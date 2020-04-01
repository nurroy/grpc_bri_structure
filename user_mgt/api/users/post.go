package users

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	pg "user_mgt/pkg/postgres"
	"user_mgt/pkg/utils/errors"
	tpspb "user_mgt/proto/v1/users"
)

const (
	// List of different Methods
	POST Method = iota
)


func (s *Server) Post(ctx context.Context, req *tpspb.Users) (*tpspb.InsertResponse, error){

	_, err := pg.InsertUserData(*req)
	if err != nil {
		s.logger.Errorf("[TYPES][POST] ERROR %v", err)
		return nil, errors.FormatError(errors.InternalServer, codes.Internal, false, "103", fmt.Sprintf("%v", err))
	}

	return &tpspb.InsertResponse{
		Success:              true,
		RespCode:             "000",
		RespDesc:             "Successfully Insert Data",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	},nil

}