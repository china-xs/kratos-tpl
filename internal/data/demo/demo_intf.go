package demo

import (
	"context"
	"github.com/china-xs/kratos-tpl/internal/data"
)

// data table demo repo 接口定义
var _ DemoRepo = (*demoRepo)(nil)

type demoRepo struct {
	data *data.Data
}

//db cache 接口定义
type DemoRepo interface {
	i()
	Register(ctx context.Context, username string, password string, mobile string) error
}

func NewDemoRepo(data *data.Data) DemoRepo {
	return &demoRepo{
		data: data,
	}
}
