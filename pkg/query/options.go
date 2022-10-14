package query

import (
	"context"
	"github.com/ppxb/go-fiber/pkg/utils"
	"gorm.io/gorm"
)

type MysqlOptions struct {
	ctx context.Context
	db  *gorm.DB
}

func WithMysqlDb(db *gorm.DB) func(*MysqlOptions) {
	return func(options *MysqlOptions) {
		if db != nil {
			getMysqlOptions(options).db = db
		}
	}
}

func WithMysqlCtx(ctx context.Context) func(*MysqlOptions) {
	return func(options *MysqlOptions) {
		if !utils.InterfaceIsNil(ctx) {
			getMysqlOptions(options).ctx = ctx
		}
	}
}

func getMysqlOptions(options *MysqlOptions) *MysqlOptions {
	if options == nil {
		return &MysqlOptions{
			ctx: context.Background(),
		}
	}
	return options
}
