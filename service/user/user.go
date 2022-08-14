package user

import (
	"context"
	"project/config"
	pb "project/genproto/auth_service"
	"project/pkg/logger"
	"project/storage"
)

type UserService struct {
	log  logger.LoggerI
	cfg  config.Config
	strg storage.StorageI
	pb.UnimplementedUserServiceServer
}

func NewUserService(log logger.LoggerI, strg storage.StorageI, cfg config.Config) *UserService {
	return &UserService{
		log: log,
		cfg: cfg,
		strg: strg,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, req *pb.UserRegisterReq) (resp *pb.UserRegisterResp, err error) {

	resp = &pb.UserRegisterResp{}


	return resp, nil
}

