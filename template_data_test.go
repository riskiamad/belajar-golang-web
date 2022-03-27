package belajar_golang_web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Achmad Rizky",
		"Adress": map[string]interface{}{
			"Street": "Jalan belum ada",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := ioutil.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

type Adress struct {
	Street string
}

type Page struct {
	Title  string
	Name   string
	Adress Adress
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", Page{
		Title: "Template Data Map",
		Name:  "Achmad Rizky",
		Adress: Adress{
			Street: "Jalan belum ada",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	body, _ := ioutil.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
