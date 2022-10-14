package service

import (
	"context"
	"github.com/ppxb/go-fiber/pkg/global"
	"github.com/ppxb/go-fiber/pkg/query"
)

type MysqlService struct {
	Q query.Mysql
}

func New(ctx context.Context) MysqlService {
	ops := []func(*query.MysqlOptions){
		query.WithMysqlCtx(ctx),
		query.WithMysqlDb(global.Mysql),
	}

	my := MysqlService{
		Q: query.NewMysql(ops...),
	}
	return my
}
