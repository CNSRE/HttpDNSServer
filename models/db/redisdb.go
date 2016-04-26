package db
import (
	"github.com/garyburd/redigo/redis"
	"time"
)

func Init() {
	pool = newPool(redisServer, redisPassword)
}

func Cli() redis.Conn{
	if( pool == nil ){
		Init()
	}
	return pool.Get()
}

func Close(){
	defer pool.Close()
}



func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle: 3,
		IdleTimeout: 240 * time.Second,
		Dial: func () (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

var (
	pool *redis.Pool
	redisServer = "127.0.0.1:6379"
	redisPassword = "httpdnsserver"
)
