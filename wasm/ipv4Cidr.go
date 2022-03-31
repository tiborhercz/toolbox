package main

import (
	"encoding/json"
	"fmt"
	"github.com/tiborhercz/cli-toolbox/internal/utils"
	"github.com/tiborhercz/cli-toolbox/pkg/ipv4"
	"syscall/js"
)

type Ipv4CidrResponse struct {
	SubnetMask       string `json:"subnetMask"`
	FirstIp          string `json:"firstIp"`
	LastIp           string `json:"lastIp"`
	TotalIpAddresses string `json:"totalIpAddresses"`
}

func processIpv4(this js.Value, args []js.Value) any {
	response := Ipv4CidrResponse{}
	ipv4Address := args[0].String()
	fmt.Println(ipv4Address)

	parsedIpAddress, parsedIpNetAddress := utils.ParseIpAddress(ipv4Address)
	cidrNumber := utils.GetCidrNumberFromIp(parsedIpNetAddress.String())

	if parsedIpAddress != nil {
		firstIp, lastIp := ipv4.GetFirstLastIp(parsedIpAddress, cidrNumber)
		subnetMask := ipv4.GetSubnetMask(cidrNumber)
		totalIpAddresses := fmt.Sprint(ipv4.GetTotalCidrIpAddresses(cidrNumber))

		response = Ipv4CidrResponse{
			SubnetMask:       subnetMask,
			FirstIp:          firstIp,
			LastIp:           lastIp,
			TotalIpAddresses: totalIpAddresses,
		}
	}

	jsonData, err := json.Marshal(response)

	if err != nil {
		fmt.Println("Unable to convert the struct to a JSON string")
	} else {
		fmt.Println(string(jsonData))
	}

	return string(jsonData)
}
