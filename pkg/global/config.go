package global

import "github.com/go-sql-driver/mysql"

type Configuration struct {
	Server ServerConf
	Mysql  MysqlConf
	Jwt    JwtConf
	Tracer TracerConf
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
