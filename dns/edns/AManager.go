package edns
import (
	"github.com/miekg/dns"
	"strings"
	"net"
	"time"
	"fmt"
	"regexp"
)

func FindA(ednsModel* EDNSModel){

	// 从缓存查询


	// 从ns服务器查询
	a := findA(ednsModel)
	if strings.EqualFold( a , "")  == false{
		a = strings.TrimRight(a, ",")
		ednsModel.A = strings.Split( a ,"," )
	}
}


func findA(ednsModel* EDNSModel)(string){

	var domain_a string

	var server string
	if( len( ednsModel.NS ) != 0 ){
		server = ednsModel.NS[0]
	}else if( len(ednsModel.SOA) != 0 ){
		server = ednsModel.SOA[0]
	}else{
		server = OPEN_DNS_SERVER
	}
	if dns.IsFqdn(server) {
		server = server[0 : len(server)-1]
	}
	if !strings.HasSuffix(server, ":53") {
		server += ":53"
	}

	domain := dns.Fqdn(ednsModel.CName[len(ednsModel.CName)-1])
	msg := new(dns.Msg)
	msg.SetQuestion(domain, dns.TypeA)
	msg.RecursionDesired = true

	if ednsModel.ClientIP != "" {

		opt := new(dns.OPT)
		opt.Hdr.Name = "."
		opt.Hdr.Rrtype = dns.TypeOPT

		e := new(dns.EDNS0_SUBNET)
		e.Code = dns.EDNS0SUBNET
		e.Family = 1 // ipv4
		e.SourceNetmask = 32
		e.SourceScope = 0
		e.Address = net.ParseIP(ednsModel.ClientIP).To4()

		opt.Option = append(opt.Option, e)
		msg.Extra = []dns.RR{opt}
	}

	client := &dns.Client{
		DialTimeout:  5 * time.Second,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	resp, rtt , err := client.Exchange(msg, server)
//	fmt.Println(resp.Answer)

	if err != nil {
		fmt.Println( rtt , err) // 记录日志  rtt是查询耗时
		return ""
	}



	for i := len(resp.Answer)-1 ; i >= 0 ; i--{
		switch resp.Answer[i].Header().Rrtype {
		case dns.TypeA:
			temp_a := resp.Answer[i].(*dns.A)
			domain_a +=  fmt.Sprint(temp_a.A , ":" , temp_a.Hdr.Ttl , ",")
			break
		case dns.TypeCNAME:
			temp_cname := resp.Answer[i].(*dns.CNAME)
			ednsModel.CName = append( ednsModel.CName , temp_cname.Target )
			break;
		}
	}

	return domain_a
}


func IsIP(ip string) (b bool) {
	if m, _ := regexp.MatchString("^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$", ip); !m {
		return false
	}
	return true
}