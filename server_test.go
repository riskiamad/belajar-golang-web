package belajar_golang_web

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
