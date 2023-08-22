package telnet

import (
	"regexp"
)

func ParseIp(ip string) bool {
	reg := regexp.MustCompile("^\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}:\\d+$")
	if reg.MatchString(ip) {
		return true
	}

	return false
}

func ParseDN(dn string) bool {
	reg := regexp.MustCompile("^[a-zA-Z0-9.-]+:\\d+$")
	if reg.MatchString(dn) {
		return true
	}
	return false

}

func Concat(addr, port string) (string, bool) {
	s := addr + ":" + port
	if ParseIp(s) {
		return s, true
	} else if ParseDN(s) {
		return s, true
	}
	return "", false
}
