package svc

import (
	"zero/pkg/orm"
	"zero/user/model"
	"zero/user/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel *model.UserMode
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := orm.NewMysql(&orm.Config{
		Dsn:             c.DB.Dsn,
		MaxIdleConns:    c.DB.MaxIdleConns,
		MaxOpenConns:    c.DB.MaxOpenConns,
		ConnMaxLifetime: c.DB.ConnMaxLifetime,
	})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUser(db.DB),
	}
}
