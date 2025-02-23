// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserRole is the golang structure of table user_role for DAO operations like Where/Data.
type UserRole struct {
	g.Meta `orm:"table:user_role, do:true"`
	UserId interface{} // 主键ID
	RoleId interface{} // 主键ID
}
