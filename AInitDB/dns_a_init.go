package main
import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"./db"
	"strings"
	"strconv"
	"container/list"
)

//------------------------------------------------------------------------
//
//[Domain][IP段ID]关联DNS记录：（Hash）[运行中增长，最大:约region_ip条数]
//
//说明：
//1.存储每个Domain下所有IP段的A记录信息
//2.该数据不会直接缓存所有region中的IP段对应的DNS信息，会根据用户命中此IP段后在进行缓存。
//3.新增一个域名就会新增加一个HSET的数据结构
//
//数据源：
//IP段ID | DNS记录
//21658,202.108.7.239|202.108.7.240|60
//
//存储格式：
//redis：HSET [domain]_dns IP段ID DNS记录
//
//------------------------------------------------------------------------

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

var domainStr = "api.weibo.cn"


func main() {

	db.RedisDial()

	v1, err := redis.Strings(db.Cli().Do("HVALS", "ip_info"))
	if err != nil {
		fmt.Println("redis err")
	}
	fmt.Println("数据总长度", len(v1))

	l := list.New()

	vlen := len(v1)

	for i := 0; i < vlen; i++ {

		ipinfo := MakeIpInfo(v1[i])

		jsonStr := httpGet(domainStr, LongStr2IP(ipinfo.StartIP))

		if len(jsonStr) < 10 {
			fmt.Println("err jsonStr: ", domainStr , ipinfo.StartIP,  jsonStr)
			l.PushBack(ipinfo.ID)
			continue
		}

		obj := json2Obj(jsonStr);

		if len(obj.Domain) != len(domainStr) {
			fmt.Println("\n\n err obj.Domain: ", domainStr , ipinfo.StartIP,  jsonStr , "\n\n")
			l.PushBack(ipinfo)
			continue
		}

		if len(obj.Device_ip) < 7 {
			fmt.Println("\n\n err obj.Device_ip: ", domainStr , ipinfo.StartIP,  jsonStr , "\n\n")
			l.PushBack(ipinfo)
			continue
		}

		dnsStr := ""
		if len(obj.DNS) >= 1 {
			for j := 0 ; j < len( obj.DNS ) ; j++{
				dnsStr += obj.DNS[j].IP + ","
			}
			dnsStr += obj.DNS[0].TTL
		}else{
			dnsStr = dnsPodHttpGet(domainStr, LongStr2IP(ipinfo.StartIP))
			dnsStr = strings.Replace(dnsStr, ";", ",", -1)
		}


		if len(dnsStr) == 0{
			fmt.Println("\n\n err obj.Device_ip: ", domainStr , ipinfo.StartIP,  jsonStr , "\n\n")
			l.PushBack(ipinfo)
			continue
		}

		_, err := db.Cli().Do("HMSET", domainStr + "_dns" , ipinfo.ID , dnsStr )
		if err != nil {
			fmt.Println("\n\n err redis do: ", domainStr , ipinfo.StartIP,  jsonStr , "\n\n")
			l.PushBack(ipinfo)
		}

		if i % 100 == 0 {
			fmt.Print(i, " ")
		}
		if i % 1000 == 0 {
			fmt.Println()
		}

	}

	fmt.Println("\n\n\n" , "错误list: \n\n")

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println("\n", e.Value)
	}

	db.RedisClose()
}


func httpGet(domain string, ip string) string {

	resp, err := http.Get("http://dns.weibo.cn/dns?domain=" + domain + "&ip=" + ip)
	if err != nil {
		fmt.Println("err : ", domain, "+", ip)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err : ", domain, "+", ip)
	}

	return string(body)
}

func dnsPodHttpGet(domain string, ip string) string {

	resp, err := http.Get("http://119.29.29.29/d?dn=" + domain + "&ip=" + ip)
	if err != nil {
		fmt.Println("err : ", domain, "+", ip)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err : ", domain, "+", ip)
	}

	return string(body)

}


func json2Obj(jsonStr string) *dataModel {
	res := &dataModel{}
	json.Unmarshal([]byte(jsonStr), &res)
	return res
}

type IpInfo struct {
	ID       string //id编号
	IP       string //ip段
	StartIP  string //开始IP
	EndIP    string //结束IP
	Country  string //国家
	Province string //省
	City     string //市
	District string //区
	Isp      string //运营商
	Type     string //类型
	Desc     string //说明
}

func MakeIpInfo(str string) *IpInfo {
	ipinfo := new(IpInfo)
	str = strings.Trim(str, "\n")
	strArr := strings.Split(str, ",")
	ipinfo.ID = strArr[0]
	ipinfo.IP = strArr[1]
	ipinfo.StartIP = strArr[2]
	ipinfo.EndIP = strArr[3]
	ipinfo.Country = strArr[4]
	ipinfo.Province = strArr[5]
	ipinfo.City = strArr[6]
	ipinfo.District = strArr[7]
	ipinfo.Isp = strArr[8]
	ipinfo.Type = strArr[9]
	ipinfo.Desc = strArr[10]
	return ipinfo
}

func LongStr2IP(lstr string) string {
	uin, _ := strconv.ParseUint(lstr, 10, 32)
	return Long2IP(uint32(uin))
}

func Long2IP(i uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", i >> 24, (i >> 16) & 0xFF, (i >> 8) & 0xFF, i & 0xFF)
}