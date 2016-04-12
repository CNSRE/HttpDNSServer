package db
import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

var cli redis.Conn = nil

func Init() {
	var err error
	cli, err = redis.Dial("tcp", "127.0.0.1:6379" )
	if err != nil {
		return
	}
}

func Cli() redis.Conn{
	if( cli == nil ){
		Init()
	}
	return cli
}

func Close(){
	defer cli.Close()
}