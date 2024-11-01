package user

import "go-clean/middlewares"

var ValidateQueryUser = middlewares.Validator{
	Query: &RequestQueryUser{},
}

var ValidateCreateUser = middlewares.Validator{
	Body: &RequestCreateUser{},
}

var ValidateUpdateUser = middlewares.Validator{
	Param: &struct {
		Id string `uri:"id" validate:"required,gte=0"`
	}{},
	Body: &RequestUpdateUser{},
}

var ValidateDeleteUser = middlewares.Validator{
	Param: &struct {
		Id string `uri:"id" validate:"required,gte=0"`
	}{},
}
