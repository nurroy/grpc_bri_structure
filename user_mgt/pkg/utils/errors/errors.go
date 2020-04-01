// Package errors is the custom error handler for BRI Open API.
package errors

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorResponse struct {
	Success         bool `json:"success,omitempty"`
	RespCode        string  `json:"respCode,omitempty"`
	RespDesc         string `json:"respDesc,omitempty"`
}

// Method is the method types.
type Method int
type Service int

const (

	FallbackError = `{"status": false, "respCode": "100", "respDesc": "internal error"}`

	// List of different Methods
	Authorization Method = iota
	InvalidFormat
	InternalServer
	Unavailable
	Other

)


// CustomHTTPError is the custom http error handler for grpc gateway.
func CustomHTTPError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, ierr error) {
	var desc string

	w.Header().Set("Content-Type", marshaler.ContentType())
	if s, ok := status.FromError(ierr); ok {
		w.WriteHeader(runtime.HTTPStatusFromCode(s.Code()))
		desc = s.Message()
	} else {
		w.WriteHeader(runtime.HTTPStatusFromCode(codes.Unknown))
		desc = ierr.Error()
	}
	b := new(ErrorResponse)
	err := json.Unmarshal([]byte(desc), b)
	if err != nil {
		w.Write([]byte(FallbackError))
	} else {
		err = json.NewEncoder(w).Encode(b)
		if err != nil {
			w.Write([]byte(FallbackError))
		}
	}
}

// FormatError is the exposed function for generating errors.
func FormatError(m Method, c codes.Code, success bool, params ...string) error {

	errRes := &ErrorResponse{
		Success:       	success,
		RespCode:       params[0],
		RespDesc:       params[1],
	}

	buf, err := json.Marshal(errRes)

	if err != nil {
		return status.Errorf(c, FallbackError)
	}
	return status.Errorf(c, string(buf))
}

// func to create error from string
func New(errs string) error {
	return errors.New(errs)
}