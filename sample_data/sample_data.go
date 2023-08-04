package sample_data

import "github.com/AdithyanMS/mta-hosting-optimizer/models"

var sampleData = []models.IpConfig{
	// I ve kept these unkeyed since this is for demonstration purposes
	{Ip: "127.0.0.1", HostName: "mta-prod-1", Active: true},
	{Ip: "127.0.0.2", HostName: "mta-prod-1", Active: false},
	{Ip: "127.0.0.3", HostName: "mta-prod-2", Active: true},
	{Ip: "127.0.0.4", HostName: "mta-prod-2", Active: true},
	{Ip: "127.0.0.5", HostName: "mta-prod-2", Active: false},
	{Ip: "127.0.0.6", HostName: "mta-prod-3", Active: false},
}

func GetSampleData() []models.IpConfig {
	return sampleData
}
