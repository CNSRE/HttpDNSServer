package edns
import (
	"github.com/miekg/dns"
	"strings"
//	"fmt"
)

func FindSoaNs(ednsModel* EDNSModel){

	// 先查询该域名是否有指定的NS


	// 查询缓存层SOA解析. 如果多数SOA TTL和域名时间一致的话,就没必要增加SOA的缓存层



	// 从本机上级DNS服务器获取域名的NS
	cname , soa , ns := findSoaNs(ednsModel.Domain)

	if strings.EqualFold( cname , "")  == false{
		cname = strings.TrimRight(cname, ",")
		ednsModel.CName = strings.Split( cname ,"," )
	}
	if strings.EqualFold( soa , "")  == false {
		soa = strings.TrimRight(soa, ",")
		ednsModel.SOA = strings.Split(soa, ",")
	}
	if strings.EqualFold( ns , "")  == false {
		ns = strings.TrimRight(ns, ",")
		ednsModel.NS = strings.Split(ns, ",")
	}
}



func findSoaNs(domain string) (string, string, string){

	var cname string
	var soa string
	var ns string

	add := func(c, s ,n string) () {
		cname += c
		soa += s
		ns += n
		return
	}

	cname += domain + ","
	m1 := new(dns.Msg)
	m1.Id = dns.Id()
	m1.RecursionDesired = true
	m1.Question = make([]dns.Question, 1)
	m1.Question[0] = dns.Question{domain , dns.TypeSOA, dns.ClassINET}
	in, _ := dns.Exchange(m1, (cf.Servers[1]+":53") )
	rrList := [...][]dns.RR{in.Answer , in.Ns , in.Extra}


	for _, rr := range rrList{
		for i := len(rr)-1 ; i >= 0 ; i--{
			switch rr[i].Header().Rrtype {
			case dns.TypeCNAME:
				temp_cname := rr[i].(*dns.CNAME)
				add(findSoaNs(temp_cname.Target))
//				fmt.Println(  "temp_cname:" , temp_cname )
				return cname , soa, ns
				break
			case dns.TypeNS:
				temp_ns := rr[i].(*dns.NS)
				ns += temp_ns.Ns + ","// + "|" +  fmt.Sprint( temp_ns.Hdr.Ttl ) + ","
//				fmt.Println(  "temp_ns:" , temp_ns )
				break
			case dns.TypeSOA:
				temp_soa := rr[i].(*dns.SOA)
				soa += temp_soa.Ns + ","// + "|" + fmt.Sprint( temp_soa.Hdr.Ttl ) + ","
//				fmt.Println( "temp_soa:" , temp_soa )
				break
			}
		}
	}

	return cname , soa , ns
}