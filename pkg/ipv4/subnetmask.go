package ipv4

import (
	"encoding/binary"
	"math"
	"math/bits"
	"strconv"
	"strings"
)

func subnetMaskToCidr(bs []byte) int {
	addr := binary.BigEndian.Uint32(bs[0:])
	return bits.OnesCount32(addr)
}

func prefixNumberToSubnetMask(ipPrefixNumber byte) []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, 0xFFFFFFFF<<(32-ipPrefixNumber))
	return bs
}

func GetSubnetMask(ipPrefixNumber byte) string {
	subnetMask := []string{"255", "255", "255", "255"}
	// Calculate the octet position also known as the subnet mask class. Subnet mask is a 4 octet number.
	octetPos := int(math.Floor(float64(ipPrefixNumber) / float64(8))) // Floor used to match array indexes

	for i, _ := range subnetMask {
		// Part of the subnet mask must remain 255
		if i < octetPos {
			continue
		}

		// Part of the subnet mask should be calculated
		if i == octetPos {
			subnetNumber := int(32 - ipPrefixNumber) // Subtract prefix from /32
			subnetMask[i] = calculateSubnetMaskNumber(subnetNumber)
		}

		// Part of the subnet mask should be set to 0
		if i > octetPos {
			subnetMask[i] = "0"
		}
	}

	return strings.Join(subnetMask, ".")
}

// calculateSubnetMaskNumber Number is calculated from the subnetNumber.
func calculateSubnetMaskNumber(subnetNumber int) string {
	var (
		bitAmount = 8 // One part of a subnet mask is 8 bits
		oneCount int
	)
	// One part of the subnet mask is 8 bits.
	// zeroCount is the number of turned off bits zeros
	zeroCount := subnetNumber % bitAmount

	// If zeroCount count is 0 all bits are turned off so all 8 bits are all zeros
	if zeroCount == 0 {
		zeroCount = bitAmount
	}

	// Calculate the amount of enabled bits (1)
	oneCount = bitAmount - zeroCount

	// Create the binary string from oneCount and zeroCount
	binaryString := strings.Repeat("1", oneCount) + strings.Repeat("0", zeroCount)

	// Transform binaryString to integer
	maskNumber, _ := strconv.ParseInt(binaryString, 2, 64)

	return strconv.Itoa(int(maskNumber))
}
