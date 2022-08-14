package auth

import (
	"context"
	"project/config"
	"project/pkg/logger"
	"project/storage"

	pb "project/genproto/auth_service"
)

type authService struct {
	log  logger.LoggerI
	strg storage.StorageI
	cfg  config.Config
}

func NewAuthService(log logger.LoggerI, cfg config.Config, strg storage.StorageI) *authService {
	return &authService{
		log:  log,
		cfg:  cfg,
		strg: strg,
	}
}

func (s *authService) GetAccess(ctx context.Context, req *pb.GetAccessReq) (resp *pb.GetAccessResp, err error) {

	return
}
