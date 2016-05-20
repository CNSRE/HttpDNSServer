package iplookup
import (
	"strings"
	"github.com/garyburd/redigo/redis"
	"../db"
	"fmt"
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

func FindIpInfo(id string) (ipInfo IpInfo, err error) {

	conn := db.Get()
	defer conn.Close()

	v1, e := redis.String(conn.Do("HGET", "ip_info", id))
	if e != nil {
		return ipInfo, fmt.Errorf("find ip info err. redis: id:", id)
	}
	str := strings.Trim(v1, "\n")
	strArr := strings.Split(str, ",")
	ipInfo.ID = strArr[0]
	ipInfo.IP = strArr[1]
	ipInfo.StartIP = strArr[2]
	ipInfo.EndIP = strArr[3]
	ipInfo.Country = strArr[4]
	ipInfo.Province = strArr[5]
	ipInfo.City = strArr[6]
	ipInfo.District = strArr[7]
	ipInfo.Isp = strArr[8]
	ipInfo.Type = strArr[9]
	ipInfo.Desc = strArr[10]
	return ipInfo, nil
}