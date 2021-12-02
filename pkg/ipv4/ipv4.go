package ipv4

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

type IP4 uint32

type IP4Chan chan IP4

const allOnes = IP4(0xFFFFFFFF)

func GetTotalCidrIpAddresses(cidrNumber byte) uint32 {
	return ipCount(cidrNumber)
}

func GetFirstLastIp(ipAddress net.IP, cidrNumber byte) (string, string) {
	firstIp := minIP(ipv4ToBinary(ipAddress), cidrNumber)
	lastIp := maxIP(ipv4ToBinary(ipAddress), cidrNumber)

	return firstIp.String(), lastIp.String()
}

func (f IP4Chan) Next() *IP4 {
	c, ok := <-f
	if !ok {
		return nil
	}
	return &c
}

func IP4s(ip IP4, cidrNumber byte) IP4Chan {
	c := make(chan IP4)
	a := minIP(ip, cidrNumber)
	limit := ipCount(cidrNumber)
	fmt.Println("limit")
	fmt.Println(limit)
	go func() {
		for {
			if limit == 0 {
				close(c)
				return
			}
			c <- a
			a = IP4(uint32(a) + 1)
			limit--
		}
	}()
	return c
}

func (ip IP4) To4() [4]byte {
	var result [4]byte
	binary.BigEndian.PutUint32(result[:], uint32(ip))
	return result
}

func (ip IP4) String() string {
	arr := ip.To4()
	return fmt.Sprintf("%v.%v.%v.%v", arr[0], arr[1], arr[2], arr[3])
}

func ipv4ToBinary(ipAddress net.IP) IP4 {
	var IPv4 IP4
	binary.Read(bytes.NewBuffer(net.ParseIP(ipAddress.String()).To4()), binary.BigEndian, &IPv4)

	return IPv4
}

func getMask(value byte) IP4 {
	return IP4(allOnes << (32 - value))
}

func inverse(ip IP4) IP4 {
	return IP4(uint32(allOnes) ^ uint32(ip))
}

func newIP4(a, b, c, d byte) IP4 {
	return IP4(uint32(a)<<24 | uint32(b)<<16 | uint32(c)<<8 | uint32(d))
}

func minIP(ip IP4, cidrNumber byte) IP4 {
	return IP4(ip & getMask(cidrNumber))
}

func maxIP(ip IP4, cidrNumber byte) IP4 {
	var mask = getMask(cidrNumber)
	return IP4(ip&mask | inverse(mask))
}

func ipCount(cidrNumber byte) uint32 {
	return 0xFFFFFFFF ^ uint32(inverse(getMask(cidrNumber))) + 1
}
