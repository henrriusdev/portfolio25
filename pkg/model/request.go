package model

type LoginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
