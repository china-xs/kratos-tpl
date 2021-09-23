package service

import (
	"context"
	"github.com/china-xs/kratos-tpl/internal/biz/demologic"

	"github.com/go-kratos/kratos/v2/log"

	pb "git.dev.enbrands.com/scrm/bed/scrm/api/demo"
)

type DemoService struct {
	pb.UnimplementedDemoServer
	log       *log.Helper
	demologic *demologic.Logic
}

func NewDemoService(logger log.Logger, logic *demologic.Logic) *DemoService {
	return &DemoService{
		log:       log.NewHelper(logger),
		demologic: logic,
	}
}

func (s *DemoService) CreateDemo(ctx context.Context, req *pb.CreateDemoRequest) (*pb.CreateDemoReply, error) {
	return &pb.CreateDemoReply{}, nil
}
