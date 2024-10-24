package xutil

import (
	"net"
)

func SplitHostPort(addr string) (host string, port string) {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return "", ""
	}
	return host, port
}
