package global

import (
	"github.com/go-sql-driver/mysql"
	"github.com/ppxb/go-fiber/pkg/log"
)

type Configuration struct {
	Server ServerConf
	Mysql  MysqlConf
	Jwt    JwtConf
	Tracer TracerConf
	Log    LogConf
}

type ServerConf struct {
	Mode           string `mapstructure:"mode" json:"mode"`
	Port           int    `mapstructure:"port" json:"port"`
	Base           string `mapstructure:"-" json:"-"`
	UrlPrefix      string `mapstructure:"url-prefix" json:"urlPrefix"`
	ApiVersion     string `mapstructure:"api-version" json:"apiVersion"`
	ConnectTimeout int    `mapstructure:"connect-timeout" json:"connectTimeout"`
}

type MysqlConf struct {
	Uri         string       `mapstructure:"uri" json:"uri"`
	TablePrefix string       `mapstructure:"table-prefix" json:"tablePrefix"`
	NoSql       bool         `mapstructure:"no-sql" json:"noSql"`
	Transaction bool         `mapstructure:"transaction" json:"transaction"`
	InitData    bool         `mapstructure:"init-data" json:"initData"`
	DSN         mysql.Config `json:"-"`
}

type JwtConf struct {
	Realm      string `mapstructure:"realm" json:"realm"`
	Key        string `mapstructure:"key" json:"key"`
	Timeout    int    `mapstructure:"timeout" json:"timeout"`
	MaxRefresh int    `mapstructure:"max-refresh" json:"maxRefresh"`
}

type TracerConf struct {
	Enable   bool              `mapstructure:"enable" json:"enable"`
	Insecure bool              `mapstructure:"insecure" json:"insecure"`
	Endpoint string            `mapstructure:"endpoint" json:"endpoint"`
	Headers  map[string]string `mapstructure:"headers" json:"headers"`
}

type LogConf struct {
	Level                    log.Level       `mapstructure:"level" json:"level"`
	Json                     bool            `mapstructure:"json" json:"json"`
	LineNum                  LogsLineNumConf `mapstructure:"line-num" json:"lineNum"`
	OperationKey             string          `mapstructure:"operation-key" json:"operationKey"`
	OperationAllowedToDelete bool            `mapstructure:"operation-allowed-to-delete" json:"operationAllowedToDelete"`
}

type LogsLineNumConf struct {
	Disable bool `mapstructure:"disable" json:"disable"`
	Level   int  `mapstructure:"level" json:"level"`
	Version bool `mapstructure:"version" json:"version"`
	Source  bool `mapstructure:"source" json:"source"`
}
