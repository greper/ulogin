// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// OauthUser is the golang structure for table oauth_user.
type OauthUser struct {
	Id        int64  `json:"id"        orm:"id"         description:"主键ID"`  // 主键ID
	CreatedAt int64  `json:"createdAt" orm:"created_at" description:"创建时间"`  // 创建时间
	UpdatedAt int64  `json:"updatedAt" orm:"updated_at" description:"更新时间"`  // 更新时间
	DeletedAt int64  `json:"deletedAt" orm:"deleted_at" description:"删除时间"`  // 删除时间
	AppId     int64  `json:"appId"     orm:"app_id"     description:""`      //
	UserId    int64  `json:"userId"    orm:"user_id"    description:"用户id"`  // 用户id
	Provider  string `json:"provider"  orm:"provider"   description:"第三方类型"` // 第三方类型
	OpenId    string `json:"openId"    orm:"open_id"    description:"第三方id"` // 第三方id
}
