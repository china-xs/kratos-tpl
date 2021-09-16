package account

// account implement 接口实现
import (
	"context"
	"git.dev.enbrands.com/scrm/bed/scrm/app/auth/internal/data/account"
)

func (this Logic) Create(ctx context.Context, create AccountLogicCreate) (int32, error) {
	d := account.AccountCreate{
		Username: create.Username,
		Password: create.Password,
		Mobile:   create.Mobile,
	}
	//do some logic
	return this.repo.Create(ctx, d)
}

func (this Logic) Update(ctx context.Context) {
	panic("implement me")
}

func (this Logic) Get(ctx context.Context, id int32) {
	panic("implement me")
}

func (this Logic) List(ctx context.Context) {
	panic("implement me")
}
