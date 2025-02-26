package boot

import "context"

func Init(ctx context.Context) {
	//数据库自动升级
	DbUp(ctx)

	//初始化系统配置
	SysSettingInit(ctx)
}
