package facecheck

import (
	"encoding/json"
	"net/http"
)

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func FaceCheckRes(w http.ResponseWriter, r *http.Request) {
	resData := r.Context().Value("resData").(map[string]interface{})

	var payload = map[string]interface{}{
		"result": resData["result"],
	}
	respondwithJSON(w, 200, payload)
}
