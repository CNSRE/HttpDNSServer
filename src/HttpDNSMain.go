package main

import (
	"./iplookup"
	"fmt"
)

func main() {

	iplookup.InitIpNet()

	ipNet := iplookup.FindIpNet( 3659177983 )

	fmt.Println(ipNet)

}
