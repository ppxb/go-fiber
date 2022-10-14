package req

type LoginCheck struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
