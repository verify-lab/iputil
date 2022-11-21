package iputil

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrivateIpV4(t *testing.T) {
	cases := []struct {
		ip        net.IP
		isPrivate bool
		msg       string
	}{
		{net.ParseIP("127.0.0.1"), true, "127.0.0.1 should be private"},
		{net.ParseIP("192.168.254.254"), true, "192.168.254.254 should be private"},
		{net.ParseIP("10.255.0.3"), true, "10.255.0.3 should be private"},
		{net.ParseIP("172.16.255.255"), true, "172.16.255.255 should be private"},
		{net.ParseIP("172.31.255.255"), true, "172.31.255.255 should be private"},
		{net.ParseIP("192.169.255.255"), false, "192.169.255.255 should not be private"},
		{net.ParseIP("9.255.0.255"), false, "9.255.0.255 should not be private"},
		{net.ParseIP("67.22.74.181"), false, "67.22.74.181 should not be private"},
		{net.ParseIP("109.86.51.55"), false, "109.86.51.55 should not be private"},
		{net.ParseIP("76.233.86.21"), false, "76.233.86.21 should not be private"},
		{net.ParseIP("109.42.112.5"), false, "109.42.112.5 should not be private"},
		{net.ParseIP("81.185.175.46"), false, "81.185.175.46 should not be private"},
	}

	for _, ts := range cases {
		if IsPrivateIP(ts.ip) != ts.isPrivate {
			t.Error(ts.msg)
		}
	}
}

func TestIsPrivateIpV6(t *testing.T) {
	cases := []struct {
		ip        net.IP
		isPrivate bool
		msg       string
	}{
		{net.ParseIP("::0"), true, "::0 should be private"},
		{net.ParseIP("::1"), true, "::1 should be private"},
		{net.ParseIP("fe80::1"), true, "fe80::1 should be private"},
		{net.ParseIP("febf::1"), true, "febf::1 should be private"},
		{net.ParseIP("ff00::1"), true, "ff00::1 should be private"},
		{net.ParseIP("ff10::1"), true, "ff10::1 should be private"},
		{net.ParseIP("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff"), true, "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff should be private"},
		{net.ParseIP("2002::"), true, "2002:: should be private"},
		{net.ParseIP("2002:ffff:ffff:ffff:ffff:ffff:ffff:ffff"), true, "2002:ffff:ffff:ffff:ffff:ffff:ffff:ffff should be private"},

		{net.ParseIP("::2"), false, "::2 should not be private"},
		{net.ParseIP("fec0::1"), false, "fec0::1 should not be private"},
		{net.ParseIP("feff::1"), false, "feff::1 should not be private"},
		{net.ParseIP("2409:4073:f:11b9:817a:81ea:a128:1955"), false, "2409:4073:f:11b9:817a:81ea:a128:1955 should not be private"},
		{net.ParseIP("2401:4900:1a8d:3eaf:84b5:5538:efd:5586"), false, "2401:4900:1a8d:3eaf:84b5:5538:efd:5586 should not be private"},
		{net.ParseIP("2409:4063:230c:e172:dd1a:4c74:de7c:cc21"), false, "2409:4063:230c:e172:dd1a:4c74:de7c:cc21 should not be private"},
		{net.ParseIP("2601:240:a:1bc6:9c39:66f0:ced8:d643"), false, "2601:240:a:1bc6:9c39:66f0:ced8:d643 should not be private"},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.isPrivate, IsPrivateIP(tc.ip), tc.ip)
	}
}
