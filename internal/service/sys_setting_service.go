package service

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/greper/ulogin/internal/base"
	"github.com/greper/ulogin/internal/dao"
	"github.com/greper/ulogin/internal/model/do"
	"github.com/greper/ulogin/internal/model/entity"
	"reflect"
	"strconv"
)

type SysSettingService struct {
	base.CrudService[do.SysSetting, entity.SysSetting]
}

var SysSettingSvc = &SysSettingService{
	base.CrudService[do.SysSetting, entity.SysSetting]{
		Table: "sys_setting",
	},
}

func (s *SysSettingService) GetByKey(ctx context.Context, key string) (result *entity.SysSetting, err error) {

	found, err := s.FindOne(ctx, &do.SysSetting{Key: key})
	if err != nil {
		return nil, err
	}
	if found != nil {
		return found, nil
	}
	return nil, nil
}

func (s *SysSettingService) GetSetting(ctx context.Context, model interface{}) (err error) {

	key := gmeta.Get(model, "key").String()

	found, err := s.GetByKey(ctx, key)
	if err != nil {
		return err
	}
	if found == nil {
		/**
		type JwtSetting struct {
			g.Meta       `key:"jwt_setting"`
			JwtSecret    string `json:"jwtSecret" `
			TokenExpires int64  `json:"tokenExpires" default:"2592000"`
		}
		*/
		//初始化一个 , 按照model TokenExpires类似的 default 配置 注入默认值

		//首先获取 TokenExpires的默认值配置

		typeof := reflect.TypeOf(model)
		valueof := reflect.ValueOf(model)
		for i := 0; i < typeof.Elem().NumField(); i++ {
			if valueof.Elem().Field(i).IsZero() {
				def := typeof.Elem().Field(i).Tag.Get("default")
				if def != "" {
					switch typeof.Elem().Field(i).Type.String() {
					case "int":
						result, _ := strconv.Atoi(def)
						valueof.Elem().Field(i).SetInt(int64(result))
					case "uint":
						result, _ := strconv.ParseUint(def, 10, 64)
						valueof.Elem().Field(i).SetUint(result)
					case "string":
						valueof.Elem().Field(i).SetString(def)
					}
				}
			}
		}
		//model转json字符串
		settingStr, _ := json.Marshal(model)

		dao.SysSetting.DB()

		title := gmeta.Get(model, "title").String()

		save := do.SysSetting{
			Key:     key,
			Title:   title,
			Setting: string(settingStr),
		}

		id, err := s.Create(ctx, &save)
		if err != nil {
			return err
		}

		save.Id = id

		found = &entity.SysSetting{}
		err = gconv.Struct(save, found)
		if err != nil {
			return err
		}
	}
	setting := found.Setting
	err = json.Unmarshal([]byte(setting), model)
	return err
}
