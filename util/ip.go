package util
import (
	"net"
	"fmt"
	"encoding/binary"
)


func IpStr2Long(ip string) uint32{
	IP := net.ParseIP(ip)
	return Ip2Long(IP)
}

func Ip2Long(ip net.IP) uint32 {
	if ip == nil && len(ip) < 4 {
		fmt.Println(ip)
		return 0
	}
	return binary.BigEndian.Uint32(ip.To4())
}

func Long2IP(i uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", i>>24, (i>>16)&0xFF, (i>>8)&0xFF, i&0xFF)
}