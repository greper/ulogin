// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Permission is the golang structure of table permission for DAO operations like Where/Data.
type Permission struct {
	g.Meta        `orm:"table:permission, do:true"`
	Id            interface{} // 主键ID
	CreatedAt     interface{} // 创建时间
	UpdatedAt     interface{} // 更新时间
	DeletedAt     interface{} // 删除时间
	ApplicationId interface{} // 应用id
	ParentId      interface{} // 父权限id
	Title         interface{} // 资源标题
	Code          interface{} // 权限代码
	Path          interface{} // 权限url
}
