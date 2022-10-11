package migrate

import (
	"database/sql"
	"fmt"
	m "github.com/go-sql-driver/mysql"
	"github.com/ppxb/go-fiber/pkg/log"
)

func Do(options ...func(*Options)) (err error) {
	ops := getOptions(nil)
	for _, f := range options {
		f(ops)
	}

	err = database(ops)
	if err != nil {
		return
	}

	//var db *sql.DB
	//db, err = sql.Open(ops.driver, ops.uri)
	//if err != nil {
	//	log.WithContext(ops.ctx).WithError(err).Error("open %s(%s) failed", ops.driver, ops.uri)
	//	return
	//}
	//
	//defer func() {
	//
	//}()

	return
}

func database(ops *Options) (err error) {
	var cfg *m.Config
	var db *sql.DB
	cfg, err = m.ParseDSN(ops.uri)
	if err != nil {
		log.WithContext(ops.ctx).WithError(err).Error("invalid uri")
		return
	}
	dbname := cfg.DBName
	cfg.DBName = ""
	db, err = sql.Open(ops.driver, cfg.FormatDSN())
	if err != nil {
		return
	}
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbname))
	if err != nil {
		log.WithContext(ops.ctx).WithError(err).Error("create database failed")
	}
	return
}
