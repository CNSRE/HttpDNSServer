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

	v1, err := redis.Strings(db.Cli().Do( "zrangebyscore", "region_ip", ip, "+inf", "LIMIT", "0", "1" ))
	if err != nil || len(v1) == 0 {
		fmt.Println( "NewRegionIP err" )
	}

	region.IP = ip
	region.ID = v1[0][ 0 : len(v1[0]) - 2 ]

	return region
}