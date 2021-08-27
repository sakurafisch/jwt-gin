package service

type LoginService interface {
	LoginUser(email string, password string) bool
}

type loginInformation struct {
	email    string
	password string
}

func (this *loginInformation) LoginUser(email string, password string) bool {
	return this.email == email && this.password == password
}

func StaticLoginService() LoginService {
	return &loginInformation{
		email:    "test@jwt-go.com",
		password: "testing",
	}
}
