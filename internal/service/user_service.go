package service

import (
	"github.com/greper/ulogin/internal/base"
	"github.com/greper/ulogin/internal/model/do"
	"github.com/greper/ulogin/internal/model/entity"
)

type UserService struct {
	base.CrudService[do.User, entity.User]
}

var UserSvc = &UserService{
	base.CrudService[do.User, entity.User]{
		Table: "user",
	},
}
