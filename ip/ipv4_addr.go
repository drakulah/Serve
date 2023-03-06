package ip

import (
	"net"
	"os"
)

func GetIPv4Addr() (ipv4Addr string) {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			ipv4Addr = ipv4.String()
		}
	}
	return
}
