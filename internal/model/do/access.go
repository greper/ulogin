// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Access is the golang structure of table access for DAO operations like Where/Data.
type Access struct {
	g.Meta           `orm:"table:access, do:true"`
	Id               interface{} // 主键ID
	CreatedAt        interface{} // 创建时间
	UpdatedAt        interface{} // 更新时间
	DeletedAt        interface{} // 删除时间
	Key              interface{} // code
	Title            interface{} // 显示名称
	Type             interface{} // 类型
	Setting          interface{} // 配置
	EncryptedSetting interface{} // 加密配置
}
