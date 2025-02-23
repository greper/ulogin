package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/greper/ulogin/internal/service"
)

type LoginReq struct {
	g.Meta   `path:"/auth/login" method:"post" tags:"Auth" summary:"Login"`
	Username string `v:"required|length:2,50" dc:"用户名"`
	Password string `v:"required|length:2,255" dc:"密码"`
}

type LoginRes struct {
	UserInfo   *service.UserInfo `json:"userInfo" dc:"用户信息"`
	AccessInfo *service.JwtToken `json:"accessInfo" dc:"访问令牌"`
}
