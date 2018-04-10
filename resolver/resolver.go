package resolver

import (
	"net"
)

// Resolve is dns resolve handler
func Resolve(host string) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	res, err := net.LookupTXT(host)
	if err == nil {
		data["TXT"] = res
	}
	res1, err := net.LookupIP(host)
	if err == nil {
		data["IP"] = res1
	}
	res2, err := net.LookupCNAME(host)
	if err == nil {
		data["CNAME"] = res2
	}
	res3, err := net.LookupMX(host)
	if err == nil {
		data["MX"] = res3
	}
	res4, err := net.LookupNS(host)
	if err == nil {
		data["NS"] = res4
	}
	res5, err := net.LookupHost(host)
	if err == nil {
		data["HOST"] = res5
	}
	return data, err
}
