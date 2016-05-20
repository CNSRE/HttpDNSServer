package main

import (
	"./dns/edns"
	"fmt"
)


var domainList = []string{
	"api.weibo.cn",
	"v.top.weibo.cn",
	"ww1.sinaimg.cn",
	"ww2.sinaimg.cn",
	"ww3.sinaimg.cn",
	"ww4.sinaimg.cn",
	"tp1.sinaimg.cn",
	"tp2.sinaimg.cn",
	"tp3.sinaimg.cn",
	"tp4.sinaimg.cn",

	"ww1.sinaimg.com",
	"ww2.sinaimg.com",
	"ww3.sinaimg.com",
	"ww4.sinaimg.com",

	"dnssdk.com",
	"www.baidu.com",
	"www.taobao.com",
	"www.hao123.com",
	"www.gexing.com",
	"www.qingguo.com",
	"www.qq.com",

}

func main() {

	edns.Init()

	for _, domain := range domainList {
		ednsModel := edns.Find(domain , "123.125.105.139")
		fmt.Println( ednsModel.Domain , " : " , ednsModel.String() )
	}


}
