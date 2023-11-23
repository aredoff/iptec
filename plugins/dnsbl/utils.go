package dnsbl

import (
	"encoding/hex"
	"errors"
	"fmt"
	"net"
)

func reverseIP(address net.IP) (string, error) {
	if address == nil {
		return "", errors.New("cant revers IP, ip is null")
	}
	segments := address.To4()
	if segments == nil {
		dst := make([]byte, hex.EncodedLen(len(address)))
		_ = hex.Encode(dst, address)
		i := 31
		reversIpv6 := ""
		for i >= 0 {
			reversIpv6 += string(dst[i]) + "."
			i--
		}
		return reversIpv6[:len(reversIpv6)-1], nil
	}
	return fmt.Sprintf("%d.%d.%d.%d", segments[3], segments[2], segments[1], segments[0]), nil
}
