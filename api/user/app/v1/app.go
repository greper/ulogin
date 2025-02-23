package v1

import (
	"github.com/greper/ulogin/api/vo"

	"github.com/gogf/gf/v2/frame/g"
)

/**
type App struct {
	Id          int64       `json:"id"          orm:"id"           description:"主键ID"`    // 主键ID
	CreatedAt   int64       `json:"createdAt"   orm:"created_at"   description:"创建时间"`    // 创建时间
	UpdatedAt   int64       `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`    // 更新时间
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"   description:"删除时间"`    // 删除时间
	UserId      int64       `json:"userId"      orm:"user_id"      description:"用户id"`    // 用户id
	Key         string      `json:"key"         orm:"key"          description:"AppKey"`  // AppKey
	Title       string      `json:"title"       orm:"title"        description:"标题"`      // 标题
	SubjectType int64       `json:"subjectType" orm:"subject_type" description:"主体类型"`    // 主体类型
	PrivateKey  string      `json:"privateKey"  orm:"private_key"  description:"私钥"`      // 私钥
	PublicKey   string      `json:"publicKey"   orm:"public_key"   description:"公钥"`      // 公钥
	TokenSecret string      `json:"tokenSecret" orm:"token_secret" description:"token密钥"` // token密钥
}
*/

type CreateReq struct {
	g.Meta      `path:"/app/add" method:"post" tags:"App" summary:"Create app"`
	Title       string `v:"required|length:2,50" dc:"应用标题"`
	UserId      int64  `v:"required" dc:"用户ID"`
	SubjectType int64  `v:"required" dc:"主体类型"`
}
type CreateRes struct {
	Id          int64  `json:"id" dc:"应用ID"`
	Key         string `json:"key" dc:"AppKey"`
	Title       string `json:"title" dc:"应用标题"`
	SubjectType int64  `json:"subjectType" dc:"主体类型"`
	PublicKey   string `json:"publicKey" dc:"公钥"`
}

type UpdateReq struct {
	g.Meta `path:"/app/update" method:"post" tags:"App" summary:"Update app"`
	Id     int64   `v:"required" dc:"应用ID"`
	Title  *string `v:"length:2,50" dc:"应用标题"`
}
type UpdateRes struct{}

type DeleteReq struct {
	g.Meta `path:"/app/delete/{id}" method:"post" tags:"App" summary:"Delete app"`
	Id     int64 `v:"required" dc:"应用ID"`
}
type DeleteRes struct{}

type InfoReq struct {
	g.Meta `path:"/app/info/{id}" method:"post" tags:"App" summary:"Get one app"`
	Id     int64 `v:"required" dc:"应用ID"`
}
type InfoRes struct {
	Id          *int64  `json:"id" dc:"应用ID"`
	Title       *string `json:"title" dc:"应用标题"`
	Key         *string `json:"key" dc:"AppKey"`
	SubjectType *string `json:"subjectType" dc:"主体类型"`
	PublicKey   *string `json:"publicKey" dc:"公钥"`
}

type ListReq struct {
	g.Meta `path:"/app/list" method:"post" tags:"App" summary:"Get apps"`
	Query  *InfoRes    `dc:"查询条件"`
	Order  *vo.OrderBy `dc:"排序"`
}

type ListRes struct {
	List []*InfoRes `json:"list" dc:"数据列表"`
}

type PageReq struct {
	g.Meta `path:"/app/page" method:"post" tags:"App" summary:"Get app pages"`
	Query  *InfoRes `dc:"查询条件"`
	Page   *vo.Page `dc:"分页"`
	Order  *vo.OrderBy
}

type PageRes struct {
	g.Meta `path:"/app/page" method:"post" tags:"App" summary:"Get app pages"`
	List   []*InfoRes `json:"list" dc:"数据列表"`
	Page   *vo.Page   `dc:"分页"`
}
