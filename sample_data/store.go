package sample_data

import "github.com/AdithyanMS/mta-hosting-optimizer/models"

func GetSampleData() []models.IpConfig {
	var sampleData = []models.IpConfig{
		{Ip: "127.0.0.1", HostName: "mta-prod-1", Active: true},
		{Ip: "127.0.0.2", HostName: "mta-prod-1", Active: false},
		{Ip: "127.0.0.3", HostName: "mta-prod-2", Active: true},
		{Ip: "127.0.0.4", HostName: "mta-prod-2", Active: true},
		{Ip: "127.0.0.5", HostName: "mta-prod-2", Active: false},
		{Ip: "127.0.0.6", HostName: "mta-prod-3", Active: false},
	}
	return sampleData
}
