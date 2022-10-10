package initialize

import (
	"embed"
	m "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/ppxb/go-fiber/pkg/global"
	"github.com/ppxb/go-fiber/pkg/log"
	"github.com/ppxb/go-fiber/pkg/migrate"
)

var sqlFs embed.FS

func Mysql() {
	cfg, err := m.ParseDSN(global.Conf.Mysql.Uri)
	if err != nil {
		panic(errors.Wrap(err, "initialized mysql failed"))
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
		panic(errors.Wrap(err, "initialize mysql failed"))
	}
	err = binlogListen()
	if err != nil {
		panic(errors.Wrap(err, "initialize mysql binlog failed"))
	}

	log.WithContext(ctx).Info("initialize mysql success")
}
