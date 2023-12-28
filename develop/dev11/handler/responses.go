package handler

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewErrorResponse(error string) ErrorResponse {
	return ErrorResponse{Error: error}
}

type SuccessResponse struct {
	Result interface{} `json:"result"`
}

func NewSuccessResponse(result interface{}) SuccessResponse {
	return SuccessResponse{Result: result}
}

func errorresponse(w http.ResponseWriter, error string) {
	r := NewErrorResponse(error)
	response, err := json.Marshal(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(response)
	return
}

func successresponse(w http.ResponseWriter, result interface{}) {
	r := NewSuccessResponse(result)
	response, err := json.Marshal(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return
}
