// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Access is the golang structure for table access.
type Access struct {
	Id               int64  `json:"id"               orm:"id"                description:"主键ID"` // 主键ID
	CreatedAt        int64  `json:"createdAt"        orm:"created_at"        description:"创建时间"` // 创建时间
	UpdatedAt        int64  `json:"updatedAt"        orm:"updated_at"        description:"更新时间"` // 更新时间
	DeletedAt        int64  `json:"deletedAt"        orm:"deleted_at"        description:"删除时间"` // 删除时间
	Key              string `json:"key"              orm:"key"               description:"code"` // code
	Title            string `json:"title"            orm:"title"             description:"显示名称"` // 显示名称
	Type             string `json:"type"             orm:"type"              description:"类型"`   // 类型
	Setting          string `json:"setting"          orm:"setting"           description:"配置"`   // 配置
	EncryptedSetting string `json:"encryptedSetting" orm:"encrypted_setting" description:"加密配置"` // 加密配置
}
