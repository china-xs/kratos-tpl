package account

import (
	"context"
	"git.dev.enbrands.com/scrm/bed/scrm/app/auth/internal/data/account"
	"git.dev.enbrands.com/scrm/bed/scrm/app/auth/internal/data/good"
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
	goodRepo good.GoodRepo
}

func NewAccountLogic(logger log.Logger, repo account.AccountRepo, goodRepo good.GoodRepo) *Logic {
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
