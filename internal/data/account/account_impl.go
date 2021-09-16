package account

import (
	"context"
	"git.dev.enbrands.com/scrm/bed/scrm/app/auth/internal/data/dao/model"
	"git.dev.enbrands.com/scrm/bed/scrm/app/auth/internal/data/dao/query"
)

// table account repo implement

func (this accountRepo) i() {
	//panic("implement me")
}

func (this accountRepo) Create(ctx context.Context, create AccountCreate) (int32, error) {
	account := model.Account{
		Name:     create.name,
		Mobile:   create.Mobile,
		Password: create.Password,
	}
	db := this.data.db.WithContext(ctx)
	//Account.Omit("code") 不更新|写入 code
	//Account.Select("code", "Username") 指定更新字段
	err := query.Use(db).Account.Create(&account)
	if err != nil {
		this.log.WithContext(ctx).Warnf("create account err %v", err)
		return 0, err
	}
	return account.ID, nil
}

func (this accountRepo) Update(ctx context.Context) {
	panic("implement me")
}

func (this accountRepo) Get(ctx context.Context, id int32) {
	panic("implement me")
}

func (this accountRepo) Delete(ctx context.Context, id int32) {
	panic("implement me")
}
