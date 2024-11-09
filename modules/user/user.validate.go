package user

import "go-clean/middlewares"

var ValidateQueryUser = middlewares.Validator{
	Query: &RequestQueryUserByAdmin{},
}

var ValidateCreateUserAdmin = middlewares.Validator{
	Body: &RequestCreateUserByAdmin{},
}

var ValidateCreateEmailByAdmin = middlewares.Validator{
	Body: &RequestCreateUserByAdmin{},
}

var ValidateCreateUser = middlewares.Validator{
	Body: &RequestCreateUser{},
}

var ValidateUpdateUser = middlewares.Validator{
	Param: &struct {
		Id string `uri:"id" validate:"required,gte=0"`
	}{},
	Body: &RequestUpdateUserByAdmin{},
}

var ValidateDeleteUser = middlewares.Validator{
	Param: &struct {
		Id string `uri:"id" validate:"required,gte=0"`
	}{},
}
