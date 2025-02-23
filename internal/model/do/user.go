// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta       `orm:"table:user, do:true"`
	Id           interface{} // 主键ID
	CreatedAt    interface{} // 创建时间
	UpdatedAt    interface{} // 更新时间
	DeletedAt    interface{} // 删除时间
	AppId        interface{} //
	Avatar       interface{} // 用户头像
	Username     interface{} // 用户登录名
	Password     interface{} // 用户登录密码
	NickName     interface{} // 用户昵称
	PhoneCode    interface{} // 手机区号
	Mobile       interface{} // 手机号
	Email        interface{} // 邮箱
	Gender       interface{} // 性别
	UserType     interface{} // 用户类型
	InviteCode   interface{} // 邀请码
	InviteUserId interface{} // 邀请人id
	OpenId       interface{} //
}
