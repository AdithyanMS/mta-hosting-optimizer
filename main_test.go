package main

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	go main()
	time.Sleep(time.Second * 2)

	resp, err := http.Get("http://localhost:8081/inefficient_hosts")
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Server not responding with 200")
}
