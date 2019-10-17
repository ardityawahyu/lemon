package service

import "github.com/ardityawahyu/lemon/repo"

type IUserService interface {
	GetUser(id int) repo.UserModel
	AddUser(um repo.UserModel) int
	UpdateUser(um repo.UserModel) bool
}

type userService struct {
	ur repo.IUserRepo
}

func NewUserService(ur repo.IUserRepo) IUserService {
	return &userService{
		ur: ur,
	}
}

func (u *userService) GetUser(id int) repo.UserModel {
	return u.ur.GetOneUser(id)
}

func (u *userService) AddUser(um repo.UserModel) int {
	return u.ur.InsertOneUser(um)
}

func (u *userService) UpdateUser(um repo.UserModel) bool {
	return u.ur.UpdateOneUser(um)
}
