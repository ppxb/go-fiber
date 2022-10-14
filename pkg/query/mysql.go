package query

import (
	"context"
	"gorm.io/gorm"
)

type Mysql struct {
	ops MysqlOptions
	Ctx context.Context
	Tx  *gorm.DB
	Db  *gorm.DB
}

func NewMysql(options ...func(*MysqlOptions)) Mysql {
	ops := getMysqlOptions(nil)
	for _, f := range options {
		f(ops)
	}

	if ops.db == nil {
		panic("[DATABASE] MYSQL db is empty")
	}
	my := Mysql{}
	my.Db = ops.db.WithContext(ops.ctx)
	return my
}
