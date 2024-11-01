package profile

import "go-clean/middlewares"

var ValidateQueryAddress = middlewares.Validator{
	Param: &struct {
		Id string `uri:"id" validate:"required,gte=0"`
	}{},
}

var ValidateCreateAddress = middlewares.Validator{
	Body: &RequestCreateAddress{},
}

var ValidateUpdateAddress = middlewares.Validator{
	Param: &struct {
		Id string `uri:"id" validate:"required,gte=0"`
	}{},
	Body: &RequestUpdateAddress{},
}

var ValidateDeleteAddress = middlewares.Validator{
	Param: &struct {
		Id string `uri:"id" validate:"required,gte=0"`
	}{},
}
