package global

import (
	"github.com/ppxb/go-fiber/pkg/ms"
	"gorm.io/gorm"
)

const (
	Name    = "fiber app"
	Version = "1.0.0"
	Mode    = "debug"
)

var (
	Conf        Configuration
	ConfBox     ms.ConfBox
	RuntimeRoot string
	Mysql       *gorm.DB
)
