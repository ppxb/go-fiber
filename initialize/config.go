package initialize

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/ppxb/go-fiber/pkg/global"
	"github.com/ppxb/go-fiber/pkg/log"
	"github.com/ppxb/go-fiber/pkg/ms"
	"github.com/spf13/viper"
	"strings"
)

const (
	configType            = "yml"
	configDir             = "conf"
	developmentConfig     = "debug.yml"
	productionConfig      = "release.yml"
	defaultConnectTimeout = 5
)

var ctx context.Context

func Config(c context.Context, conf embed.FS) {
	ctx = c
	box := ms.ConfBox{
		Ctx: ctx,
		Fs:  conf,
		Dir: configDir,
	}
	global.ConfBox = box

	var configName string
	v := viper.New()
	if strings.ToLower(global.Mode) == gin.ReleaseMode {
		configName = productionConfig
	} else {
		configName = developmentConfig
	}

	readConfig(box, v, configName)
	settings := v.AllSettings()
	for index, setting := range settings {
		v.SetDefault(index, setting)
	}

	if err := v.Unmarshal(&global.Conf); err != nil {
		panic(errors.Wrapf(err, "initialize config failed, config env: %s_CONF: %s", global.Mode, box.Dir))
	}

	if global.Conf.Server.ConnectTimeout < 1 {
		global.Conf.Server.ConnectTimeout = defaultConnectTimeout
	}

	if strings.TrimSpace(global.Conf.Server.UrlPrefix) == "" {
		global.Conf.Server.UrlPrefix = "api"
	}

	if strings.TrimSpace(global.Conf.Server.ApiVersion) == "" {
		global.Conf.Server.ApiVersion = "v1"
	}

	global.Conf.Server.Mode = global.Mode
	global.Conf.Server.Base = fmt.Sprintf("/%s/%s", global.Conf.Server.UrlPrefix, global.Conf.Server.ApiVersion)

	log.WithContext(ctx).Info("initialize config success, config env: %s", global.Mode)
}

func readConfig(box ms.ConfBox, v *viper.Viper, configFile string) {
	v.SetConfigType(configType)
	config := box.Get(configFile)
	if len(config) == 0 {
		log.WithContext(box.Ctx).Error("initialize configs failed, configs env = [%s] file = [%s]", strings.Split(configFile, ".")[0], configFile)
	}
	if err := v.ReadConfig(bytes.NewReader(config)); err != nil {
		log.WithContext(box.Ctx).Error(errors.Wrapf(err, "initialize configs failed, configs env = [%s] file = [%s]", strings.Split(configFile, ".")[0], configFile))
	}
}
