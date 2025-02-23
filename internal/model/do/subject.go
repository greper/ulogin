// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Subject is the golang structure of table subject for DAO operations like Where/Data.
type Subject struct {
	g.Meta         `orm:"table:subject, do:true"`
	Id             interface{} // 主键ID
	CreatedAt      interface{} // 创建时间
	UpdatedAt      interface{} // 更新时间
	DeletedAt      *gtime.Time // 删除时间
	UserId         interface{} // 绑定用户id
	AppId          interface{} // appId
	AppKey         interface{} // appKey
	SubjectId      interface{} // 主体Id
	SubjectType    interface{} // 主体类型
	Level          interface{} // 会员等级,default:1
	ActiveAt       interface{} // 激活时间
	ExpiresAt      interface{} // 过期时间
	Valid          interface{} // 状态
	License        interface{} // 授权文件
	LicenseVersion interface{} // 授权文件版本
	Remark         interface{} // 备注
	LastCode       interface{} // 激活码
	ActiveCount    interface{} // 激活次数
	Secret         interface{} // 密钥
	VipType        interface{} // 会员类型
	AppOwnerId     interface{} // app用户id
	Url            interface{} // 站点URL
	InstallAt      interface{} // 安装时间
	TrialUsed      interface{} // 试用已发放
}
