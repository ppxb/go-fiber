package initialize

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ppxb/go-fiber/pkg/global"
	"github.com/ppxb/go-fiber/pkg/log"
	"github.com/ppxb/go-fiber/pkg/ms"
	"github.com/spf13/viper"
	"strings"
)

const (
	configType            = "yml"
	configDir             = "./initialize/conf"
	developmentConfig     = "dev.yml"
	productionConfig      = "prod.yml"
	defaultConnectTimeout = 5
)

//go:embed conf/*.yml
var conf embed.FS

func Config(c context.Context) {
	box := ms.ConfBox{
		Ctx: c,
		Fs:  conf,
		Dir: configDir,
	}
	global.ConfBox = box

	var configName string
	v := viper.New()
	if strings.ToLower(global.ProMode) == "release" {
		configName = productionConfig
	} else {
		configName = developmentConfig
	}

	// read configs in global configs box
	readConfig(box, v, configName)
	settings := v.AllSettings()
	for index, setting := range settings {
		v.SetDefault(index, setting)
	}

	if err := v.Unmarshal(&global.Conf); err != nil {
		panic(errors.Wrapf(err, "[Server] initialize configs failed, configs env = [%s] path = [%s/%s]", global.ProMode, box.Dir, configName))
	}

	// initialize some other options
	if global.Conf.Server.ConnectTimeout < 1 {
		global.Conf.Server.ConnectTimeout = defaultConnectTimeout
	}

	if strings.TrimSpace(global.Conf.Server.UrlPrefix) == "" {
		global.Conf.Server.UrlPrefix = "api"
	}

	if strings.TrimSpace(global.Conf.Server.ApiVersion) == "" {
		global.Conf.Server.ApiVersion = "v1"
	}

	global.Conf.Server.Base = fmt.Sprintf("/%s/%s", global.Conf.Server.UrlPrefix, global.Conf.Server.ApiVersion)

	log.WithContext(c).Info("[Server] initialize configs success.configs env = [%s] path = [%s/%s]", global.ProMode, box.Dir, configName)
}

func readConfig(box ms.ConfBox, v *viper.Viper, configFile string) {
	v.SetConfigType(configType)
	config := box.Get(configFile)
	if len(config) == 0 {
		panic(fmt.Sprintf("[Server] initialize configs failed, configs env = [%s] path = [%s/%s]", global.ProMode, box.Dir, configFile))
	}
	if err := v.ReadConfig(bytes.NewReader(config)); err != nil {
		panic(errors.Wrapf(err, "[Server] initialize configs failed, configs env = [%s] path = [%s/%s]", global.ProMode, box.Dir, configFile))
	}
}
