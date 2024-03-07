package main

import (
	"encoding/json"
	"net/http"
)

type HelloHandler struct{}

func (HelloHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	jsonEncoder := json.NewEncoder(writer)

	err := jsonEncoder.Encode(HelloResponse{
		Message: "Hello, World!",
	})
	if err != nil {
		panic(err)
	}
}

type HelloResponse struct {
	Message string
}
