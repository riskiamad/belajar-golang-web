package belajar_golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var dirTemplates embed.FS

var myTemplates = template.Must(template.ParseFS(dirTemplates, "templates/*.gohtml"))

func TemplateCache(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "simple.gohtml", "Cara ini lebih cepat karena hanya melakukan parse 1 kali di awal")
}

func TestTemplateCache(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCache(recorder, request)

	body, _ := ioutil.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
