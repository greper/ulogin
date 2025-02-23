// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// OauthClient is the golang structure for table oauth_client.
type OauthClient struct {
	Id           int64  `json:"id"           orm:"id"            description:"主键ID"`  // 主键ID
	CreatedAt    int64  `json:"createdAt"    orm:"created_at"    description:"创建时间"`  // 创建时间
	UpdatedAt    int64  `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`  // 更新时间
	DeletedAt    int64  `json:"deletedAt"    orm:"deleted_at"    description:"删除时间"`  // 删除时间
	AppId        int64  `json:"appId"        orm:"app_id"        description:""`      //
	Key          string `json:"key"          orm:"key"           description:"code"`  // code
	Title        string `json:"title"        orm:"title"         description:"显示名称"`  // 显示名称
	Type         string `json:"type"         orm:"type"          description:"类型"`    // 类型
	AccessId     string `json:"accessId"     orm:"access_id"     description:"appid"` // appid
	AccessSecret string `json:"accessSecret" orm:"access_secret" description:"密钥"`    // 密钥
}
