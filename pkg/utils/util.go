package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(jsonData)
}
