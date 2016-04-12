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
		// 这里要处理成启动失败
		return
	}
}

func Cli() redis.Conn{
	if( cli == nil ){
		fmt.Println( "init redis" )
		Init()
	}
	return cli
}

func Close(){
	defer cli.Close()
}