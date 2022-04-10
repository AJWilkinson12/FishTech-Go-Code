package main

import (
	"net"
)

func Valid(ip string) bool {

	ipObj := net.ParseIP(ip)

	if ipObj == nil {
		return false
	} else {
		return true
	}

}
