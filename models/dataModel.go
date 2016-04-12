package models
import "strings"

type dataModel struct {
	Domain    string `json:"domain"`
	Device_ip string `json:"device_ip"`
	Device_sp string `json:"device_sp"`
	DNS       []dnsModel `json:"dns"`
}

type dnsModel struct {
	Priority string `json:"priority"`
	IP       string `json:"ip"`
	TTL      string `json:"ttl"`
}

func NewdataModel(domain , ip , sp , a string) *dataModel {

	data := new(dataModel)

	data.Domain = domain
	data.Device_ip = ip
	data.Device_sp = sp

	tempArr := strings.Split(a, ",")
	data.DNS = make( []dnsModel, len(tempArr) - 1  )
	for i := 0 ; i < len(tempArr) - 1 ; i++{
		data.DNS[i] = dnsModel{ "0", tempArr[i], tempArr[len(tempArr) - 1] }
	}

	return data
}