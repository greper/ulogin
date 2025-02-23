// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// User is the golang structure for table user.
type User struct {
	Id           int64  `json:"id"           orm:"id"             description:"主键ID"`   // 主键ID
	CreatedAt    int64  `json:"createdAt"    orm:"created_at"     description:"创建时间"`   // 创建时间
	UpdatedAt    int64  `json:"updatedAt"    orm:"updated_at"     description:"更新时间"`   // 更新时间
	DeletedAt    int64  `json:"deletedAt"    orm:"deleted_at"     description:"删除时间"`   // 删除时间
	AppId        int64  `json:"appId"        orm:"app_id"         description:""`       //
	Avatar       string `json:"avatar"       orm:"avatar"         description:"用户头像"`   // 用户头像
	Username     string `json:"username"     orm:"username"       description:"用户登录名"`  // 用户登录名
	Password     string `json:"password"     orm:"password"       description:"用户登录密码"` // 用户登录密码
	NickName     string `json:"nickName"     orm:"nick_name"      description:"用户昵称"`   // 用户昵称
	PhoneCode    string `json:"phoneCode"    orm:"phone_code"     description:"手机区号"`   // 手机区号
	Mobile       string `json:"mobile"       orm:"mobile"         description:"手机号"`    // 手机号
	Email        string `json:"email"        orm:"email"          description:"邮箱"`     // 邮箱
	Gender       string `json:"gender"       orm:"gender"         description:"性别"`     // 性别
	UserType     int64  `json:"userType"     orm:"user_type"      description:"用户类型"`   // 用户类型
	InviteCode   string `json:"inviteCode"   orm:"invite_code"    description:"邀请码"`    // 邀请码
	InviteUserId int64  `json:"inviteUserId" orm:"invite_user_id" description:"邀请人id"`  // 邀请人id
	OpenId       string `json:"openId"       orm:"open_id"        description:""`       //
}
