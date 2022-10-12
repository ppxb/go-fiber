package initialize

import (
	"context"
	"embed"
	"fmt"
	m "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/ppxb/go-fiber/pkg/global"
	"github.com/ppxb/go-fiber/pkg/log"
	"github.com/ppxb/go-fiber/pkg/migrate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

//go:embed db/*.sql
var sqlFs embed.FS

func Mysql(ctx context.Context) {
	cfg, err := m.ParseDSN(global.Conf.Mysql.Uri)
	if err != nil {
		panic(errors.Wrap(err, "[Server] initialized mysql failed"))
	}
	global.Conf.Mysql.DSN = *cfg
	uri := global.Conf.Mysql.Uri
	err = migrate.Do(
		migrate.WithCtx(ctx),
		migrate.WithUri(uri),
		migrate.WithFs(sqlFs),
		migrate.WithFsRoot("db"),
		migrate.WithBefore(beforeMigrate),
	)
	if err != nil {
		panic(errors.Wrap(err, "[Server] initialize mysql failed"))
	}
	// initialized binlog listener
	//err = binlogListen()
	//if err != nil {
	//	panic(errors.Wrap(err, "initialize mysql binlog failed"))
	//}

	log.WithContext(ctx).Info("[Server] initialize mysql success")
}

func beforeMigrate(ctx context.Context) (err error) {
	init := false
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(global.Conf.Server.ConnectTimeout)*time.Second)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				if !init {
					panic(fmt.Sprintf("[Server] initialize mysql failed: connect timeout(%ds)", global.Conf.Server.ConnectTimeout))
				}

				// avoid goroutine deadlock
				return
			}
		}
	}()

	l := log.NewDefaultGormLogger()
	if global.Conf.Mysql.NoSql {
		l = l.LogMode(glogger.Silent)
	} else {
		l = l.LogMode(glogger.Info)
	}
	var db *gorm.DB
	db, err = gorm.Open(mysql.Open(global.Conf.Mysql.DSN.FormatDSN()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   global.Conf.Mysql.TablePrefix + "_",
			SingularTable: true,
		},
		// select * from xxx => select a,b,c from xxx
		QueryFields: true,
		Logger:      l,
	})
	if err != nil {
		return
	}
	init = true
	global.Mysql = db
	//autoMigrate(ctx)
	return
}

//func autoMigrate(ctx context.Context) {
//	// migrate tables change to sql-migrate: initialize/db/***.sql
//	// auto migrate fsm
//	fsm.Migrate(fsm.WithDb(global.Mysql), fsm.WithCtx(ctx))
//}
