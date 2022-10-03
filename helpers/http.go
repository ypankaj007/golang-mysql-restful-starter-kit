package helpers

import (
	"encoding/json"
	"net/http"
)

func SendRespond(w http.ResponseWriter, code int, payload interface{}, message string) {
	respData := map[string]interface{}{
		"data":    payload,
		"message": message,
	}
	response, _ := json.Marshal(respData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
