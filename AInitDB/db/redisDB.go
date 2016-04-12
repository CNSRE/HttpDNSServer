package db

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

var cli redis.Conn = nil

var REDIS_SERVER_IP = "127.0.0.1:6379"

func RedisDial() {
	var err error
	cli, err = redis.Dial("tcp", REDIS_SERVER_IP )
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Cli() redis.Conn{
	if( cli == nil ){
		fmt.Println( "init redis" )
		RedisDial()
	}
	return cli
}

func RedisClose(){
	defer cli.Close()
}