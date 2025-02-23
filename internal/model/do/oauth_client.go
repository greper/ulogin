// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// OauthClient is the golang structure of table oauth_client for DAO operations like Where/Data.
type OauthClient struct {
	g.Meta       `orm:"table:oauth_client, do:true"`
	Id           interface{} // 主键ID
	CreatedAt    interface{} // 创建时间
	UpdatedAt    interface{} // 更新时间
	DeletedAt    interface{} // 删除时间
	AppId        interface{} //
	Key          interface{} // code
	Title        interface{} // 显示名称
	Type         interface{} // 类型
	AccessId     interface{} // appid
	AccessSecret interface{} // 密钥
}
