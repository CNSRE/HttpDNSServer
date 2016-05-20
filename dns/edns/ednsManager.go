package edns
import (
	"github.com/miekg/dns"
	"fmt"
)

var DEFAULT_RESOLV_FILE = "/etc/resolv.conf"
var OPEN_DNS_SERVER = "114.114.114.114"
var cf* dns.ClientConfig

func Init(){
	var el error
	cf, el = dns.ClientConfigFromFile(DEFAULT_RESOLV_FILE)
	if el != nil {
		fmt.Println( "DEFAULT_RESOLV_FILE ERR" )
	}
}

func Find( domain string , ip string)(EDNSModel){

	// 判断域名是否是标准domain
	domain = dns.Fqdn(domain)

	// 存储对象
	ednsModel := EDNSModel{domain,ip,nil,nil,nil,nil}

	// 查询SOA记录
	FindSoaNs( &ednsModel )

	// 查询A记录
	FindA(&ednsModel)

	return ednsModel
}


type EDNSModel struct {
	Domain  	string
	ClientIP	string
	A[]			string
	CName[]		string
	SOA[]		string
	NS[]		string
}

func (e *EDNSModel) String()(string) {
	str := "dns find : \n"
	str += fmt.Sprint( "Domain :" , e.Domain , "\n" )
	str += fmt.Sprint( "ClientIP :" , e.ClientIP , "\n" )
	str += fmt.Sprint( "A : " , e.A , "\n" )
	str += fmt.Sprint(  "CName :" , e.CName , "\n" )
	str += fmt.Sprint(  "SOA : " , e.SOA , "\n" )
	str += fmt.Sprint(  "NS :" , e.NS , "\n\n" )
	return str
}