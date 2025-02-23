# Change Log

All notable changes to this project will be documented in this file.
See [Conventional Commits](https://conventionalcommits.org) for commit guidelines.

## [1.30.5](https://github.com/certd/certd/compare/v1.30.4...v1.30.5) (2025-02-14)

**Note:** Version bump only for package @certd/ui-client

## [1.30.4](https://github.com/certd/certd/compare/v1.30.3...v1.30.4) (2025-02-14)

**Note:** Version bump only for package @certd/ui-client

## [1.30.3](https://github.com/certd/certd/compare/v1.30.2...v1.30.3) (2025-02-13)

**Note:** Version bump only for package @certd/ui-client

## [1.30.2](https://github.com/certd/certd/compare/v1.30.1...v1.30.2) (2025-02-09)

### Bug Fixes

* 当前置任务被删除时进行校验 ([c89686a](https://github.com/certd/certd/commit/c89686a2fda251484930f0ae715417b618c21690))

### Performance Improvements

* 上传自定义证书 ([75a38d9](https://github.com/certd/certd/commit/75a38d95f305b4271d9106babe7cffc1c89ae8f3))

## [1.30.1](https://github.com/certd/certd/compare/v1.30.0...v1.30.1) (2025-01-20)

### Performance Improvements

* 创建流水线时，默认成功时也发送通知 ([52ae690](https://github.com/certd/certd/commit/52ae6902d203ca56e0312692b50c55cb6ddd3e39))
* http方式校验，选择sftp时，支持修改文件访问权限比如777 ([15d6eaf](https://github.com/certd/certd/commit/15d6eaf5532ed25acd4f8d58c429353a2f44206c))

# [1.30.0](https://github.com/certd/certd/compare/v1.29.5...v1.30.0) (2025-01-19)

### Bug Fixes

* 修复查看任务日志偶发性无法自动滚动底部的bug ([7e482f7](https://github.com/certd/certd/commit/7e482f798c0142bce1866f84676cb40210f9638a))

### Performance Improvements

* 证书仓库 ([91e7f45](https://github.com/certd/certd/commit/91e7f45a1c5ea1e0ec0aa3236b80028f03a6d0aa))

## [1.29.5](https://github.com/certd/certd/compare/v1.29.4...v1.29.5) (2025-01-07)

### Bug Fixes

* 修复授权管理，点击了查看原文按钮后，无法修改值的bug ([85c99f7](https://github.com/certd/certd/commit/85c99f7f80761ac6efaf3255c03b933442db1686))

## [1.29.4](https://github.com/certd/certd/compare/v1.29.3...v1.29.4) (2025-01-06)

### Bug Fixes

* 修复站点监控域名校验无法通过的bug ([1cb4a53](https://github.com/certd/certd/commit/1cb4a539cc523721ffd4b22d40d0e3d2d68cd915))

## [1.29.3](https://github.com/certd/certd/compare/v1.29.2...v1.29.3) (2025-01-04)

### Performance Improvements

* 优化站点证书检查页面，检查增加3次重试 ([e6dd7cd](https://github.com/certd/certd/commit/e6dd7cd54a3e23897031b5df6e0c3cdc0545d35a))
* 支持http校验方式申请证书 ([405591c](https://github.com/certd/certd/commit/405591c5d08fa1a3b228ee3980199e7731cfec4a))
* http校验方式，支持七牛云oss、阿里云oss、腾讯云cos ([3f74d4d](https://github.com/certd/certd/commit/3f74d4d9e5f5d0e629b44cff1895b3f7a8fbcafc))

## [1.29.2](https://github.com/certd/certd/compare/v1.29.1...v1.29.2) (2024-12-25)

**Note:** Version bump only for package @certd/ui-client

## [1.29.1](https://github.com/certd/certd/compare/v1.29.0...v1.29.1) (2024-12-25)

### Bug Fixes

* 免费套餐支持购买 ([f5ec987](https://github.com/certd/certd/commit/f5ec9870fd6af1f0c9099852bbdb4d07813ccce8))
* 修复某处金额转换丢失精度的bug ([d2d6f12](https://github.com/certd/certd/commit/d2d6f12218cbe7bd55f4ae082b93084be85f0a7b))
* 修复新版本小红点显示错误问题 ([fe4786e](https://github.com/certd/certd/commit/fe4786e168afe03a5243dd67971476c348339809))

### Performance Improvements

* 用户创建证书流水线没有购买套餐或者超限时提前报错 ([472f06c](https://github.com/certd/certd/commit/472f06c2d190d0ae48e8b53c18bc278437656a1c))
* 优化插件名称显示 ([26adf7d](https://github.com/certd/certd/commit/26adf7d437e674385f26a8f92fded6521a620671))

# [1.29.0](https://github.com/certd/certd/compare/v1.28.4...v1.29.0) (2024-12-24)

### Bug Fixes

* 修复手机模式下，查询框被文字遮盖的bug ([040788c](https://github.com/certd/certd/commit/040788c793642c3bb2a3ede87fe30fcf3be471bd))
* 修复左侧菜单收起时无法展开子菜单的bug ([0056223](https://github.com/certd/certd/commit/005622307e612717a5408aa1484717ef03003a22))

### Features

* 基础版不再限制流水线数量 ([cb27d4b](https://github.com/certd/certd/commit/cb27d4b4906b2782eaceb0a95bbdc5d0534370d2))
* 套餐购买支持易支付、支付宝支付 ([faa28f8](https://github.com/certd/certd/commit/faa28f88f954cba4c1dd29125562e5acd2fd99af))
* 用户套餐，用户支付功能 ([a019956](https://github.com/certd/certd/commit/a019956698acaf2c4beb620b5ad8c18918ead6a1))
* 支持微信支付 ([45d6347](https://github.com/certd/certd/commit/45d6347f5b6199493b11aabdd74177f6dca2cea4))

### Performance Improvements

* 调整创建证书表单字段的顺序 ([d393521](https://github.com/certd/certd/commit/d3935219f2aa50d6662c5b5ebf7ee25ad696ab2b))
* 同一时间只允许一个套餐生效 ([8ebf95a](https://github.com/certd/certd/commit/8ebf95a222a900d1707716c7b1f3b39f8a6d8f94))
* 用户名支持修改 ([89c7f07](https://github.com/certd/certd/commit/89c7f070343e86453c84677ebe1669f9b266d871))
* 优化证书申请跳过的状态显示，成功通知现在在跳过时不会发送 ([67d762b](https://github.com/certd/certd/commit/67d762b6a520f1fa24719a124e5ae975a81f5f82))
* 站点证书监控通知发送，每天定时检查 ([bb4910f](https://github.com/certd/certd/commit/bb4910f4e57234e42b44505f4620ae7af66025c5))

## [1.28.4](https://github.com/certd/certd/compare/v1.28.3...v1.28.4) (2024-12-12)

**Note:** Version bump only for package @certd/ui-client

## [1.28.3](https://github.com/certd/certd/compare/v1.28.2...v1.28.3) (2024-12-12)

### Bug Fixes

* 修复授权被删除后，无法清空的bug ([b45977c](https://github.com/certd/certd/commit/b45977c29a29084c11e496bec3415eaaebafdd74))

### Performance Improvements

* 点击版本红点按钮，跳转到升级帮助页面 ([454fbda](https://github.com/certd/certd/commit/454fbda581bbe22abca5b91e5086ea9d9d58a020))
* 通知标题优化 ([ff083ce](https://github.com/certd/certd/commit/ff083ce6848a8bee3c8248e4b881086ae1517c28))

## [1.28.2](https://github.com/certd/certd/compare/v1.28.1...v1.28.2) (2024-12-09)

### Bug Fixes

* 修复创建流水线通知设置无效的bug ([498cf34](https://github.com/certd/certd/commit/498cf34999fddfa24ce088e2e678469fa669abb8))

## [1.28.1](https://github.com/certd/certd/compare/v1.28.0...v1.28.1) (2024-12-08)

### Bug Fixes

* 修复cname排查方法 nslookup命令显示黑色的问题 ([3dfeeec](https://github.com/certd/certd/commit/3dfeeec899d7d0d7292695ce410f78548e076c03))

### Performance Improvements

* 通知选择器优化 ([2c0cbdd](https://github.com/certd/certd/commit/2c0cbdd29ecb74cc939b2ae7ee86b8d40f70ba31))
* 新增七牛云插件分组 ([49e7dc5](https://github.com/certd/certd/commit/49e7dc56e1a95fbdea3e30cdeb945b48415b69e3))
* 新增server酱3通知 ([6aa4872](https://github.com/certd/certd/commit/6aa487269c9f6862e188b37a0d6c73f79c937d94))
* 支持邀请奖励 ([618ec93](https://github.com/certd/certd/commit/618ec937866b24ebcf8164db43acb1ed66a5b329))
* 支持易发云短信 ([94fa77f](https://github.com/certd/certd/commit/94fa77fcd2b9bea294fb05736c0d8cdc81f56103))
* favicon支持自定义 ([8b9c47d](https://github.com/certd/certd/commit/8b9c47daf194515006689a212ae9cf586bdf5993))

# [1.28.0](https://github.com/certd/certd/compare/v1.27.9...v1.28.0) (2024-11-30)

### Features

* 手机号登录、邮箱验证码注册 ([7b55337](https://github.com/certd/certd/commit/7b55337c5edb470cca7aa62201eda8d274784004))

### Performance Improvements

* 部署到IIS插件 ([1534f45](https://github.com/certd/certd/commit/1534f4523633265d219d7b3a249a9ea1af99c512))
* 流水线支持批量修改分组，批量删除 ([a847e66](https://github.com/certd/certd/commit/a847e66c4fc843b98f1520b2b8072d3586ce8b81))
* 首页新增修改密码提示 ([0772d3b](https://github.com/certd/certd/commit/0772d3b3fd24afdde4086d9f09ef19d037b431b4))
* 选项显示图标 ([aedc462](https://github.com/certd/certd/commit/aedc46213571a3bd93809b7af7fa17a08d546237))
* 优化证书申请成功通知发送方式 ([8002a56](https://github.com/certd/certd/commit/8002a56efc5998aa03db5711ae87f9eb4bc9e160))
* 支持短信验证码登录 ([387bcc5](https://github.com/certd/certd/commit/387bcc5fa418cdeea81a06da5e3f8cd6b43cd082))
* 自定义webhook显示详细的错误信息 ([3254afc](https://github.com/certd/certd/commit/3254afc75640eed3729d0fc02a818fefbe5c7fc3))

## [1.27.9](https://github.com/certd/certd/compare/v1.27.8...v1.27.9) (2024-11-26)

### Performance Improvements

* 通知支持自定义webhook、anpush、iyuu、server酱 ([cbccd9e](https://github.com/certd/certd/commit/cbccd9e3d0a4c24aba772af62734666d40b22c57))
* 通知支持vocechat、bark、telegram、discord、slack ([642f57f](https://github.com/certd/certd/commit/642f57ff6d7152a9e14f59c7fc0e32a6b1751fb7))

## [1.27.8](https://github.com/certd/certd/compare/v1.27.7...v1.27.8) (2024-11-25)

**Note:** Version bump only for package @certd/ui-client

## [1.27.7](https://github.com/certd/certd/compare/v1.27.6...v1.27.7) (2024-11-25)

### Performance Improvements

* 通知管理 ([d9a00ee](https://github.com/certd/certd/commit/d9a00eeaf72735ced67c59d7983d84e3c730064a))
* 通知渠道支持测试按钮 ([b54ae27](https://github.com/certd/certd/commit/b54ae272ebc2d31b32b049d44e2299a6be7f153c))
* 支持企业微信群聊机器人通知 ([b805a29](https://github.com/certd/certd/commit/b805a2925984144a31575b8aaa622f0c30d41b56))

## [1.27.6](https://github.com/certd/certd/compare/v1.27.5...v1.27.6) (2024-11-19)

### Performance Improvements

* 当步骤全部都禁用时，任务本身显示删除线 ([9ab9a6e](https://github.com/certd/certd/commit/9ab9a6e8b083e19793894f23e59f29c604ec98e5))

## [1.27.5](https://github.com/certd/certd/compare/v1.27.4...v1.27.5) (2024-11-18)

### Bug Fixes

* 修复1Panel面板本身证书更新导致判定执行失败的问题 ([2689e6d](https://github.com/certd/certd/commit/2689e6d6c03aba21da90d5d45232c6ba08696be1))
* 修复Cname情况下，无法使用DNS类型的bug ([26dad39](https://github.com/certd/certd/commit/26dad399d5768b3205da099ddc11809aef7d6224))

### Performance Improvements

* 日志查看自动滚动到底部 ([4a2f7eb](https://github.com/certd/certd/commit/4a2f7ebf87b7c027cebff7cb763f8f35f6d2aa36))
* 系统设置中的代理设置优化为可全局生效，环境变量中的https_proxy设置将无效 ([381a37f](https://github.com/certd/certd/commit/381a37fbaa6b61c887eda743897ae00afb825bdf))
* 新手导航在非编辑模式下不显示 ([18bfcc2](https://github.com/certd/certd/commit/18bfcc24ad0bde57bb04db8a4209861ec6b8ff1d))
* 专业版试用，无需绑定账号 ([c7c4318](https://github.com/certd/certd/commit/c7c4318c11b65a76089787aa58939832d338a232))

## [1.27.4](https://github.com/certd/certd/compare/v1.27.3...v1.27.4) (2024-11-14)

### Performance Improvements

* 公共cname服务支持关闭 ([f4ae512](https://github.com/certd/certd/commit/f4ae5125dc4cd97816976779cb3586b5ee78947e))

## [1.27.3](https://github.com/certd/certd/compare/v1.27.2...v1.27.3) (2024-11-13)

### Bug Fixes

* 修复邮件配置，忽略证书校验设置不生效的bug ([66a9690](https://github.com/certd/certd/commit/66a9690dc958732e1b3c672d965db502296446f9))

### Performance Improvements

* 修复站点个性化，浏览器标题没有生效的bug ([bcfac02](https://github.com/certd/certd/commit/bcfac02c96ceaf23d1a0b05b48d8047da933beaf))
* 优化上传到主机插 路径选择，根据证书格式显示 ([8c3f86c](https://github.com/certd/certd/commit/8c3f86c6909ed91f48bb2880e78834e22f6f6a29))
* ipv6支持 ([da6ac16](https://github.com/certd/certd/commit/da6ac1626b3574be2fabeeb18a1f10d60bdcbe49))

## [1.27.2](https://github.com/certd/certd/compare/v1.27.1...v1.27.2) (2024-11-08)

### Performance Improvements

* 优化流水线页面切换回来不丢失查询条件 ([4dcf6e8](https://github.com/certd/certd/commit/4dcf6e87bc5f7657ce8a56c5331e8723a0fee8ee))
* 支持公共cname服务 ([3c919ee](https://github.com/certd/certd/commit/3c919ee5d1aef5d26cf3620a7c49d920786bc941))
* 执行历史支持点击查看流水线详情 ([8968639](https://github.com/certd/certd/commit/89686399f90058835435b92872fc236fac990148))
* 专业版7天试用 ([c58250e](https://github.com/certd/certd/commit/c58250e1f065a9bd8b4e82acc1df754504c0010c))

## [1.27.1](https://github.com/certd/certd/compare/v1.27.0...v1.27.1) (2024-11-04)

### Bug Fixes

* 修复头像没有更新的bug ([9b4a31f](https://github.com/certd/certd/commit/9b4a31fa6a32b9cab2e22bd141cf96ca29120445))

### Performance Improvements

* 禁止页面缓存，点击tab页签可以刷新数据 ([7ad4b55](https://github.com/certd/certd/commit/7ad4b55ee000c1dd0747832b11107f32b0ffb889))
* 优化时间选择器，自动填写分钟和秒钟 ([396dc34](https://github.com/certd/certd/commit/396dc34a841c7d016b033736afdba8366fb2d211))
* cname 域名映射记录可读性优化 ([b1117ed](https://github.com/certd/certd/commit/b1117ed54a3ef015752999324ff72b821ef5e4b9))

# [1.27.0](https://github.com/certd/certd/compare/v1.26.16...v1.27.0) (2024-10-31)

### Bug Fixes

* 修复历史记录不能按名称查询的bug ([6113c38](https://github.com/certd/certd/commit/6113c388b7fc58b11ca19ff05cc1286d096c8d28))

### Features

* 首页全新改版 ([63ec5b5](https://github.com/certd/certd/commit/63ec5b5519c760a3330569c0da6dac157302a330))

### Performance Improvements

* 管理控制台数据统计 ([babd589](https://github.com/certd/certd/commit/babd5897ae013ff7c04ebfcbfac8a00d84dd627c))
* 增加向导 ([6d9ef26](https://github.com/certd/certd/commit/6d9ef26ecab71d752c2c55d75aed4fb5f6c05a39))

## [1.26.16](https://github.com/certd/certd/compare/v1.26.15...v1.26.16) (2024-10-30)

**Note:** Version bump only for package @certd/ui-client

## [1.26.15](https://github.com/certd/certd/compare/v1.26.14...v1.26.15) (2024-10-28)

### Bug Fixes

* 顶部菜单变...的bug ([6dabad7](https://github.com/certd/certd/commit/6dabad76baba96be0f8af36a3fbfb9f5182aecf1))

### Performance Improvements

* 默认证书更新时间设置为35天，增加腾讯云删除过期证书插件，可以避免腾讯云过期证书邮件 ([51b6fed](https://github.com/certd/certd/commit/51b6fed468eaa6f28ce4497ce303ace1a52abb96))
* 授权加密支持解密查看 ([5575c83](https://github.com/certd/certd/commit/5575c839705f6987ad2bdcd33256b0962c6a9c6a))
* 重置管理员密码同时启用管理员账户，避免之前禁用了，重置密码还是登录不进去 ([f92d918](https://github.com/certd/certd/commit/f92d918a1e28e29b794ad4754661ea760c18af46))

## [1.26.14](https://github.com/certd/certd/compare/v1.26.13...v1.26.14) (2024-10-26)

### Bug Fixes

* 修复启动时自签证书无法保存的bug ([526c484](https://github.com/certd/certd/commit/526c48450bcd37b3ccded9b448f17de8140bdc6e))

### Performance Improvements

* 顶部菜单自定义 ([54d136c](https://github.com/certd/certd/commit/54d136cc6ae122f7c891b7a5c7232fe5de8e5cb5))
* 禁用readonly用户 ([d10d42e](https://github.com/certd/certd/commit/d10d42e20619bb55a50d636b8867ff33db4e3b4b))
* 限制其他用户流水线数量 ([315e437](https://github.com/certd/certd/commit/315e43746baf01682737f82e41579237a48409af))
* 用户管理优化头像上传 ([661293c](https://github.com/certd/certd/commit/661293c189a3abf3cdc953b5225192372f57930d))

## [1.26.13](https://github.com/certd/certd/compare/v1.26.12...v1.26.13) (2024-10-26)

### Bug Fixes

* 修复对话框全屏按钮与关闭按钮重叠的bug ([95df56c](https://github.com/certd/certd/commit/95df56cc5ca5e3eb843cd17cb7078cde47729f1e))

### Performance Improvements

* 支持同时监听https端口，7002 ([d5a17f9](https://github.com/certd/certd/commit/d5a17f9e6afd63fda2df0981118480f25a1fac2e))

## [1.26.12](https://github.com/certd/certd/compare/v1.26.11...v1.26.12) (2024-10-25)

### Performance Improvements

* 部署到阿里云任意云资源，阿里云部署大杀器 ([4075be7](https://github.com/certd/certd/commit/4075be7849b140acb92bd8da8a9acbf4eef85180))
* 文件名特殊字符限制输入 ([c4164c6](https://github.com/certd/certd/commit/c4164c66e29f3ec799f98108a344806ca61e94ff))
* 新增部署到百度云CDN插件 ([f126f9f](https://github.com/certd/certd/commit/f126f9f932d37fa01fff1accc7bdd17d349f8db5))
* 优化cron选择器，增加下次触发时间显示 ([5b148b7](https://github.com/certd/certd/commit/5b148b7ed960ca6f7f5b733b2eadd56eeecbd4c2))
* 支持配置公共ZeroSSL授权 ([a90d1e6](https://github.com/certd/certd/commit/a90d1e68ee9cbc3705223457b8a86f071b150968))

## [1.26.11](https://github.com/certd/certd/compare/v1.26.10...v1.26.11) (2024-10-23)

### Bug Fixes

* 修复移动任务后出现空阶段的bug ([4ea3edd](https://github.com/certd/certd/commit/4ea3edd59e93ca4f5b2e43b20dd4ef33909caddb))
* 允许七牛云cdn插件输入.号开头的通配符域名 ([18ee87d](https://github.com/certd/certd/commit/18ee87daff6eafc2201b58e28d85aafd3cb7a5b9))

### Performance Improvements

* 优化日志颜色 ([1291e98](https://github.com/certd/certd/commit/1291e98e821c5b1810aab7f0aebe3f5f5cd44a20))
* 优化证书申请速度和成功率，反代地址优化，google基本可以稳定请求。增加请求重试。 ([41d9c3a](https://github.com/certd/certd/commit/41d9c3ac8398def541e65351cbe920d4a927182d))
* 优化pfx密码密码输入框，让浏览器不自动填写密码 ([ffeede3](https://github.com/certd/certd/commit/ffeede38afa70c5ff6f2015516bead23d2c4df87))

## [1.26.10](https://github.com/certd/certd/compare/v1.26.9...v1.26.10) (2024-10-20)

### Bug Fixes

* 修复cname服务普通用户access访问权限问题 ([c1e3e2e](https://github.com/certd/certd/commit/c1e3e2ee1f923ee5806479dd5f178c3286a01ae0))

## [1.26.9](https://github.com/certd/certd/compare/v1.26.8...v1.26.9) (2024-10-19)

### Bug Fixes

* 修复切换普通用户登录时，左侧菜单没有同步更新的bug ([12116a8](https://github.com/certd/certd/commit/12116a89f43cf8b98f16d2ea6073f6b72a643215))
* 修正邮箱设置跳转路由 ([17d8890](https://github.com/certd/certd/commit/17d88900a1f0e3af609b74597f5b1978230db32d))

### Performance Improvements

* 触发证书重新申请input变化对比规则优化，减少升级版本后触发申请证书的情况 ([c46a2a9](https://github.com/certd/certd/commit/c46a2a9a399c2a9a8bb59a48b9fb6e93227cce9b))
* 优化菜单 ([1f4f157](https://github.com/certd/certd/commit/1f4f15757de1015cf7563f7022599eef58cc93d7))
* 增加文档站 https://certd.docmirror.cn ([6e2ac1c](https://github.com/certd/certd/commit/6e2ac1c089f6ddccb396f1f2738509c05333e1bb))

## [1.26.8](https://github.com/certd/certd/compare/v1.26.7...v1.26.8) (2024-10-15)

### Performance Improvements

* 密钥备份 ([1c6028a](https://github.com/certd/certd/commit/1c6028abcf8849163462bb2f8441b6838357e09b))
* 证书直接查看 ([5dde5bd](https://github.com/certd/certd/commit/5dde5bd3f76db3959d411619d29bfb8064e3b307))
* sqlite数据库备份插件 ([77f1631](https://github.com/certd/certd/commit/77f163144f7dcfb0431475c55508fecfd6d969f8))

## [1.26.7](https://github.com/certd/certd/compare/v1.26.6...v1.26.7) (2024-10-14)

**Note:** Version bump only for package @certd/ui-client

## [1.26.6](https://github.com/certd/certd/compare/v1.26.5...v1.26.6) (2024-10-14)

**Note:** Version bump only for package @certd/ui-client

## [1.26.5](https://github.com/certd/certd/compare/v1.26.4...v1.26.5) (2024-10-14)

**Note:** Version bump only for package @certd/ui-client

## [1.26.4](https://github.com/certd/certd/compare/v1.26.3...v1.26.4) (2024-10-14)

### Performance Improvements

* [comm] 支持插件管理 ([e8b617b](https://github.com/certd/certd/commit/e8b617b80ce882dd63006f0cfc719a80a1cc6acc))
* 新增代理设置功能 ([273ab61](https://github.com/certd/certd/commit/273ab6139f5807f4d7fe865cc353b97f51b9a668))
* EAB授权支持绑定邮箱，支持公共EAB设置 ([07043af](https://github.com/certd/certd/commit/07043aff0ca7fd29c56dd3c363002cb15d78b464))

## [1.26.3](https://github.com/certd/certd/compare/v1.26.2...v1.26.3) (2024-10-12)

### Performance Improvements

* 优化系统设置加载时机 ([7396253](https://github.com/certd/certd/commit/73962536d5a4769902d760d005f3f879465addcc))

## [1.26.2](https://github.com/certd/certd/compare/v1.26.1...v1.26.2) (2024-10-11)

### Performance Improvements

* 邮箱设置改为系统设置，普通用户无需配置发件邮箱 ([4244569](https://github.com/certd/certd/commit/42445692117184a3293e63bef84a74cbb5984b0e))

## [1.26.1](https://github.com/certd/certd/compare/v1.26.0...v1.26.1) (2024-10-10)

**Note:** Version bump only for package @certd/ui-client

# [1.26.0](https://github.com/certd/certd/compare/v1.25.9...v1.26.0) (2024-10-10)

### Bug Fixes

* 修复历史记录根据流水线名称查询报错的bug ([ce9a986](https://github.com/certd/certd/commit/ce9a9862f122fce2186e7727eaa4b251b59e6032))

### Features

* 域名验证方法支持CNAME间接方式，此方式支持所有域名注册商，且无需提供Access授权，但是需要手动添加cname解析 ([f3d3508](https://github.com/certd/certd/commit/f3d35084ed44f9f33845f7045e520be5c27eed93))
* 站点个性化设置 ([11a9fe9](https://github.com/certd/certd/commit/11a9fe9014d96cba929e5a066e78f2af7ae59d14))

### Performance Improvements

* 并行任务名称改成添加任务，取消并行，可以在同一个阶段获取上一个task的输出 ([c5e5877](https://github.com/certd/certd/commit/c5e58770d1c5edc19c6f9ea1618f44b68e091f35))
* 调整静态资源到static目录 ([0584b36](https://github.com/certd/certd/commit/0584b3672b40f9042a2ed87e5627022606d046cd))
* 调整全部静态资源到static目录 ([a218890](https://github.com/certd/certd/commit/a21889080d6c7ffdf0af526a3a21f0b2d1c77288))
* 检查cname是否正确配置 ([b5d8935](https://github.com/certd/certd/commit/b5d8935159374fbe7fc7d4c48ae0ed9396861bdd))
* 优化宝塔网站部署插件远程获取数据的提示 ([2a3ca9f](https://github.com/certd/certd/commit/2a3ca9f552d96594ec6690a1c4c91f598451b9a1))
* 优化缩短首页缓存时间 ([49395e8](https://github.com/certd/certd/commit/49395e8cb65f4b30c0145329ed5de48be4ef3842))
* 域名输入增加校验提示，避免输入错误的域名 ([0c8e83e](https://github.com/certd/certd/commit/0c8e83e1254a9ce4d5a4e7888eb1710394a4b77c))
* cname校验配置增加未校验通过提示 ([77cc3c4](https://github.com/certd/certd/commit/77cc3c4a5cbd81f8233a8e0bb33fab0621c0905f))
* google eab授权支持自动获取，不过要配置代理 ([592791d](https://github.com/certd/certd/commit/592791d1356fc252fbb70d7f168567aee9585507))

## [1.25.9](https://github.com/certd/certd/compare/v1.25.8...v1.25.9) (2024-10-01)

### Bug Fixes

* 修复西部数码账户级别apikey不可用的bug ([f8f3e8b](https://github.com/certd/certd/commit/f8f3e8b43fd5d815887bcb53b95f46dc96424b79))

## [1.25.8](https://github.com/certd/certd/compare/v1.25.7...v1.25.8) (2024-09-30)

### Performance Improvements

* 群晖获取deviceid优化 ([8d42273](https://github.com/certd/certd/commit/8d4227366548eb70f6bc04303829e6933168f906))

## [1.25.7](https://github.com/certd/certd/compare/v1.25.6...v1.25.7) (2024-09-29)

**Note:** Version bump only for package @certd/ui-client

## [1.25.6](https://github.com/certd/certd/compare/v1.25.5...v1.25.6) (2024-09-29)

### Performance Improvements

* 部署支持1Panel ([d047234](https://github.com/certd/certd/commit/d047234d98d31504f2e5a472b66e1b75806af26e))
* 增加使用教程 ([9d9c021](https://github.com/certd/certd/commit/9d9c0218195af5b9896cce7109b26a433480571d))

## [1.25.5](https://github.com/certd/certd/compare/v1.25.4...v1.25.5) (2024-09-26)

**Note:** Version bump only for package @certd/ui-client

## [1.25.4](https://github.com/certd/certd/compare/v1.25.3...v1.25.4) (2024-09-25)

**Note:** Version bump only for package @certd/ui-client

## [1.25.3](https://github.com/certd/certd/compare/v1.25.2...v1.25.3) (2024-09-24)

**Note:** Version bump only for package @certd/ui-client

## [1.25.2](https://github.com/certd/certd/compare/v1.25.1...v1.25.2) (2024-09-24)

**Note:** Version bump only for package @certd/ui-client

## [1.25.1](https://github.com/certd/certd/compare/v1.25.0...v1.25.1) (2024-09-24)

**Note:** Version bump only for package @certd/ui-client

# [1.25.0](https://github.com/certd/certd/compare/v1.24.4...v1.25.0) (2024-09-24)

### Bug Fixes

* 修复首次创建任务运行时不自动设置当前运行情况的bug ([ecd83ee](https://github.com/certd/certd/commit/ecd83ee136abdd3df9ed2f21ec2ff0f24c0ed9d9))

### Features

* 账号绑定 ([e046640](https://github.com/certd/certd/commit/e0466409d0c021bb415abd94df448c8a0d4799e9))
* 支持中间证书 ([e86756e](https://github.com/certd/certd/commit/e86756e4c65a53dd23106d7ecbfe2fa987cc13f3))
* 支持vip转移 ([361e8fe](https://github.com/certd/certd/commit/361e8fe7ae5877e23fd5de31bc919bedd09c57f5))

### Performance Improvements

* 群晖支持OTP双重验证登录 ([8b8039f](https://github.com/certd/certd/commit/8b8039f42bbce10a4d0e737cdeeeef9bb17bee5a))
* 任务支持禁用 ([8ed16b3](https://github.com/certd/certd/commit/8ed16b3ea2dfe847357863a0bfa614e4fa5fc041))
* 优化收件邮箱输入 ([22ef28f](https://github.com/certd/certd/commit/22ef28f6338a78465bd52ccbad13e66e80263b2f))
* 支持阿里云ACK证书部署 ([d331fea](https://github.com/certd/certd/commit/d331fea47789122650e057ec7c9e85ee8e66f09b))
* 支持七牛云 ([8ecc2f9](https://github.com/certd/certd/commit/8ecc2f9446a9ebd11b9bfbffbb6cf7812a043495))
* plugins增加图标 ([a8da658](https://github.com/certd/certd/commit/a8da658a9723342b4f43a579f7805bfef0648efb))

## [1.24.4](https://github.com/certd/certd/compare/v1.24.3...v1.24.4) (2024-09-09)

### Performance Improvements

* 插件选择支持搜索 ([d1498a7](https://github.com/certd/certd/commit/d1498a71601b74d38343b1d070eadd03705dd9d5))

## [1.24.3](https://github.com/certd/certd/compare/v1.24.2...v1.24.3) (2024-09-06)

**Note:** Version bump only for package @certd/ui-client

## [1.24.2](https://github.com/certd/certd/compare/v1.24.1...v1.24.2) (2024-09-06)

### Bug Fixes

* 修复复制流水线出现的各种问题 ([6314e8d](https://github.com/certd/certd/commit/6314e8d7eb58cd52e2a7bd3b5ffb9112b0b69577))

### Performance Improvements

* 阶段、任务、步骤全面支持拖动排序 ([bd73a16](https://github.com/certd/certd/commit/bd73a163cd0497f062bd424ddc6bc9bbc95f81ea))
* 任务配置不需要的字段可以自动隐藏 ([192d9dc](https://github.com/certd/certd/commit/192d9dc7e36737d684c769f255f407c28b1152ac))
* 任务支持拖动排序 ([1e9b563](https://github.com/certd/certd/commit/1e9b5638aa36a8ce70019a9c750230ba41938327))
* client 请求超时时间延长为10s ([ff46771](https://github.com/certd/certd/commit/ff46771d8dd43e71c1ca70e3ba783945750342cc))

## [1.24.1](https://github.com/certd/certd/compare/v1.24.0...v1.24.1) (2024-09-02)

### Bug Fixes

* 激活仅限管理员 ([1c17970](https://github.com/certd/certd/commit/1c17970b981f0987c506744ee6b2283fd5e40493))

### Performance Improvements

* 授权配置支持加密 ([42a56b5](https://github.com/certd/certd/commit/42a56b581d754c3e5f9838179d19ab0d004ef2eb))
* 支持阿里云 DCDN ([98b77f8](https://github.com/certd/certd/commit/98b77f80843834616fb26f83b4c42245326abd06))
* 支持已跳过的步骤重新运行 ([ea775ad](https://github.com/certd/certd/commit/ea775adae18d57a04470cfba6b9460d761d74035))
* 支持cdnfly ([724a850](https://github.com/certd/certd/commit/724a85028b4a7146c9e3b4df4497dcf2a7bf7c67))
* 支持ftp上传 ([b9bddbf](https://github.com/certd/certd/commit/b9bddbfabb5664365f1232e9432532187c98006c))

# [1.24.0](https://github.com/certd/certd/compare/v1.23.1...v1.24.0) (2024-08-25)

### Bug Fixes

* 部署到腾讯云cdn选择证书任务步骤限制只能选证书 ([3345c14](https://github.com/certd/certd/commit/3345c145b802170f75a098a35d0c4b8312efcd17))
* 修复执行日志没有清理的bug ([22a3363](https://github.com/certd/certd/commit/22a336370a88a7df2a23c967043bae153da71ed5))

### Features

* 支持ECC类型 ([a7424e0](https://github.com/certd/certd/commit/a7424e02f5c7e02ac1688791040785920ce67473))

### Performance Improvements

* 优化成功后跳过的提示 ([7b451bb](https://github.com/certd/certd/commit/7b451bbf6e6337507f4627b5a845f5bd96ab4f7b))
* 优化证书申请成功率 ([968c469](https://github.com/certd/certd/commit/968c4690a07f69c08dcb3d3a494da4e319627345))
* email proxy ([453f1ba](https://github.com/certd/certd/commit/453f1baa0b9eb0f648aa1b71ccf5a95b202ce13f))

## [1.23.1](https://github.com/certd/certd/compare/v1.23.0...v1.23.1) (2024-08-06)

### Performance Improvements

* 优化默认值设置 ([1af19f0](https://github.com/certd/certd/commit/1af19f0ac053fe109782882964533636b5969d6b))

# [1.23.0](https://github.com/certd/certd/compare/v1.22.9...v1.23.0) (2024-08-05)

### Features

* use node 20 ([e8ed972](https://github.com/certd/certd/commit/e8ed97206bf28e83f942db2ef4ea07fa76fd3567))

## [1.22.9](https://github.com/certd/certd/compare/v1.22.8...v1.22.9) (2024-08-05)

### Performance Improvements

* 优化定时任务 ([87e440e](https://github.com/certd/certd/commit/87e440ee2a8b10dc571ce619f28bc83c1e5eb147))

## [1.22.8](https://github.com/certd/certd/compare/v1.22.7...v1.22.8) (2024-08-05)

### Performance Improvements

* 修复删除历史记录没有删除log的bug，新增history管理页面，演示站点启动时不自动启动非管理员用户的定时任务 ([f78ae93](https://github.com/certd/certd/commit/f78ae93eedfe214008c3d071ca3d77c962137a64))

## [1.22.7](https://github.com/certd/certd/compare/v1.22.6...v1.22.7) (2024-08-04)

**Note:** Version bump only for package @certd/ui-client

## [1.22.6](https://github.com/certd/certd/compare/v1.22.5...v1.22.6) (2024-08-03)

### Bug Fixes

* 修复在相同的cron时偶尔无法触发定时任务的bug ([680941a](https://github.com/certd/certd/commit/680941af119619006b592e3ab6fb112cb5556a8b))

### Performance Improvements

* 流水线支持名称模糊查询 ([59897c4](https://github.com/certd/certd/commit/59897c4ceae992ebe2972ca9e8f9196616ffdfd7))

## [1.22.5](https://github.com/certd/certd/compare/v1.22.4...v1.22.5) (2024-07-26)

**Note:** Version bump only for package @certd/ui-client

## [1.22.3](https://github.com/certd/certd/compare/v1.22.2...v1.22.3) (2024-07-25)

**Note:** Version bump only for package @certd/ui-client

## [1.22.2](https://github.com/certd/certd/compare/v1.22.1...v1.22.2) (2024-07-23)

### Bug Fixes

* 修复创建流水线时，无法根据dns类型默认正确的dns授权的bug ([a2c43b5](https://github.com/certd/certd/commit/a2c43b50a6069ed48958fd142844a8568c2af452))

## [1.22.1](https://github.com/certd/certd/compare/v1.22.0...v1.22.1) (2024-07-20)

### Performance Improvements

* 创建证书任务可以选择lege插件 ([affef13](https://github.com/certd/certd/commit/affef130378030c517250c58a4e787b0fc85d7d1))
* 创建证书任务增加定时任务和邮件通知输入 ([427620d](https://github.com/certd/certd/commit/427620d34f3b8ad6933005faf1878908441a2453))

# [1.22.0](https://github.com/certd/certd/compare/v1.21.2...v1.22.0) (2024-07-19)

### Features

* 升级midway，支持esm ([485e603](https://github.com/certd/certd/commit/485e603b5165c28bc08694997726eaf2a585ebe7))

### Performance Improvements

* 优化一些小细节 ([b168852](https://github.com/certd/certd/commit/b1688525dbbbfd67e0ab1cf5b4ddfbe9d394f370))
* 增加备案号设置 ([bd3d959](https://github.com/certd/certd/commit/bd3d959944db63a5690b55ee150e1007133868b9))
* 自动生成jwtkey，无需手动配置 ([390e485](https://github.com/certd/certd/commit/390e4853a570390a97df6a3b3882579f9547eeb4))

## [1.21.2](https://github.com/certd/certd/compare/v1.21.1...v1.21.2) (2024-07-08)

**Note:** Version bump only for package @certd/ui-client

## [1.21.1](https://github.com/certd/certd/compare/v1.21.0...v1.21.1) (2024-07-08)

### Performance Improvements

* 说明优化，默认值优化 ([970c7fd](https://github.com/certd/certd/commit/970c7fd8a0f557770e973d8462ee5684ef742810))

# [1.21.0](https://github.com/certd/certd/compare/v1.20.17...v1.21.0) (2024-07-03)

**Note:** Version bump only for package @certd/ui-client

## [1.20.17](https://github.com/certd/certd/compare/v1.20.16...v1.20.17) (2024-07-03)

**Note:** Version bump only for package @certd/ui-client

## [1.20.16](https://github.com/certd/certd/compare/v1.20.15...v1.20.16) (2024-07-01)

**Note:** Version bump only for package @certd/ui-client

## [1.20.15](https://github.com/certd/certd/compare/v1.20.14...v1.20.15) (2024-06-28)

### Bug Fixes

* 修复无法强制取消任务的bug ([9cc01db](https://github.com/certd/certd/commit/9cc01db1d569a5c45bb3e731f35d85df324a8e62))

### Performance Improvements

* 支持windows文件上传 ([7f61cab](https://github.com/certd/certd/commit/7f61cab101fa13b4e88234e9ad47434e6130fed2))

## [1.20.14](https://github.com/certd/certd/compare/v1.20.13...v1.20.14) (2024-06-23)

### Bug Fixes

* 修复修改密码功能异常问题 ([f740ff5](https://github.com/certd/certd/commit/f740ff517f521dce361284c2c54bccc68aee0ea2))

## [1.20.13](https://github.com/certd/certd/compare/v1.20.12...v1.20.13) (2024-06-18)

### Bug Fixes

* 日志高度越界 ([c4c9adb](https://github.com/certd/certd/commit/c4c9adb8bfd513f57252e523794e3799a9b220f8))
* 修复邮箱设置页面SMTP拼写错误的问题 ([b98f1c0](https://github.com/certd/certd/commit/b98f1c0dd0bc6c6b4f814c578692afdf6d90b88d))
* 修复logo问题 ([7e483e6](https://github.com/certd/certd/commit/7e483e60913d509b113148c735fe13ba1d72dddf))

### Performance Improvements

* 增加警告，修复一些样式错乱问题 ([fd54c2f](https://github.com/certd/certd/commit/fd54c2ffac492222e85ff2f5f49a9ee5cfc73588))

## [1.20.12](https://github.com/certd/certd/compare/v1.20.10...v1.20.12) (2024-06-17)

### Bug Fixes

* 修复aliyun域名超过100个找不到域名的bug ([5b1494b](https://github.com/certd/certd/commit/5b1494b3ce93d1026dc56ee741342fbb8bf7be24))

### Performance Improvements

* 增加系统设置，可以关闭自助注册功能 ([20feace](https://github.com/certd/certd/commit/20feacea12d43386540db6a600f391d786be4014))
* 支持cloudflare域名 ([fbb9a47](https://github.com/certd/certd/commit/fbb9a47e8f7bb805289b9ee64bd46ffee0f01c06))

## [1.20.10](https://github.com/certd/certd/compare/v1.20.9...v1.20.10) (2024-05-30)

### Performance Improvements

* 上传到主机插件支持复制到本机路径 ([92446c3](https://github.com/certd/certd/commit/92446c339936f98f08f654b8971a7393d8435224))
* 优化文件下载包名 ([d9eb927](https://github.com/certd/certd/commit/d9eb927b0a1445feab08b1958aa9ea80637a5ae6))
* 增加任务复制功能 ([39ad759](https://github.com/certd/certd/commit/39ad7597fa0e19cc1f7631bbd6fea0a9e05a62c9))

## [1.20.9](https://github.com/certd/certd/compare/v1.20.8...v1.20.9) (2024-03-22)

**Note:** Version bump only for package @certd/ui-client

## [1.20.8](https://github.com/certd/certd/compare/v1.20.7...v1.20.8) (2024-03-22)

**Note:** Version bump only for package @certd/ui-client

## [1.20.7](https://github.com/certd/certd/compare/v1.20.6...v1.20.7) (2024-03-22)

**Note:** Version bump only for package @certd/ui-client

## [1.20.6](https://github.com/certd/certd/compare/v1.20.5...v1.20.6) (2024-03-21)

### Bug Fixes

* 调整按钮图标到居中位置 ([836d18f](https://github.com/certd/certd/commit/836d18f07e22d00faf2f213bc3301a6672b5bafc))

### Performance Improvements

* 插件贡献文档及示例 ([72fb20a](https://github.com/certd/certd/commit/72fb20abf3ba5bdd862575d2907703a52fd7eb17))

## [1.20.5](https://github.com/certd/certd/compare/v1.20.2...v1.20.5) (2024-03-11)

### Bug Fixes

* 修复腾讯云cdn部署无法选择端点的bug ([154409b](https://github.com/certd/certd/commit/154409b1dfee3ea1caae740ad9c1f99a6e7a9814))

## [1.20.2](https://github.com/certd/certd/compare/v1.2.1...v1.20.2) (2024-02-28)

**Note:** Version bump only for package @certd/ui-client

## [1.2.1](https://github.com/certd/certd/compare/v1.2.0...v1.2.1) (2023-12-12)

**Note:** Version bump only for package @certd/ui-client

**Note:** Version bump only for package @certd/ui-client

# [1.2.0](https://github.com/certd/certd/compare/v1.1.6...v1.2.0) (2023-10-27)

* 🔱: [client] sync upgrade with 2 commits [trident-sync] ([aa3207f](https://github.com/certd/certd/commit/aa3207fca5f15f7c3da789989d99c8ae7d1c4551))

### BREAKING CHANGES

* search支持自定义布局，search.layout、search.collapse转移到 search.container之下。如果想使用原来的search组件，请配置search.is=fs-search-v1

## [1.1.6](https://github.com/certd/certd/compare/v1.1.5...v1.1.6) (2023-07-10)

**Note:** Version bump only for package @certd/ui-client

## [1.1.5](https://github.com/certd/certd/compare/v1.1.4...v1.1.5) (2023-07-03)

**Note:** Version bump only for package @certd/ui-client

## [1.1.4](https://github.com/certd/certd/compare/v1.1.3...v1.1.4) (2023-07-03)

### Bug Fixes

* 成功图标转动的问题 ([f87eee3](https://github.com/certd/certd/commit/f87eee3b9ff1ef9874e79a81fe0ed7104cb9ee8c))

### Performance Improvements

* cancel task ([bc65c0a](https://github.com/certd/certd/commit/bc65c0a786360c087fe95cad93ec6a87804cc5ee))
* flush log ([891a43a](https://github.com/certd/certd/commit/891a43ae6716ff98ed06643f7da2e35199ee195c))

## [1.1.3](https://github.com/certd/certd/compare/v1.1.2...v1.1.3) (2023-07-03)

**Note:** Version bump only for package @certd/ui-client

## [1.1.2](https://github.com/certd/certd/compare/v1.1.1...v1.1.2) (2023-07-03)

**Note:** Version bump only for package @certd/ui-client

## [1.1.1](https://github.com/certd/certd/compare/v1.1.0...v1.1.1) (2023-06-28)

**Note:** Version bump only for package @certd/ui-client

# [1.1.0](https://github.com/certd/certd/compare/v1.0.6...v1.1.0) (2023-06-28)

### Bug Fixes

* 修复access选择类型trigger ([2851a33](https://github.com/certd/certd/commit/2851a33eb2510f038fadb55da29512597a4ba512))

### Features

* 权限控制 ([27a4c81](https://github.com/certd/certd/commit/27a4c81c6d70e70abb3892c3ea58d4719988808a))
* 邮件通知 ([937e3fa](https://github.com/certd/certd/commit/937e3fac19cd03b8aa91db8ba03fda7fcfbacea2))
* cert download ([5a51c14](https://github.com/certd/certd/commit/5a51c14de521cb8075a80d2ae41a16e6d5281259))
* save files ([671d273](https://github.com/certd/certd/commit/671d273e2f9136d16896536b0ca127cf372f1619))

## [1.0.6](https://github.com/certd/certd/compare/v1.0.5...v1.0.6) (2023-05-25)

**Note:** Version bump only for package @certd/ui-client

## [1.0.5](https://github.com/certd/certd/compare/v1.0.4...v1.0.5) (2023-05-25)

**Note:** Version bump only for package @certd/ui-client

## [1.0.4](https://github.com/certd/certd/compare/v1.0.3...v1.0.4) (2023-05-25)

**Note:** Version bump only for package @certd/ui-client

## [1.0.3](https://github.com/certd/certd/compare/v1.0.2...v1.0.3) (2023-05-25)

**Note:** Version bump only for package @certd/ui-client

## [1.0.2](https://github.com/certd/certd/compare/v1.0.1...v1.0.2) (2023-05-24)

**Note:** Version bump only for package @certd/ui-client

## [1.0.1](https://github.com/certd/certd/compare/v1.0.0...v1.0.1) (2023-05-24)

**Note:** Version bump only for package @certd/ui-client

## [1.9.2](https://github.com/fast-crud/fast-crud/compare/v1.9.1...v1.9.2) (2023-03-01)

**Note:** Version bump only for package @fast-crud/fs-admin-antdv

## [1.9.1](https://github.com/fast-crud/fast-crud/compare/v1.9.0...v1.9.1) (2023-03-01)

**Note:** Version bump only for package @fast-crud/fs-admin-antdv

## [0.10.5](https://github.com/fast-crud/fast-crud/compare/v0.10.4...v0.10.5) (2021-07-01)

### Performance Improvements

* fs-admin 与crud demo ([4e6b20f](https://github.com/fast-crud/fast-crud/commit/4e6b20fe19434460853841f371b9fd5f16e5e2d3))
* fs-admin纳入子模块 ([2940d30](https://github.com/fast-crud/fast-crud/commit/2940d30f419bf4bde1e8e791f1fbdb9184818285))
