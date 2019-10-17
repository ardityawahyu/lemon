package service

import (
	"reflect"

	"github.com/ardityawahyu/lemon/repo"
)

type ILoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	u repo.IUserRepo
}

func NewLoginService(u repo.IUserRepo) ILoginService {
	return &loginService{
		u: u,
	}
}

func (l *loginService) Login(username string, password string) bool {
	um := l.u.GetUserLogin(username, password)
	if reflect.DeepEqual(um, repo.UserModel{}) {
		return false
	}

	return true
}
