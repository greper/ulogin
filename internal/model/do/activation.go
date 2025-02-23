// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Activation is the golang structure of table activation for DAO operations like Where/Data.
type Activation struct {
	g.Meta     `orm:"table:activation, do:true"`
	Id         interface{} // 主键ID
	CreatedAt  interface{} // 创建时间
	UpdatedAt  interface{} // 更新时间
	DeletedAt  *gtime.Time // 删除时间
	Code       interface{} // 激活码
	UserId     interface{} // 用户id
	AppId      interface{} // 应用id
	AppKey     interface{} // AppKey
	Level      interface{} // 会员等级,default:1
	Duration   interface{} // 有效时长
	UseAt      interface{} // 使用时间
	StartAt    interface{} // 开始时间
	ExpiresAt  interface{} // 到期时间
	SubjectId  interface{} // 主体id
	Valid      interface{} // 状态
	Remark     interface{} // 备注
	Exported   interface{} // 是否已导出
	Version    interface{} // 版本
	VipType    interface{} // 会员类型
	AppOwnerId interface{} // 用户id
}
