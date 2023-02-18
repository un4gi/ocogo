package main

import (
	"net"
	"strings"
)

func reverseLookup(ip string) string {
	var host string

	ptr, _ := net.LookupAddr(ip)
	for _, host = range ptr {
		if strings.HasSuffix(host, ".") {
			host = strings.TrimSuffix(host, ".")
		}
	}
	return host
}
