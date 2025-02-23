package vo

type Page struct {
	PageNo   *int `json:"page" dc:"页码"`
	PageSize *int `json:"size" dc:"每页条数"`
	Total    *int `json:"total" dc:"总条数"`
}

type OrderBy struct {
	Prop *string `json:"order" dc:"排序"`
	By   *string `json:"by" dc:"排序方式"`
	Asc  *bool   `json:"asc" dc:"是否升序"`
}

type PageRes[T any] struct {
	Page *Page `json:"page" dc:"分页信息"`
	List []*T  `json:"list" dc:"数据列表"`
}

type PageReq[T any] struct {
	Page    *Page    `json:"page" dc:"分页信息"`
	Query   *T       `json:"query" dc:"查询条件"`
	OrderBy *OrderBy `json:"orderBy" dc:"排序"`
}

type ListReq[T any] struct {
	Query   *T       `json:"query" dc:"查询条件"`
	OrderBy *OrderBy `json:"orderBy" dc:"排序"`
}
