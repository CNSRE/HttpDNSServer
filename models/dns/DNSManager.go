package dns
import (
	"fmt"
	"../iplookup"
	"../cache"
)

type dataModel struct {
	Domain    string `json:"domain"`
	Device_ip string `json:"device_ip"`
	Device_sp string `json:"device_sp"`
	DNS       []dnsModel `json:"dns"`
}

type dnsModel struct {
	Priority string `json:"priority"`
	IP       string `json:"ip"`
	TTL      string `json:"ttl"`
}

func Find(domain, ip, userIp string) ( data dataModel, err error ){

	// 正则检测是否是合法域名
	if len( domain ) == 0 {
		return data , fmt.Errorf( "dnsManager domain is null" )
	}

	// 正则检测是否是合法IP
	if len( ip ) == 0 {
		ip = userIp
		//return data , fmt.Errorf( "dnsManager ip is null" )
	}

	// iplookup 模块查询ip对应的region id
	regID , ipErr := iplookup.FindID(domain ,ip )
	if ipErr != nil {
		return data , fmt.Errorf( "终止执行,错误返回,记录日志" , domain , userIp , err )
	}

	// cache 查询数据
	a , t , cacheErr := cache.Find(domain, regID)
	if cacheErr != nil{
		return data , fmt.Errorf( "终止执行,错误返回,记录日志" , domain , userIp , err )
	}

	// 查询IP相关信息
	ipInfo , infoErr := iplookup.FindIpInfo(regID)
	if infoErr != nil {

	}

	// ends递归查询数据( 等震动 )


	// 组装数据
	data.Domain = domain
	data.Device_ip = ip
	data.Device_sp = ipInfo.Isp
	data.DNS = make( []dnsModel, len(a) )
	for i := 0 ; i < len(a) ; i++ {
		data.DNS[i] = dnsModel{ "0" , a[i] , t }
	}

	return data , nil
}