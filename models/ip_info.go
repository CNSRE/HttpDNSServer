package models
import (
	"github.com/garyburd/redigo/redis"
	"../db"
	"fmt"
	"strings"
)

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

func NewIpInfo(id string) *IpInfo {
	ipinfo := new(IpInfo)
	v1, err := redis.String(db.Cli().Do("HGET", "ip_info", id))
	if err != nil {
		fmt.Println("NewIpInfo err")
	}
	str := strings.Trim(v1, "\n")
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