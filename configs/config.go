package configs

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/ppxb/go-fiber/pkg/env"
	"github.com/ppxb/go-fiber/pkg/log"
	"github.com/ppxb/go-fiber/pkg/ms"
	"github.com/spf13/viper"
	"strings"
)

//go:embed conf/*.yml
var conf embed.FS

var ctx = context.Background()

var Conf *Config

const (
	configType            = "yml"
	configDir             = "./configs/conf"
	developmentConfig     = "dev.yml"
	productionConfig      = "prod.yml"
	defaultConnectTimeout = 5
)

type Config struct {
	Server `mapstructure:"server" json:"server"`
	Mysql  `mapstructure:"mysql" json:"mysql"`
	Jwt    `mapstructure:"jwt" json:"jwt"`
}

type Server struct {
	Name           string `mapstructure:"name" json:"name"`
	Port           int    `mapstructure:"port" json:"port"`
	Version        string `mapstructure:"version" json:"version"`
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

func init() {
	box := ms.ConfBox{
		Ctx: ctx,
		Fs:  conf,
		Dir: configDir,
	}

	var configFile string
	v := viper.New()
	if env.Active().Value() == "prod" {
		configFile = productionConfig
	} else {
		configFile = developmentConfig
	}

	readConfig(box, v, configFile)
	settings := v.AllSettings()
	for index, setting := range settings {
		v.SetDefault(index, setting)
	}

	if err := v.Unmarshal(&Conf); err != nil {
		panic(errors.Wrapf(err, "[Server] initialize configs failed, configs env = [%s] path = [%s/%s]", env.Active().Value(), box.Dir, configFile))
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

	log.WithContext(ctx).Info("[Server] initialize configs success.configs env = [%s] path = [%s/%s]", env.Active().Value(), box.Dir, configFile)
}

func readConfig(box ms.ConfBox, v *viper.Viper, configFile string) {
	v.SetConfigType(configType)
	config := box.Get(configFile)
	if len(config) == 0 {
		panic(fmt.Sprintf("[Server] initialize configs failed, configs env = [%s] path = [%s/%s]", env.Active().Value(), box.Dir, configFile))
	}
	if err := v.ReadConfig(bytes.NewReader(config)); err != nil {
		panic(errors.Wrapf(err, "[Server] initialize configs failed, configs env = [%s] path = [%s/%s]", env.Active().Value(), box.Dir, configFile))
	}
}
