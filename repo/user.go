package repo

type IUserRepo interface {
	GetUserLogin(username string, password string) (um UserModel)
	GetOneUser(id int) UserModel
	InsertOneUser(um UserModel) int
	UpdateOneUser(um UserModel) bool
}

type userRepo struct{}

type UserModel struct {
	ID       int
	Username string
	Password string
	Email    string
}

var userData []UserModel

func NewUserRepo() IUserRepo {
	return &userRepo{}
}

func (u *userRepo) GetUserLogin(username string, password string) (um UserModel) {
	for _, ud := range userData {
		if ud.Username == username && ud.Password == password {
			um = ud
			break
		}
	}
	return
}

func (u *userRepo) GetOneUser(id int) UserModel {
	for _, ud := range userData {
		if ud.ID == id {
			return ud
		}
	}

	return UserModel{}
}

func (u *userRepo) InsertOneUser(um UserModel) int {
	newID := 0
	if len(userData) > 0 {
		newID = userData[len(userData)-1].ID + 1
	}

	um.ID = newID
	userData = append(userData, um)
	return um.ID
}

func (u *userRepo) UpdateOneUser(um UserModel) bool {
	for _, ud := range userData {
		if um.ID == ud.ID {
			userData[um.ID] = um
			return true
		}
	}
	return false
}
