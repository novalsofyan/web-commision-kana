// Package jsonresp menyediakan JSON responder
package jsonresp

import (
	"encoding/json"
	"net/http"
)

type JSONResponder interface {
	WriteData(w http.ResponseWriter, status int, data any)
}

type responder struct{}

func NewResponder() JSONResponder {
	return &responder{}
}

func (r *responder) WriteData(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := map[string]any{
		"data": data,
	}

	json.NewEncoder(w).Encode(resp)
}
