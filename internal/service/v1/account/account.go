package service

import (
	"context"
	"github.com/china-xs/kratos-tpl/internal/biz/account"
	"github.com/go-kratos/kratos/v2/log"

	pb "github.com/china-xs/kratos-tpl/api/v1/account"
)

type AccountService struct {
	pb.UnimplementedAccountServer
	accBiz *account.Logic
	log    *log.Helper
}

func NewAccountService(accBiz *account.Logic, logger log.Logger) *AccountService {
	return &AccountService{
		accBiz: accBiz,
		log:    log.NewHelper(logger),
	}
}

func (s *AccountService) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountReply, error) {
	c := account.AccountLogicCreate{
		Username: req.Username,
		Password: req.Password,
		Mobile:   req.Mobile,
	}
	id, err := s.accBiz.Create(ctx, c)
	if err != nil {
		return nil, err
	}
	metadata := pb.CreateAccountReply_MetaData{
		Id:       id,
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
	}
	return &pb.CreateAccountReply{
		Code:     201,
		Message:  "ok",
		Reason:   "succ",
		Metadata: &metadata,
	}, nil
}
func (s *AccountService) UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.UpdateAccountReply, error) {
	return &pb.UpdateAccountReply{}, nil
}
func (s *AccountService) DeleteAccount(ctx context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountReply, error) {
	return &pb.DeleteAccountReply{}, nil
}
func (s *AccountService) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountReply, error) {
	return &pb.GetAccountReply{}, nil
}
func (s *AccountService) ListAccount(ctx context.Context, req *pb.ListAccountRequest) (*pb.ListAccountReply, error) {
	return &pb.ListAccountReply{}, nil
}
