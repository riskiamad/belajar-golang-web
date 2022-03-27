package belajar_golang_web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

type MyPage struct {
	Name string
}

func (mypage MyPage) SayHello(name string) string {
	return "Hello " + name + " My Name is " + mypage.Name
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("FUNCTION").Parse(`{{.SayHello "Budi"}}`)
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "FUNCTION", MyPage{Name: "Achmad Rizky"})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ := ioutil.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{Name: "Achmad Rizky"})
}

func TestTemplateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateGlobal(recorder, request)

	body, _ := ioutil.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateCreateGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse("{{upper .Name}}"))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{Name: "Achmad Rizky"})
}

func TestTemplateCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCreateGlobal(recorder, request)

	body, _ := ioutil.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplatePipelines(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(value string) string {
			return "Hello " + value
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse("{{sayHello .Name | upper}}"))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{Name: "Achmad Rizky"})
}

func TestTemplatePipelines(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplatePipelines(recorder, request)

	body, _ := ioutil.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
