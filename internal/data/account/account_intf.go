package account

import (
	"context"
	data2 "git.dev.enbrands.com/scrm/bed/scrm/app/auth/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

// data table account repo 接口定义
var _ AccountRepo = (*accountRepo)(nil)

type accountRepo struct {
	data *data2.Data
	log  *log.Helper
}

//db cache 接口定义
type AccountRepo interface {
	i()
	Create(ctx context.Context, create AccountCreate) (int32, error)
	Update(ctx context.Context)
	Get(ctx context.Context, id int32)
	Delete(ctx context.Context, id int32)
	//List(ctx context.Context)
}

func NewAccountRepo(data *data2.Data, logger log.Logger) AccountRepo {
	return &accountRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type AccountCreate struct {
	Username string
	name     string
	Password string
	Mobile   string
}

//// NewGreeterRepo .
//func NewAccountRepo(data *Data, logger log.Logger) biz.GreeterRepo {
//	return &greeterRepo{
//		data: data,
//		log:  log.NewHelper(logger),
//	}
//}
