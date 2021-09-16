package data

import (
	"github.com/china-xs/kratos-tpl/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

// ProviderSet is data providers.
// 数据库 New方法必须在这里注册
//var ProviderSet = wire.NewSet(
//	NewData,
//	NewAccountRepo,
//	good.NewGoodRepo,
//)

// Data .
type Data struct {
	DB  *gorm.DB
	RDB *redis.Client
	Log log.Helper
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	l := log.NewHelper(logger)
	//db 初始化 暂时直接链接gorm 后面配置多库后根据配置再调整
	dsn := c.Database.Source
	config := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "u_",
			SingularTable: true,
		},
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	}
	db, err := gorm.Open(mysql.Open(dsn), &config)
	if err != nil {
		l.Errorf("db connect err:%v", err)
		return nil, nil, err
	}
	//redis 初始化 项目初期不配置可以注释
	redisOps := redis.Options{
		Addr: c.Redis.Addr,
	}
	rdb := redis.NewClient(&redisOps)
	//redis tracing
	rdb.AddHook(redisotel.TracingHook{})

	return &Data{
		DB:  db,
		RDB: rdb,
		Log: *log.NewHelper(logger),
	}, cleanup, nil
}
