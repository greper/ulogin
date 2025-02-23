// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysSetting is the golang structure of table sys_setting for DAO operations like Where/Data.
type SysSetting struct {
	g.Meta    `orm:"table:sys_setting, do:true"`
	Id        interface{} // 主键ID
	CreatedAt interface{} // 创建时间
	UpdatedAt interface{} // 更新时间
	DeletedAt interface{} // 删除时间
	Key       interface{} // 应用code
	Title     interface{} // 显示名称
	Setting   interface{} // Logo
}
