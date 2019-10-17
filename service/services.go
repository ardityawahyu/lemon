package service

type Services struct {
	Login ILoginService
	User  IUserService
}

func InitializeDependency() Services {
	l := NewLoginService()
	u := NewUserService()

	return Services{
		Login: l,
		User:  u,
	}
}
