// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RolePermission is the golang structure of table role_permission for DAO operations like Where/Data.
type RolePermission struct {
	g.Meta       `orm:"table:role_permission, do:true"`
	RoleId       interface{} // 主键ID
	PermissionId interface{} // 主键ID
}
