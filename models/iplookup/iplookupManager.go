package iplookup

import (
	"github.com/garyburd/redigo/redis"
	"../db"
	"../util"
	"fmt"
)
type RegionIP struct {
	ID string
	IP uint32
}

func FindID(domain string, ipStr string) (string, error) {

	conn := db.Get()
	defer conn.Close()

	// ip 转 int类型
	ip := util.IpStr2Long(ipStr)

	// 查询 [Domain]IP库关联ID
	v1, err := redis.Strings(conn.Do("zrangebyscore", domain + "_ip", ip, "+inf", "LIMIT", "0", "1"))
	if err != nil {
		return "", fmt.Errorf("iplookup find id err. redis : ", err, " domain:", domain, " ip:", ipStr)
	}
	if ( len(v1) > 0 ) {
		return v1[0][ 0 : len(v1[0]) - 2 ], nil
	}

	// 默认IP库关联ID
	v1, err = redis.Strings(conn.Do("zrangebyscore", "region_ip", ip, "+inf", "LIMIT", "0", "1"))
	if err != nil {
		return "", fmt.Errorf("iplookup find id err. redis : ", err, " domain:", domain, " ip:", ipStr)
	}
	if len(v1) == 0 {
		// 如果默认IP库没有命中 为该ip创建 region id 记录
	}

	return v1[0][ 0 : len(v1[0]) - 2 ], nil
}
