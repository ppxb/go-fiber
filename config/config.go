package config

import (
	"bytes"
	"context"
	"embed"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/ppxb/go-fiber/pkg/logger"
	"github.com/ppxb/go-fiber/pkg/ms"
	"github.com/spf13/viper"
	"strings"
)

const (
	configType            = "yml"
	configDir             = "./config/conf"
	developmentConfig     = "debug.yml"
	productionConfig      = "release.yml"
	defaultConnectTimeout = 5
)

//go:embed conf/*.yml
var conf embed.FS

var Conf Config

type Config struct {
	Server
	Mysql
	Jwt
}

type Server struct {
	Mode           string `mapstructure:"mode" json:"mode"`
	Port           int    `mapstructure:"port" json:"port"`
	UrlPrefix      string `mapstructure:"url-prefix" json:"urlPrefix"`
	ApiVersion     string `mapstructure:"api-version" json:"apiVersion"`
	ConnectTimeout int    `mapstructure:"connect-timeout" json:"connectTimeout"`
}

type Mysql struct {
	Uri         string       `mapstructure:"uri" json:"uri"`
	TablePrefix string       `mapstructure:"table-prefix" json:"tablePrefix"`
	NoSql       bool         `mapstructure:"no-sql" json:"noSql"`
	Transaction bool         `mapstructure:"transaction" json:"transaction"`
	InitData    bool         `mapstructure:"init-data" json:"initData"`
	DSN         mysql.Config `json:"-"`
}

type Jwt struct {
	Realm      string `mapstructure:"realm" json:"realm"`
	Key        string `mapstructure:"key" json:"key"`
	Timeout    int    `mapstructure:"timeout" json:"timeout"`
	MaxRefresh int    `mapstructure:"max-refresh" json:"maxRefresh"`
}

func Init(c context.Context, mode string) {
	box := ms.ConfBox{
		Ctx: c,
		Fs:  conf,
		Dir: configDir,
	}

	var configName string
	v := viper.New()
	if strings.ToLower(mode) == "release" {
		configName = productionConfig
	} else {
		configName = developmentConfig
	}

	readConfig(box, v, configName)
	settings := v.AllSettings()
	for index, setting := range settings {
		v.SetDefault(index, setting)
	}

	if err := v.Unmarshal(&Conf); err != nil {
		logger.WithContext(c).Panic(errors.Wrapf(err, "[Server] initialize configs failed, configs env = [%s] file = [%s]", mode, configName))
	}

	if Conf.Server.ConnectTimeout < 1 {
		Conf.Server.ConnectTimeout = defaultConnectTimeout
	}

	if strings.TrimSpace(Conf.Server.UrlPrefix) == "" {
		Conf.Server.UrlPrefix = "api"
	}

	if strings.TrimSpace(Conf.Server.ApiVersion) == "" {
		Conf.Server.ApiVersion = "v1"
	}

	if strings.TrimSpace(Conf.Server.Mode) == "" {
		Conf.Server.Mode = mode
	}

	logger.WithContext(c).Info("[Server] Config initialize successful.")
}

func readConfig(box ms.ConfBox, v *viper.Viper, configFile string) {
	v.SetConfigType(configType)
	config := box.Get(configFile)
	if len(config) == 0 {
		logger.WithContext(box.Ctx).Panicf("[Server] initialize configs failed, configs env = [%s] file = [%s]", strings.Split(configFile, ".")[0], configFile)
	}
	if err := v.ReadConfig(bytes.NewReader(config)); err != nil {
		logger.WithContext(box.Ctx).Panic(errors.Wrapf(err, "[Server] initialize configs failed, configs env = [%s] file = [%s]", strings.Split(configFile, ".")[0], configFile))
	}
}
