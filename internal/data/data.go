package data

import (
	"context"
	"github.com/china-xs/kratos-tpl/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

// Data init
type Data struct {
	//TODO wrapped database client
	Log log.Helper
	db  *gorm.DB
	rdb *redis.Client
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
		db:  db,
		rdb: rdb,
		Log: *log.NewHelper(logger),
	}, cleanup, nil
}

type Dao interface {
	i()
	GetDb(ctx context.Context) *gorm.DB
	GetRdb(ctx context.Context) *redis.Client
	GetLog(ctx context.Context) *log.Helper
}

func (this Data) GetDb(ctx context.Context) *gorm.DB {
	return this.db
}
func (this Data) GetRdb(ctx context.Context) *redis.Client {
	return this.rdb.WithContext(ctx)
}

func (this Data) GetLog(ctx context.Context) *log.Helper {
	return this.Log.WithContext(ctx)
}

func (this Data) i() {

}
