package boot

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/greper/ulogin/internal/service"
	"github.com/greper/ulogin/internal/service/setting"
	"github.com/greper/ulogin/utility"
)

func SysSettingInit(ctx context.Context) {
	jwtSetting := &setting.JwtSetting{}
	err := service.SysSettingSvc.GetSetting(ctx, &jwtSetting)
	if err != nil {
		glog.Fatal(ctx, err)
	}

	if jwtSetting.JwtSecret == "" {
		jwtSetting.JwtSecret = utility.GenSecret()
		jwtSetting.TokenExpires = 2592000
	}

	err = service.SysSettingSvc.SaveSetting(ctx, jwtSetting)
	if err != nil {
		glog.Fatal(ctx, err)
	}
}
