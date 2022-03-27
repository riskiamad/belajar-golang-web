package belajar_golang_web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	fmt.Fprint(w, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/", nil)
	request.Header.Add("Content-type", "application/json")

	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)
	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))
}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("x-powered-by", "iamrisk")
	fmt.Fprint(w, "OK")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/", nil)

	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)
	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.Header.Get("x-powered-by"))
}
