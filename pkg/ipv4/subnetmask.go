package ipv4

import (
	"encoding/binary"
	"fmt"
	"net"
)

func GetSubnetMask(cidrNumber byte) string {
	subnetMask := cidrNumberToSubnetMask(cidrNumber)

	return fmt.Sprint(net.IP(subnetMask))
}

func cidrNumberToSubnetMask(cidrNumber byte) []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, 0xFFFFFFFF<<(32-cidrNumber))
	return bs
}
