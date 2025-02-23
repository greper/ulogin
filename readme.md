# 激活码平台接口
替代之前的 `handsfree/platform/api/activation` 模块

## 功能
- 开发者可以注册开发者账号，创建APP
- 为APP生成激活码，导出发放给用户
- 用户使用激活码在真正的应用中请求激活
- 激活后获得License， 应用本地可以校验license， 并且还需要定期远程校验license有效性
- 开发者可以在后台查看激活情况
- 激活码可以自定义类型，比如专业版，商业版之类
- 激活码可以设置有效期，比如1年，2年之类
- 用户邀请其他用户来购买激活码并激活时，可以获得奖励积分，积分可以延长专业版、和商业版的有效期

## 技术栈
- 前端：vue3+ antdv
- 后端：GoFrame + Gorm + MySQL
- 认证：JWT


## 命令帮助

根据数据库表生成dao代码
gf gen dao

根据接口定义生成controller代码
gf gen ctrl
