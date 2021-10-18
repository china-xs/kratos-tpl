package main

import (
	"fmt"
	"github.com/golang/mock/mockgen/model"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

const (
	host   = "127.0.0.1"
	user   = "***"
	pwd    = "***"
	port   = 3306
	dbname = "db_username"
	preFix = "u_"
)

func main() {
	// 指定生成代码的具体目录，默认为：
	//gen.WithoutContext 不强制使用ctx 调用方法需要自带WithContext(ctx)
	g := gen.NewGenerator(gen.Config{
		OutPath: "../../internal/data/dao/query",
		//Mode:    gen.WithoutContext, //不限制上下文使用,实际业务必须强制上下文
	})
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		user,
		pwd,
		host,
		port,
		dbname,
	)
	config := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   preFix,
			SingularTable: true,
		},
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	}
	db, err := gorm.Open(mysql.Open(dsn), &config)
	if err != nil {
		fmt.Printf("链接异常 %s \n", err)
		return
	}
	// 复用工程原本使用的SQL连接配置
	g.UseDB(db)
	// 所有需要实现查询方法的结构体 增加表不能把原来表删除
	g.ApplyBasic(
		g.GenerateModel("table_name"),
	)

	// 为指定的数据库表实现除基础方法外的相关方法
	g.ApplyInterface(
		func(method model.Method) {},
	)

	// 执行并生成代码
	g.Execute()
}
