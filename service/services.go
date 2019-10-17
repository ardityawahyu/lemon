package service

import "github.com/ardityawahyu/lemon/repo"

type Services struct {
	Login ILoginService
	User  IUserService
}

func InitializeDependency() Services {
	ur := repo.NewUserRepo()
	l := NewLoginService(ur)
	u := NewUserService(ur)

	return Services{
		Login: l,
		User:  u,
	}
}
