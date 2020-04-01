package users

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"user_mgt/pkg/utils/errors"
	tpspb "user_mgt/proto/v1/users"
	pg "user_mgt/pkg/postgres"
	)

// Method is the method types.
type Method int

const (
	// List of different Methods
	Get Method = iota
)

// ConfigStore is the common configuration store interface.
type ConfigStore interface {
	GetKey(string) string
}

// Logger is the logging interface
type Logger interface {
	Info(...interface{})
	Infof(string, ...interface{})
	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
}

// Server is the server object for this api service.
type Server struct {
	config ConfigStore
	logger Logger
}

// New creates a new server.
func New(config ConfigStore, logger Logger) (*Server, error) {
	return &Server{
		config: config,
		logger: logger,
	}, nil
}

// isValidRequest validates the inquirysaca request
func IsValidRequest(m Method, req *tpspb.RetrieveRequest) error {
	errmsg := "invalid request: %s"

	if req.UId == "" {
		return errors.FormatError(errors.InvalidFormat, codes.InvalidArgument, false, "102", fmt.Sprintf(errmsg, "missing acctNo"))
	}

	return nil
}

func (s *Server) Get(ctx context.Context, req *tpspb.RetrieveRequest) (*tpspb.RetrieveResponse, error){
	err := IsValidRequest(Get, req)
	if err != nil {
		s.logger.Errorf("[TYPES][GET] ERROR %v", err)
		return nil, err
	}

	rows, err := pg.SelectUserData(*req)
	fmt.Println("rows", rows)
	if err != nil {
		s.logger.Errorf("[TYPES][GET] ERROR %v", err)
		return nil, errors.FormatError(errors.InternalServer, codes.Internal, false, "103", fmt.Sprintf("%v", err))
	}

	var row = &tpspb.Users{
		UId:                  rows.UId,
		UName:                rows.UName,
		UAddress:             rows.UAddress,
		Age:                  string(rows.Age),
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	fmt.Println("Rows :  ", row)

	return &tpspb.RetrieveResponse{
		Success:              true,
		RespCode:             "000",
		RespDesc:             "Successfully get users data",
		User:                 row,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	},nil

}