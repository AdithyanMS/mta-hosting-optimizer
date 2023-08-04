package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type InefficientHostsTestCase struct {
	name       string
	minIpCount any
	method     string
	expected   int
}

func TestInefficientHosts(t *testing.T) {
	initial := os.Getenv("MIN_IP_COUNT")
	t.Cleanup(func() {
		os.Setenv("MIN_IP_COUNT", initial)
	})

	tests := []InefficientHostsTestCase{
		{
			name:       "test1",
			minIpCount: 0,
			method:     "GET",
			expected:   200,
		},
		{
			name:       "test2",
			minIpCount: "1",
			method:     "POST",
			expected:   405,
		},
		{
			name:       "test3",
			minIpCount: "",
			expected:   200,
		},
		{
			name:       "test4",
			minIpCount: "2",
			method:     "GET",
			expected:   200,
		},
		{
			name:       "test5",
			minIpCount: false,
			method:     "GET",
			expected:   200,
		},
	}

	for _, test := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			resRecorder := httptest.NewRecorder()
			var response []string
			os.Setenv("MIN_IP_COUNT", fmt.Sprintf("%v", test.minIpCount))

			req, err := http.NewRequest(test.method, "/inefficient_hosts", nil)
			if err != nil {
				t.Fatalf("Error creating a request. error: %v", err)
			}

			InefficientHosts(resRecorder, req)

			assert.Equal(t, test.expected, resRecorder.Code, fmt.Sprintf("Expected status code %v", test.expected))
			if resRecorder.Code == 200 {
				err = json.Unmarshal(resRecorder.Body.Bytes(), &response)
				if err != nil {
					t.Errorf("Response is not a slice of string; resp: %v. error: %s", response, err)
				}
			}
		})
	}

}
