package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func AutoEscape(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  "<p>Hello World</p>",
	})
}

func TestAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	AutoEscape(recorder, request)

	body, _ := ioutil.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(AutoEscape),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func AutoEscapeDisabled(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  template.HTML("<p>Hello World</p>"),
	})
}

func TestAutoEscapeDisabled(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	AutoEscapeDisabled(recorder, request)

	body, _ := ioutil.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestAutoEscapeDisabledServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(AutoEscapeDisabled),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
