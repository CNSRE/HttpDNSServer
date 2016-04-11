package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"strconv"
	"sort"
)

var REDIS_SERVER_IP = "127.0.0.1:6379"

type IpNet struct {
	ID       int    //id编号
	IP       string //ip段
	StartIP  int64  //开始IP
	StopIP   int64  //结束IP
	Country  int    //国家
	Province int    //省
	City     int    //市
	District int    //区
	Isp      int    //运营商
	Type     string //类型
	Desc     string //说明
}


func (p IpNet) String() string {
	str := ""
	str += strconv.Itoa(p.ID) + ","
	str += strconv.FormatInt(p.StartIP, 10) + ","
	str += strconv.FormatInt(p.StopIP, 10) + ","
	str += strconv.Itoa(p.Country) + ","
	str += strconv.Itoa(p.Province) + ","
	str += strconv.Itoa(p.City) + ","
	str += strconv.Itoa(p.District) + ","
	str += strconv.Itoa(p.Isp) + ","
	str += p.Type + ","
	str += p.Desc
	return str
}

type IpNetArray []IpNet

var ipNetArray IpNetArray

func (p IpNetArray) Len() int {
	return len(p)
}
func (p IpNetArray) Less(i, j int) bool {
	return p[i].StartIP < p[j].StopIP
}
func (p IpNetArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}


var filename = "res/ip.db"

func InitIpNet() {

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		return
	}
	arr := strings.Split(string(buf), "\n")
	ipNetArray = make(IpNetArray, len(arr))

	for i, v := range arr {
		tempArr := strings.Split(v, ",")
		if ( len(tempArr) < 10 ) { continue }
		ipNetArray[i] = IpNet{}
		ipNetArray[i].ID = i
		ipNetArray[i].IP = tempArr[0]
		ipNetArray[i].StartIP, _ = strconv.ParseInt(tempArr[1], 10, 64)
		ipNetArray[i].StopIP, _ = strconv.ParseInt(tempArr[2], 10, 64)
		ipNetArray[i].Country, _ = strconv.Atoi(tempArr[3])
		ipNetArray[i].Province, _ = strconv.Atoi(tempArr[4])
		ipNetArray[i].City, _ = strconv.Atoi(tempArr[5])
		ipNetArray[i].District, _ = strconv.Atoi(tempArr[6])
		ipNetArray[i].Isp, _ = strconv.Atoi(tempArr[7])
		ipNetArray[i].Type = tempArr[8]
		ipNetArray[i].Desc = tempArr[9]
	}

	sort.Sort(ipNetArray)

	for i, _ := range ipNetArray {
		ipNetArray[i].ID = i
	}

}


func main() {
	InitIpNet()
	redisDial();
	defer cli.Close()

	delDB_ip_info()
	insertDB_ip_info()
	printAll_ip_info()


	delDB_region_ip()
	insertDB_region_ip()
	printAll_region_ip()

}

var cli redis.Conn = nil

func redisDial() {
	var err error
	cli, err = redis.Dial("tcp", config.REDIS_SERVER_IP )
	if err != nil {
		fmt.Println(err)
		return
	}
}


func insertDB_ip_info() {
	for i, v := range ipNetArray {
		str := v.String()
		v1, err := cli.Do("HMSET", "ip_info", v.ID, str)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(i, " = ", str, " - ", v1)
	}
}

func printAll_ip_info(){
	for _, v := range ipNetArray {
		v1, err := redis.String(cli.Do("HGET", "ip_info", v.ID))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(v1)
	}
}

func delDB_ip_info(){
	v1, err := cli.Do("DEL", "ip_info")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v1)
}

func insertDB_region_ip(){
	for i, v := range ipNetArray {
		str := v.String()
		v1, err := cli.Do("ZADD", "region_ip", v.StartIP, strconv.Itoa(v.ID) + "_S" , v.StopIP, strconv.Itoa(v.ID) + "_E" ) //redis：ZADD region_ip 开始IP ID_S 结束IP ID_E
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(i, " = ", str, " - ", v1)
	}
}

func printAll_region_ip(){

	for _, v := range ipNetArray {

		v1, err := redis.String(cli.Do("ZSCORE", "region_ip", strconv.Itoa(v.ID) + "_S"))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(strconv.Itoa(v.ID)+"_S" , " = " , v1)

		v1, err = redis.String(cli.Do("ZSCORE", "region_ip", strconv.Itoa(v.ID) + "_E"))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(strconv.Itoa(v.ID)+"_E" , " = " , v1)
	}

}

func delDB_region_ip(){
	v1, err := cli.Do("DEL", "region_ip")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v1)
}
