package models
import (
	"github.com/garyburd/redigo/redis"
	"../db"
	"fmt"
)
type RegionIP struct {
	ID string
	IP uint32
}

func NewRegionIP(domain string, ip uint32) *RegionIP {

	region := new(RegionIP)

	// 查询 [Domain]IP库关联ID
	v1, err := redis.Strings(db.Cli().Do( "zrangebyscore", domain+"_ip", ip, "+inf", "LIMIT", "0", "1" ))
	if err != nil  {
		fmt.Println( domain , "NewRegionIP err" )
	}
	if( len(v1) > 0 ){
		region.IP = ip
		region.ID = v1[0][ 0 : len(v1[0]) - 2 ]
		return region
	}

 	// 默认IP库关联ID
	v1, err = redis.Strings(db.Cli().Do( "zrangebyscore", "region_ip", ip, "+inf", "LIMIT", "0", "1" ))
	if err != nil || len(v1) == 0 {
		fmt.Println( "NewRegionIP err" )
	}
	region.IP = ip
	region.ID = v1[0][ 0 : len(v1[0]) - 2 ]
	return region
}