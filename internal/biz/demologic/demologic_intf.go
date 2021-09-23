package demologic

import (
	"context"
	"github.com/china-xs/kratos-tpl/internal/data/demo"
	"github.com/go-kratos/kratos/v2/log"
)

// package demo_interface
// DemoLogic 实现 demoLogic
var _ demoLogic = (*Logic)(nil)

type TestInfo struct {
	Id int64
}

type demoLogic interface {
	Register(ctx context.Context, username string, password string, mobile string) (*TestInfo, error)
}

type Logic struct {
	repo demo.DemoRepo
	log  *log.Helper
}

func NewDemoLogic(logger log.Logger, repo demo.DemoRepo) *Logic {
	return &Logic{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}
