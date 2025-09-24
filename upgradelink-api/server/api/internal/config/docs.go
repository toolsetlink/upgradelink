package config

// 错误信息及docs 文档
const (
	Err1Msg  = "服务内部错误"
	Err1Docs = "UpgradeLink 服务内部错误, 联系管理员解决。7*24小时 请联系我们协助您解决: http://upgrade.toolsetlink.com/upgrade/contact-us.html"

	Err100Msg  = "请求参数fileKey不能为空"
	Err100Docs = "请参考文档: http://upgrade.toolsetlink.com/upgrade/get-file-upgrade-strategy.html"

	Err101Msg  = "请求参数versionCode格式错误"
	Err101Docs = "请参考文档: http://upgrade.toolsetlink.com/upgrade/get-file-upgrade-strategy.html"

	Err102Msg  = "请求file应用不存在"
	Err102Docs = "未查询到fileKey对应的应用，请确认在系统中是否已经正确创建。 请参考文档: http://upgrade.toolsetlink.com/upgrade/recommend/file/app.html"

	Err103Msg  = "未获取到可使用版本"
	Err103Docs = "未查询到对应的应用版本，请确认在系统中是否已经正确创建。 请参考文档: http://upgrade.toolsetlink.com/upgrade/recommend/file/app-version.html"

	Err200Msg  = "请求参数urlKey不能为空"
	Err200Docs = "请参考文档: http://upgrade.toolsetlink.com/upgrade/get-url-upgrade-strategy.html"

	Err201Msg  = "请求参数versionCode格式错误"
	Err201Docs = "请参考文档: http://upgrade.toolsetlink.com/upgrade/get-url-upgrade-strategy.html"

	Err202Msg  = "请求url应用不存在"
	Err202Docs = "未查询到urlKey对应的应用，请确认在系统中是否已经正确创建。 请参考文档: http://upgrade.toolsetlink.com/upgrade/recommend/url/app.html"

	Err203Msg  = "未获取到可使用版本"
	Err203Docs = "未查询到对应的应用版本，请确认在系统中是否已经正确创建。 请参考文档: http://upgrade.toolsetlink.com/upgrade/recommend/url/app-version.html"

	Err300Msg  = "请求参数apkKey不能为空"
	Err300Docs = "请参考文档: http://upgrade.toolsetlink.com/upgrade/get-apk-upgrade-strategy.html"

	Err301Msg  = "请求参数versionCode格式错误"
	Err301Docs = "请参考文档: http://upgrade.toolsetlink.com/upgrade/get-apk-upgrade-strategy.html"

	Err302Msg  = "请求安卓应用不存在"
	Err302Docs = "未查询到apkKey对应的应用，请确认在系统中是否已经正确创建。 请参考文档: http://upgrade.toolsetlink.com/upgrade/recommend/apk/app.html"

	Err303Msg  = "未获取到可使用版本"
	Err303Docs = "未查询到对应的应用版本，请确认在系统中是否已经正确创建。 请参考文档: http://upgrade.toolsetlink.com/upgrade/recommend/apk/app-version.html"

	Err400Msg  = "请求参数configurationKey不能为空"
	Err400Docs = "请参考文档: http://upgrade.toolsetlink.com/upgrade/get-configuration-upgrade-strategy.html"

	Err401Msg  = "请求参数versionCode格式错误"
	Err401Docs = "请参考文档: http://upgrade.toolsetlink.com/upgrade/get-configuration-upgrade-strategy.html"

	Err402Msg  = "请求安卓应用不存在"
	Err402Docs = "未查询到apkKey对应的应用，请确认在系统中是否已经正确创建。 请参考文档: http://upgrade.toolsetlink.com/upgrade/recommend/configuration/app.html"

	Err403Msg  = "未获取到可使用版本"
	Err403Docs = "未查询到对应的应用版本，请确认在系统中是否已经正确创建。 请参考文档: http://upgrade.toolsetlink.com/upgrade/recommend/configuration/app-version.html"

	Err42901Msg  = "应用任务策略请求流控限制"
	Err42901Docs = "当前升级任务配置的请求流控策略，已经达到限制，请稍后再试。"
)
