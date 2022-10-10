package global

import "github.com/go-sql-driver/mysql"

type Configuration struct {
	Server ServerConfiguration `mapstructure:"server" json:"server"`
	Mysql  MysqlConfiguration  `mapstructure:"mysql" json:"mysql"`
}

type ServerConfiguration struct {
	Port           int    `mapstructure:"port" json:"port"`
	Base           string `mapstructure:"-" json:"-"`
	UrlPrefix      string `mapstructure:"url-prefix" json:"urlPrefix"`
	ApiVersion     string `mapstructure:"api-version" json:"apiVersion"`
	ConnectTimeout int    `mapstructure:"connect-timeout" json:"connectTimeout"`
}

type MysqlConfiguration struct {
	Uri         string       `mapstructure:"uri" json:"uri"`
	TablePrefix string       `mapstructure:"table-prefix" json:"tablePrefix"`
	NoSql       bool         `mapstructure:"no-sql" json:"noSql"`
	Transaction bool         `mapstructure:"transaction" json:"transaction"`
	InitData    bool         `mapstructure:"init-data" json:"initData"`
	DSN         mysql.Config `json:"-"`
}