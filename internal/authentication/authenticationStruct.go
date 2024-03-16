package authentication

type User struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Access   int    `json:"access"`
}

/*
type UserManager interface {
	CreateUser()
	EnterUser()
	ChangeUserData()
}
*/
