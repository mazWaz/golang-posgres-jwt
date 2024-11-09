package profile

import (
	"go-clean/middlewares"
	"go-clean/modules/user"
)

var ValidateQueryAddress = middlewares.Validator{
	Param: &struct {
		Id string `uri:"id" validate:"required,gte=0"`
	}{},
}

var ValidateCreateAddress = middlewares.Validator{
	Body: &user.RequestUpdateAddress{},
}

var ValidateUpdateAddress = middlewares.Validator{
	Param: &struct {
		Id string `uri:"id" validate:"required,gte=0"`
	}{},
	Body: &user.RequestUpdateUser{},
}

var ValidateDeleteAddress = middlewares.Validator{
	Param: &struct {
		Id string `uri:"id" validate:"required,gte=0"`
	}{},
}
