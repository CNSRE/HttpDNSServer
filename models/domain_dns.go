package models
import (
	"fmt"
	"../db"
	"github.com/garyburd/redigo/redis"
)

type DomainDNS struct {
	ID string
	A string
}

func NewDomainDNS(domain , id string) *DomainDNS {

	dns := new(DomainDNS)
	v1, err := redis.String(db.Cli().Do( "HGET", domain+"_dns", id))
	if err != nil || len(v1) == 0 {
		fmt.Println( "NewRegionIP err" )
	}
	dns.ID = id
	dns.A = v1

	return dns
}
