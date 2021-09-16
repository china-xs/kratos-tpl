package good

import (
	"context"
	"github.com/china-xs/kratos-tpl/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

var _ Repo = (*goodRepo)(nil)

type goodRepo struct {
	data *data.Data
	log  *log.Helper
}

func (g goodRepo) Create(ctx context.Context, id int32) error {
	panic("implement me")
}

type Repo interface {
	Create(ctx context.Context, id int32) error
}

func NewGoodRepo(data *data.Data, logger log.Logger) Repo {
	return &goodRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
