package service

import (
	"os"
	"strconv"

	"github.com/AdithyanMS/mta-hosting-optimizer/models"
)

func minimumIpAddresses() int {
	xString := os.Getenv("MIN_IP_COUNT")
	x, err := strconv.Atoi(xString)
	if err != nil {
		x = 1
	}
	return x
}

func GetInefficientHosts(data []models.IpConfig) []string {
	activeIPs := make(map[string]int)
	inefficientHosts := []string{}

	for _, host := range data {
		if host.Active {
			activeIPs[host.HostName]++
		} else {
			_, present := activeIPs[host.HostName]
			if !present {
				activeIPs[host.HostName] = 0
			}
		}
	}

	// make global
	x := minimumIpAddresses()

	for hostName, activeNum := range activeIPs {
		if activeNum <= x {
			inefficientHosts = append(inefficientHosts, hostName)
		}
	}

	return inefficientHosts

}
