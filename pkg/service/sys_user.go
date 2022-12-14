package service

import (
	"github.com/pkg/errors"
	"github.com/ppxb/go-fiber/models"
	"github.com/ppxb/go-fiber/pkg/req"
	"github.com/ppxb/go-fiber/pkg/response"
	"github.com/ppxb/go-fiber/pkg/utils"
	"time"
)

func (my MysqlService) LoginCheck(r req.LoginCheck) (u models.SysUser, err error) {
	err = my.Q.Db.Where("username = ?", r.Username).First(&u).Error
	if err != nil {
		err = errors.Errorf(response.LoginCheckErrorMsg)
		return
	}

	if ok := utils.ComparePwd(r.Password, u.Password); !ok {
		err = errors.Errorf(response.LoginCheckErrorMsg)
		return
	}
	err = my.UserLastLogin(u.Id)
	return
}

func (my MysqlService) UserLastLogin(id uint) (err error) {
	m := make(map[string]interface{})
	m["last_login"] = time.Now()
	err = my.Q.Db.
		Model(&models.SysUser{}).
		Where("id = ?", id).
		Updates(&m).
		Error
	return
}
