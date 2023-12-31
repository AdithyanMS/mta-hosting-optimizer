package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AdithyanMS/mta-hosting-optimizer/sample_data"
	"github.com/AdithyanMS/mta-hosting-optimizer/service"
)

func InefficientHosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	data := sample_data.GetSampleData()
	hosts := service.GetInefficientHosts(data)

	res, _ := json.Marshal(hosts)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(res)
	if err != nil {
		log.Println("error writing response")
	}
	return
}
