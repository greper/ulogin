// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// OauthUser is the golang structure of table oauth_user for DAO operations like Where/Data.
type OauthUser struct {
	g.Meta    `orm:"table:oauth_user, do:true"`
	Id        interface{} // 主键ID
	CreatedAt interface{} // 创建时间
	UpdatedAt interface{} // 更新时间
	DeletedAt interface{} // 删除时间
	AppId     interface{} //
	UserId    interface{} // 用户id
	Provider  interface{} // 第三方类型
	OpenId    interface{} // 第三方id
}
