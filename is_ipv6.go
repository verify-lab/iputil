package iputil

import "net"

const ipV6Delimiter = ':'

func IsIPv6(ip net.IP) bool {
	return IsStringIPv6(ip.String())
}

func IsStringIPv6(ip string) bool {
	for i := 0; i < len(ip); i++ {
		if ipV6Delimiter == ip[i] {
			return true
		}
	}

	return false
}
