// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// LoginLog is the golang structure for table login_log.
type LoginLog struct {
	Id        int64       `json:"id"        orm:"id"         description:""` //
	UserId    int64       `json:"userId"    orm:"user_id"    description:""` //
	LoginTime *gtime.Time `json:"loginTime" orm:"login_time" description:""` //
	AppId     int64       `json:"appId"     orm:"app_id"     description:""` //
	ClientId  int64       `json:"clientId"  orm:"client_id"  description:""` //
	LoginIp   string      `json:"loginIp"   orm:"login_ip"   description:""` //
}
