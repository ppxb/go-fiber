package models

import (
	"time"
)

type SysUser struct {
	M
	Username  string    `gorm:"index:idx_username,unique;comment:user login name" json:"username"`
	Password  string    `gorm:"comment:password" json:"password"`
	Mobile    string    `gorm:"comment:mobile number" json:"mobile"`
	Avatar    string    `gorm:"comment:avatar url" json:"avatar"`
	Name      string    `gorm:"comment:name" json:"name"`
	Status    *uint     `gorm:"type:tinyint(1);default:1;comment:status(0: disabled, 1: enable)" json:"status"`
	DeptId    uint      `gorm:"comment:department id" json:"deptId"`
	RoleId    uint      `gorm:"comment:role id" json:"roleId"`
	LastLogin time.Time `gorm:"comment:last login time" json:"lastLogin"`
	Locked    uint      `gorm:"type:tinyint(1);default:0;comment:locked(0: unlock, 1: locked)" json:"locked"`
}
