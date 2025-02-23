// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// App is the golang structure of table app for DAO operations like Where/Data.
type App struct {
	g.Meta    `orm:"table:app, do:true"`
	Id        interface{} // 主键ID
	CreatedAt interface{} // 创建时间
	UpdatedAt interface{} // 更新时间
	DeletedAt interface{} // 删除时间
	Logo      interface{} // Logo
	Key       interface{} // 应用code
	Title     interface{} // 显示名称
}
