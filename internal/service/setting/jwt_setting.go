package setting

import "github.com/gogf/gf/v2/frame/g"

type JwtSetting struct {
	g.Meta       `key:"jwt_setting" title:"Jwt设置"`
	JwtSecret    string `json:"jwtSecret" `
	TokenExpires int64  `json:"tokenExpires" default:"2592000"`
}
