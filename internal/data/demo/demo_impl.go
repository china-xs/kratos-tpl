package demo

import (
	"context"
	"fmt"
)

// table demo repo implement

func (this demoRepo) i() {
	//panic("implement me")
}

func (this demoRepo) Register(ctx context.Context, username string, password string, mobile string) error {
	fmt.Println("=======", "进来了")
	return nil
}
