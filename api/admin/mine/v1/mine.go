package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/greper/ulogin/internal/service"
)

type InfoReq struct {
	g.Meta `path:"/mine/info" method:"post" tags:"Mine" summary:"Info"`
}

type InfoRes struct {
	UserInfo *service.UserInfo `json:"userInfo" dc:"用户信息"`
}
