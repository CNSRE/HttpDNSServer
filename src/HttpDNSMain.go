package main

import (
	"./iplookup"
	"fmt"
	"time"
)

func main() {

	iplookup.InitIpNet()

	for i := 0 ; i < 1000 ; i++{
		t1 := time.Now().UnixNano();
		ipNet := iplookup.FindIpNet( 253606 )
		t2 := time.Now().UnixNano();
		fmt.Println(t2-t1);
		fmt.Println(ipNet)
	}



}