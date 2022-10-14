package response

import "time"

type User struct {
	Base
	Username  string    `json:"username"`
	Mobile    string    `json:"mobile"`
	Avatar    string    `json:"avatar"`
	Name      string    `json:"name"`
	Status    *uint     `json:"status"`
	DeptId    uint      `json:"deptId"`
	RoleId    uint      `json:"roleId"`
	LastLogin time.Time `json:"lastLogin"`
	Locked    uint      `json:"locked"`
}
