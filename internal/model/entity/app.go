// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// App is the golang structure for table app.
type App struct {
	Id        int64  `json:"id"        orm:"id"         description:"主键ID"`   // 主键ID
	CreatedAt int64  `json:"createdAt" orm:"created_at" description:"创建时间"`   // 创建时间
	UpdatedAt int64  `json:"updatedAt" orm:"updated_at" description:"更新时间"`   // 更新时间
	DeletedAt int64  `json:"deletedAt" orm:"deleted_at" description:"删除时间"`   // 删除时间
	Logo      string `json:"logo"      orm:"logo"       description:"Logo"`   // Logo
	Key       string `json:"key"       orm:"key"        description:"应用code"` // 应用code
	Title     string `json:"title"     orm:"title"      description:"显示名称"`   // 显示名称
}
