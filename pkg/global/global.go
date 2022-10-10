package global

import (
	"github.com/ppxb/go-fiber/pkg/ms"
	"gorm.io/gorm"
)

var (
	Conf        Configuration
	ConfBox     ms.ConfBox
	RuntimeRoot string
	Mysql       *gorm.DB
)
