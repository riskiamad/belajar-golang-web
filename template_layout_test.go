package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/base.gohtml",
		"./templates/content.gohtml",
	))
	t.ExecuteTemplate(w, "content", map[string]interface{}{
		"Title": "Template Data Layout",
		"Name":  "Achmad Rizky",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := ioutil.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
