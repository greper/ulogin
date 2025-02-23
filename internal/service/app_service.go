package service

import (
	"github.com/greper/ulogin/internal/base"
	"github.com/greper/ulogin/internal/model/do"
	"github.com/greper/ulogin/internal/model/entity"
)

type AppService struct {
	base.CrudService[do.App, entity.App]
}

var AppSvc = &AppService{
	base.CrudService[do.App, entity.App]{
		Table: "app",
	},
}
