// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Role is the golang structure of table role for DAO operations like Where/Data.
type Role struct {
	g.Meta    `orm:"table:role, do:true"`
	Id        interface{} // 主键ID
	CreatedAt interface{} // 创建时间
	UpdatedAt interface{} // 更新时间
	DeletedAt interface{} // 删除时间
	Code      interface{} // 角色名
	Name      interface{} // 显示名称
}
