package account

import (
	"context"
	"github.com/china-xs/kratos-tpl/internal/data/account"
	"github.com/china-xs/kratos-tpl/internal/data/good"
	"github.com/go-kratos/kratos/v2/log"
)

// package account_interface
// AccountLogic 实现 accountLogic
var _ accountLogic = (*Logic)(nil)

type accountLogic interface {
	Create(ctx context.Context, create AccountLogicCreate) (int32, error)
	Update(ctx context.Context)
	Get(ctx context.Context, id int32)
	List(ctx context.Context)
}

type Logic struct {
	repo     account.AccountRepo
	log      *log.Helper
	goodRepo good.Repo
}

func NewAccountLogic(logger log.Logger, repo account.AccountRepo, goodRepo good.Repo) *Logic {
	return &Logic{
		repo:     repo,
		goodRepo: goodRepo,
		log:      log.NewHelper(logger),
	}
}

type AccountLogicCreate struct {
	Username string
	Mobile   string
	Password string
}
