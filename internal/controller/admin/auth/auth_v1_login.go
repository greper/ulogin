package auth

import (
	"context"
	"github.com/greper/ulogin/internal/service"

	"github.com/greper/ulogin/api/admin/auth/v1"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {

	loginRes, err := service.AuthSvc.LoginByPassword(ctx, &service.LoginByPasswordIn{
		Username: req.Username,
		Password: req.Password,
		UserType: 1,
	})
	if err != nil {
		return nil, err
	}

	return &v1.LoginRes{
		UserInfo:   loginRes.UserInfo,
		AccessInfo: loginRes.AccessInfo,
	}, nil

}
