package config
import "strings"

type ConfigModel struct {

	//************************************************************************************************************

	//lib库开关  {-1 永久关闭} {0 关闭Lib库} {1开启Lib库}
	HTTPDNS_SWITCH string

	//************************************************************************************************************

	//dig检测开关
	HTTPDNS_DIG_SWITCH string

	//************************************************************************************************************

	//是否启用云端更新配置策略 云端一旦关闭,客户端只有安装新版本才能开启
	ENABLE_UPDATE_CONFIG string

	//HTTPDNS 配置接口api地址
	CONFIG_API string

	//************************************************************************************************************

	//日志记录开关
	HTTPDNS_LOG_SWITCH string

	//日志上传的间隔时间
	SCHEDULE_LOG_INTERVAL string

	//************************************************************************************************************

	//是否启用自己家的HTTP_DNS服务器 默认不启用 | 1启用 0不启用
	IS_MY_HTTP_SERVER string

	//自己家HTTP_DNS服务API地址 使用时直接在字符串后面拼接domain地址 | 示例（http://202.108.7.153/dns?domain=）+ domain
	HTTPDNS_SERVER_API []string

	//************************************************************************************************************

	//是否启用udpdns服务器 默认不启用 | 1启用 0不启用
	IS_UDPDNS_SERVER  string

	//udp dnsserver的地址
	UDPDNS_SERVER_API string

	//************************************************************************************************************

	//是否启用dnspod服务器 默认不启用 | 1启用 0不启用
	IS_DNSPOD_SERVER string

	//DNSPOD HTTP_DNS 服务器API地址 | 默认（http://119.29.29.29/d?ttl=1&dn=）
	DNSPOD_SERVER_API string

	//DNSPOD 企业级ID配置选项
	DNSPOD_ID string

	//DNSPOD 企业级KEY配置选项
	DNSPOD_KEY string

	//************************************************************************************************************

	//是否开启 本地排序插件算法 默认开启 | 1开启 0不开启
	IS_SORT string

	//速度插件 比重分配值：默认40%
	SPEEDTEST_PLUGIN_NUM string

	//服务器推荐优先级插件 比重分配：默认30% （需要自家HTTP_DNS服务器支持）
	PRIORITY_PLUGIN_NUM string

	//历史成功次数计算插件 比重分配：默认10%
	SUCCESSNUM_PLUGIN_NUM string

	//历史错误次数计算插件 比重分配：默认10%
	ERRNUM_PLUGIN_NUM string

	//最后一次成功时间计算插件 比重分配：默认10%
	SUCCESSTIME_PLUGIN_NUM string

	//************************************************************************************************************

	//测速间隔时间
	SCHEDULE_SPEED_INTERVAL string

	//timer轮询器的间隔时间
	SCHEDULE_TIMER_INTERVAL string

	//ip数据过期延迟差值
	IP_OVERDUE_DELAY string

	//************************************************************************************************************

	//白名单
	DOMAIN_SUPPORT_LIST []string

	//************************************************************************************************************

}


// 计算字符串：新浪微博+com.sina.weibo+android+httpdns  key: e4982a86a993f9691fe57feb9ef2b200
// 计算字符串：新浪微博+com.sina.weibo+ios+httpdns  key: ed42adc6e52fd7fd81e53fb6424497a1

// 计算字符串：微博头条+com.sina.app.weiboheadline+android+httpdns  key: 18e3116aa3603d53eefe74e5f0a4ade9
// 计算字符串：微博头条+com.sina.app.weiboheadline+ios+httpdns  key: 93a392d81b2799abfe7359ccd188db56

func Find(appkey string) ( data ConfigModel, err error ){

	if strings.EqualFold(appkey, "18e3116aa3603d53eefe74e5f0a4ade9"){
		data = ConfigModel{
			"1",
			"0",
			"1", "api.dnssdk.com:8080/config",
			"1", "3600000",
			"1" , []string{"http://dns.weibo.cn","http://202.108.7.232","http://221.179.190.246","http://58.63.236.228"},
			"0", "114.114.114.114",
			"0", "http://119.29.29.29/d?ttl=1&dn=", "22", "j2cjxCp2",
			"1", "50", "30", "10", "10", "10",
			"60000", "60000", "60",
			[]string{"v.top.weibo.cn",},
		}
	}

	return data, err
}


