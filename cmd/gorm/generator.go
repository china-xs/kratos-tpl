package main

import (
	"fmt"
	"git.dev.enbrands.com/scrm/bed/scrm/pkg/db"
	"github.com/golang/mock/mockgen/model"
	"gorm.io/gen"
)

const (
	host     = "127.0.0.1"
	user     = "prod"
	pwd      = "Prodc959ed5e61ce4451803Enbrands"
	port     = 3308
	dataName = "yjf_scrm"
	preFix   = "u_"
)

func main() {
	// 指定生成代码的具体目录，默认为：
	g := gen.NewGenerator(gen.Config{OutPath: "../../internal/data/dao/query"})
	DB, err := db.NewDB(
		db.Host(host),
		db.User(user),
		db.Pwd(pwd),
		db.Name(dataName),
		db.Port(port),
		db.PreFix(preFix),
	).GetDB("default")
	if err != nil {
		fmt.Printf("链接异常 %s \n", err)
		return
	}
	// 复用工程原本使用的SQL连接配置
	g.UseDB(DB)
	// 所有需要实现查询方法的结构体 增加表不能把原来表删除
	g.ApplyBasic(
		g.GenerateModel(
			"u_account",
			gen.FieldIgnore("create_at"),
			gen.FieldIgnore("update_at"),
		),
	)
	// 为指定的数据库表实现除基础方法外的相关方法
	g.ApplyInterface(
		func(method model.Method) {},
	)
	// 执行并生成代码
	g.Execute()
}
