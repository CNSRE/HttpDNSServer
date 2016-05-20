package cache
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"../db"
	"strings"
)

type DomainDNS struct {
	ID string
	A  string
}

func Find(domain, id string) (a []string, ttl string, err error) {

	conn := db.Get()
	defer conn.Close()

	v1, e := redis.String(conn.Do("HGET", domain + "_dns", id))
	if err != nil || len(v1) == 0 {
		return nil, "", fmt.Errorf("cache find err. redis : ", e, " domain:", domain, " id:", id)
	}

	tempArr := strings.Split(v1, ",")
	if len( tempArr ) < 2{
		return nil, "", fmt.Errorf("cache find err. no dns data : "," domain:", domain, " id:", id)
	}

	return tempArr[0:len(tempArr) - 1 ], tempArr[len(tempArr) - 1] , nil
}