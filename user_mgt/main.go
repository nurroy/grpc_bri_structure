package main

import (
	"user_mgt/api/users"
	tpspb "user_mgt/proto/v1/users"
	"context"
	"log"
	"github.com/brankas/envcfg"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	pg "user_mgt/pkg/postgres"
	"user_mgt/pkg/utils/errors"
)
var (
	config *envcfg.Envcfg
	logger *logrus.Logger
)

func main() {
	var err error
	config, logger, err = Setup()
	if err != nil {
		log.Fatal(err)
	}

	logger.Infof("[SERVER] Environment %s is ready", config.Env())

	logger.Infof("[POSTGRES] Starting to initialize postgres conf")

	// setup postgres
	err = pg.New(config)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Infof("[POSTGRES] Finish initialize postgres conf")

	// create run group
	g, _ := errgroup.WithContext(context.Background())

	var servers []*http.Server
	// goroutine to check for signals to gracefully finish all functions
	g.Go(func() error {
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

		select {
		case sig := <-signalChannel:
			logger.Infof("received signal: %s\n", sig)

			for i, s := range servers {
				if err := s.Shutdown(context.Background()); err != nil {
					if err == nil {
						logger.Infof("error shutting down server %d: %v", i, err)
						return err
					}
				}
			}
			os.Exit(1)
		}
		return nil
	})

	g.Go(func() error {return NewGRPCServer(config, logger)})
	g.Go(func() error {return NewHTTPServer()})

	if err := g.Wait(); !IgnoreErr(err) {
		logger.Fatal(err)
	}
	logger.Infoln("done.")
}

// setup sets up the environment.
func Setup() (*envcfg.Envcfg, *logrus.Logger, error) {
	// create config and logger
	config, err := envcfg.New()
	if err != nil {
		return nil, nil, err
	}
	logger := logrus.New()

	// force all writes to regular log to logger
	log.SetOutput(logger.Writer())
	log.SetFlags(0)

	// configure logging for environment
	tf := new(logrus.TextFormatter)
	//tf.ForceColors = logrus.IsTerminal()
	tf.FullTimestamp = true
	logger.Formatter = tf

	return config, logger, nil
}

// ignoreErr returns true when err can be safely ignored.
func IgnoreErr(err error) bool {
	switch {
	case err == nil || err == http.ErrServerClosed || err == cmux.ErrListenerClosed:
		return true
	}
	if opErr, ok := err.(*net.OpError); ok {
		return opErr.Err.Error() == "use of closed network connection"
	}
	return false
}

// newGRPCServer creates the grpc server serve mux.
func NewGRPCServer(config *envcfg.Envcfg, logger *logrus.Logger) error {
	// create service handlers each grpc service server
	typesSvc, err := users.New(config, logger)
	if err != nil {
		return err
	}

	lis, err := net.Listen("tcp", ":"+string(config.GetKey("grpc.port")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	// register grpc service server
	grpcServer := grpc.NewServer()
	tpspb.RegisterTypesServiceServer(grpcServer, typesSvc)

	// add reflection service
	reflection.Register(grpcServer)

	// running gRPC server
	log.Println("[SERVER] GRPC server is ready")
	grpcServer.Serve(lis)

	return nil
}

// newHTTPServer creates the http server serve mux.
func NewHTTPServer() error {
	gwruntime.HTTPError = errors.CustomHTTPError

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Connect to the GRPC server
	addr := "0.0.0.0:"+ string(config.GetKey("grpc.port"))
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
		return err
	}
	defer conn.Close()

	// Create new grpc-gateway
	rmux := gwruntime.NewServeMux()

	// register gateway endpoints
	for _, f := range []func(ctx context.Context, mux *gwruntime.ServeMux, conn *grpc.ClientConn) error{
		// register grpc service handler
		tpspb.RegisterTypesServiceHandler,
	} {
		if err = f(ctx, rmux, conn); err != nil {
			log.Fatal(err)
			return err
		}
	}


	// create http server mux
	mux := http.NewServeMux()
	mux.Handle("/", rmux)

	// run swagger server
	if config.GetKey("runtime.environment") == "development" {
		NewHTTPEncodedAndSwaggerServer(mux)
	}

	// running rest http server
	log.Println("[SERVER] REST HTTP server is ready")
	err = http.ListenAndServe("0.0.0.0:"+config.PortString(), mux)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// newHTTPEncodedAndSwaggerServer creates the swagger server serve mux.
func NewHTTPEncodedAndSwaggerServer(gwmux *http.ServeMux) {
	// register swagger service server
	gwmux.HandleFunc("/api/customer/v1.0/docs.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "swagger/docs.json")
	})

	// load swagger-ui file
	fs := http.FileServer(http.Dir("swagger/swagger-ui"))
	gwmux.Handle("/api/customer/v1.0/docs/", http.StripPrefix("/api/customer/v1.0/docs", fs))
}