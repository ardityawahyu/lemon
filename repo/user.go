package repo

type IUserRepo interface{}

type userRepo struct{}

type UserModel struct {
	Username string
	Password string
	Email    string
}

func NewUserRepo() IUserRepo {
	return &userRepo{}
}

func (u *userRepo) GetOneUser(username string, password string) (um UserModel) {
	// TODO: getting data user here
	return
}

func (u *userRepo) InsertOneUser(um UserModel) (id int) {
	// TODO: insert user
	return
}

func (u *userRepo) UpdateOneUser(um UserModel) (updatedRowCount int) {
	// TODO: update user
	return
}
