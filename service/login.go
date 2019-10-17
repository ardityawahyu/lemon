package service

import "github.com/ardityawahyu/lemon/repo"

type ILoginInterface interface{}

type loginService struct {
	u repo.User
}

func NewLoginService(u repo.User) ILoginInterface {
	return &loginService{
		u: u,
	}
}

func (l *loginService) Login(username string, password string) bool {
	// TODO: login logic here
	return true
}
