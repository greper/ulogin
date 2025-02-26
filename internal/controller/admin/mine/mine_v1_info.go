package mine

import (
	"context"
	"github.com/greper/ulogin/api/admin/mine/v1"
	"github.com/greper/ulogin/internal/service"
)

func (c *ControllerV1) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {

	loginUser := service.GetLoginUser(ctx)

	get, err := service.UserSvc.Get(ctx, loginUser.UserId)
	if err != nil {
		return nil, err
	}
	get.Password = ""
	return &v1.InfoRes{
		UserInfo: &service.UserInfo{
			Id:       get.Id,
			Username: get.Username,
			Avatar:   get.Avatar,
			NickName: get.NickName,
		},
	}, nil
}
