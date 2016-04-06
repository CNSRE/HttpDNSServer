package iplookup

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	"sort"
)

type IpNet struct{
	ID int //id编号
	IP string //ip段
	StartIP int64 //开始IP
	StopIP int64  //结束IP
	Country int //国家
	Province int //省
	City int //市
	District int //区
	Isp int //运营商
	Type string //类型
	Desc string //说明
}

func (p IpNetArray) Len() int           {
	return len(p)
}
func (p IpNetArray) Less(i, j int) bool {
	return p[i].StartIP < p[j].StopIP
}
func (p IpNetArray) Swap(i, j int)      {
	p[i], p[j] = p[j], p[i]
}

type IpNetArray []IpNet
var ipNetArray IpNetArray

var filename = "res/ip.db"

func InitIpNet(){

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		return
	}
	arr := strings.Split(string(buf), "\n")
	ipNetArray = make( IpNetArray , len(arr)  )

	for i , v := range arr {
		tempArr := strings.Split(v,",")
		if( len( tempArr ) < 10 ) { continue }
		ipNetArray[i] = IpNet{ }
		ipNetArray[i].ID = i
		ipNetArray[i].IP = tempArr[0]
		ipNetArray[i].StartIP , _ = strconv.ParseInt(tempArr[1], 10, 64)
		ipNetArray[i].StopIP , _ = strconv.ParseInt(tempArr[2], 10, 64)
		ipNetArray[i].Country ,_ = strconv.Atoi( tempArr[3])
		ipNetArray[i].Province ,_ = strconv.Atoi( tempArr[4])
		ipNetArray[i].City ,_ = strconv.Atoi( tempArr[5])
		ipNetArray[i].District ,_ = strconv.Atoi( tempArr[6])
		ipNetArray[i].Isp ,_ = strconv.Atoi( tempArr[7])
		ipNetArray[i].Type = tempArr[8]
		ipNetArray[i].Desc = tempArr[9]
	}

	sort.Sort(ipNetArray)

}


func FindIpNet(ip int64)IpNet{

	index := sort.Search(len(ipNetArray), func(i int) bool { return ipNetArray[i].StopIP >= ip })

	return ipNetArray[index]
}