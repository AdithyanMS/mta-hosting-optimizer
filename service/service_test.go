package service

import (
	"fmt"
	"os"
	"sort"
	"testing"

	"github.com/AdithyanMS/mta-hosting-optimizer/models"
	"github.com/stretchr/testify/assert"
)

type InefficientHostsConfigTest struct {
	name     string
	data     []models.IpConfig
	expected []string
}

type MinIpCountTest struct {
	name     string
	data     any
	expected int
}

func TestMinimumIpAddresses(t *testing.T) {
	initial := os.Getenv("MIN_IP_COUNT")
	t.Cleanup(func() {
		os.Setenv("MIN_IP_COUNT", initial)
	})

	// expectation is that the function should return an integer if it finds the value in .env, otherwise it should return 1 in all other cases
	tests := []MinIpCountTest{
		{
			name:     "test 1",
			data:     0,
			expected: 0,
		},
		{
			name:     "test 2",
			data:     "",
			expected: 1,
		},
		{
			name:     "test 3",
			data:     4,
			expected: 4,
		},
		{
			name:     "test 4",
			data:     1,
			expected: 1,
		},
		{
			name:     "test 5",
			data:     3,
			expected: 3,
		},
		{
			name:     "test 6",
			data:     "Adi",
			expected: 1,
		},
		{
			name:     "test 7",
			data:     true,
			expected: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val := fmt.Sprintf("%v", test.data)
			os.Setenv("MIN_IP_COUNT", val)
			minIpCount := minimumIpAddresses()
			assert.Equal(t, test.expected, minIpCount, fmt.Sprintf("failed %s; expected %v; got %v", test.name, test.expected, minIpCount))
		})
	}
}

func TestGetInefficientHosts(t *testing.T) {
	initial := os.Getenv("MIN_IP_COUNT")
	t.Cleanup(func() {
		os.Setenv("MIN_IP_COUNT", initial)
	})
	assert := assert.New(t)

	// the following tests are run based on different values of Active and different values of the environment variable
	os.Setenv("MIN_IP_COUNT", "1")
	tests := []InefficientHostsConfigTest{
		{
			name: "test 1",
			data: []models.IpConfig{
				{Ip: "127.0.0.1", HostName: "mta-prod-1", Active: true},
				{Ip: "127.0.0.2", HostName: "mta-prod-1", Active: false},
				{Ip: "127.0.0.3", HostName: "mta-prod-2", Active: true},
				{Ip: "127.0.0.4", HostName: "mta-prod-2", Active: true},
				{Ip: "127.0.0.5", HostName: "mta-prod-2", Active: false},
				{Ip: "127.0.0.6", HostName: "mta-prod-3", Active: false},
			},
			expected: []string{
				"mta-prod-3",
				"mta-prod-1",
			},
		},
		{
			name: "test 2",
			data: []models.IpConfig{
				{Ip: "127.0.0.1", HostName: "mta-prod-1", Active: true},
				{Ip: "127.0.0.2", HostName: "mta-prod-1", Active: true},
				{Ip: "127.0.0.3", HostName: "mta-prod-2", Active: true},
				{Ip: "127.0.0.4", HostName: "mta-prod-2", Active: true},
				{Ip: "127.0.0.5", HostName: "mta-prod-2", Active: false},
				{Ip: "127.0.0.6", HostName: "mta-prod-3", Active: false},
			},
			expected: []string{
				"mta-prod-3",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := GetInefficientHosts(test.data)
			sort.Strings(response)
			sort.Strings(test.expected)
			assert.Equal(test.expected, response, fmt.Sprintf("failed %s; expected %v; got %v", test.name, test.expected, response))
		})
	}

	os.Setenv("MIN_IP_COUNT", "0")
	tests = []InefficientHostsConfigTest{
		{
			name: "test 3",
			data: []models.IpConfig{
				{Ip: "127.0.0.1", HostName: "mta-prod-1", Active: true},
				{Ip: "127.0.0.2", HostName: "mta-prod-1", Active: false},
				{Ip: "127.0.0.3", HostName: "mta-prod-2", Active: true},
				{Ip: "127.0.0.4", HostName: "mta-prod-2", Active: true},
				{Ip: "127.0.0.5", HostName: "mta-prod-2", Active: false},
				{Ip: "127.0.0.6", HostName: "mta-prod-3", Active: false},
			},
			expected: []string{
				"mta-prod-3",
			},
		},
		{
			name: "test 4",
			data: []models.IpConfig{
				{Ip: "127.0.0.1", HostName: "mta-prod-1", Active: true},
				{Ip: "127.0.0.2", HostName: "mta-prod-1", Active: false},
				{Ip: "127.0.0.3", HostName: "mta-prod-2", Active: true},
				{Ip: "127.0.0.4", HostName: "mta-prod-2", Active: true},
				{Ip: "127.0.0.5", HostName: "mta-prod-2", Active: false},
				{Ip: "127.0.0.6", HostName: "mta-prod-3", Active: true},
			},
			expected: []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := GetInefficientHosts(test.data)
			sort.Strings(response)
			sort.Strings(test.expected)
			assert.Equal(test.expected, response, fmt.Sprintf("failed %s; expected %v; got %v", test.name, test.expected, response))
		})
	}

	os.Setenv("MIN_IP_COUNT", "2")
	tests = []InefficientHostsConfigTest{
		{
			name: "test 5",
			data: []models.IpConfig{
				{Ip: "127.0.0.1", HostName: "mta-prod-1", Active: true},
				{Ip: "127.0.0.2", HostName: "mta-prod-1", Active: false},
				{Ip: "127.0.0.3", HostName: "mta-prod-2", Active: true},
				{Ip: "127.0.0.4", HostName: "mta-prod-2", Active: true},
				{Ip: "127.0.0.5", HostName: "mta-prod-2", Active: false},
				{Ip: "127.0.0.6", HostName: "mta-prod-3", Active: false},
			},
			expected: []string{
				"mta-prod-3",
				"mta-prod-1",
				"mta-prod-2",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := GetInefficientHosts(test.data)
			sort.Strings(response)
			sort.Strings(test.expected)
			assert.Equal(test.expected, response, fmt.Sprintf("failed %s; expected %v; got %v", test.name, test.expected, response))
		})
	}
}
