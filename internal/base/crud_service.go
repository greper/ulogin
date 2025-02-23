package base

import (
	"context"
	"github.com/greper/ulogin/api/vo"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type CrudService[T any, E any] struct {
	Table string
}

// Create 创建记录
func (s *CrudService[T, E]) Create(ctx context.Context, data *T) (id int64, err error) {
	return g.DB().Model(s.Table).Ctx(ctx).Data(data).InsertAndGetId()
}

// Delete 删除记录
func (s *CrudService[T, E]) Delete(ctx context.Context, id interface{}) (err error) {
	_, err = g.DB().Model(s.Table).Ctx(ctx).Where("id", id).Delete()
	return
}

// Update 更新记录
func (s *CrudService[T, E]) Update(ctx context.Context, id interface{}, data *T) error {
	_, err := g.DB().Model(s.Table).Ctx(ctx).Data(data).Where("id", id).Update()
	return err
}

// Get 获取单条记录
func (s *CrudService[T, E]) Get(ctx context.Context, id interface{}) (result *E, err error) {
	err = g.DB().Model(s.Table).Ctx(ctx).Where("id", id).Scan(&result)
	return
}

// 查找单条记录
func (s *CrudService[T, E]) FindOne(ctx context.Context, Query *T) (result *E, err error) {
	list, err := s.GetList(ctx, &vo.ListReq[T]{Query: Query})
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	return nil, nil
}

// GetList 获取列表
func (s *CrudService[T, E]) GetList(ctx context.Context, listReq *vo.ListReq[T]) (list []*E, err error) {
	m := g.DB().Model(s.Table).Ctx(ctx)
	if listReq.Query != nil {
		m = m.Data(listReq.Query)
	}
	if listReq.OrderBy != nil && listReq.OrderBy.Prop != nil {
		if listReq.OrderBy.Asc != nil && *listReq.OrderBy.Asc {
			m = m.OrderAsc(*listReq.OrderBy.Prop)
		} else {
			m = m.OrderDesc(*listReq.OrderBy.Prop)
		}
	}
	err = m.Scan(&list)
	return
}

// GetPage 获取分页列表
func (s *CrudService[T, E]) GetPage(ctx context.Context, pageReq *vo.PageReq[T], customWhere func(*gdb.Model)) (list []*E, total int, err error) {
	m := g.DB().Model(s.Table).Ctx(ctx)

	m.Data(pageReq.Query)

	if customWhere != nil {
		customWhere(m)
	}
	if pageReq.OrderBy != nil && pageReq.OrderBy.Prop != nil {
		if pageReq.OrderBy.Asc != nil && *pageReq.OrderBy.Asc {
			m = m.OrderAsc(*pageReq.OrderBy.Prop)
		} else {
			m = m.OrderDesc(*pageReq.OrderBy.Prop)
		}
	}

	total, err = m.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	var page = pageReq.Page
	err = m.Page(int(*page.PageNo), int(*page.PageSize)).Scan(&list)
	return
}
