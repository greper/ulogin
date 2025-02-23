// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// OauthBound is the golang structure of table oauth_bound for DAO operations like Where/Data.
type OauthBound struct {
	g.Meta    `orm:"table:oauth_bound, do:true"`
	Id        interface{} // 主键ID
	CreatedAt interface{} // 创建时间
	UpdatedAt interface{} // 更新时间
	DeletedAt interface{} // 删除时间
	UserId    interface{} // 用户id
	Provider  interface{} // 第三方类型
	OpenId    interface{} // 第三方id
}
