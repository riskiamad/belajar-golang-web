package belajar_golang_web

import (
	_ "embed"
	"io"
	"net/http"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	err := myTemplates.ExecuteTemplate(w, "upload_form.gohtml", nil)
	if err != nil {
		panic(err)
	}
}

func Upload(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	name := r.PostFormValue("name")
	myTemplates.ExecuteTemplate(w, "upload_success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestServerUpload(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// //go:embed resources/pngwing.com.png
// var image []byte

// func TestUpload(t *testing.T) {
// 	body := new(bytes.Buffer)
// 	writer := multipart.NewWriter(body)
// 	writer.WriteField("name", "Achmad Rizky")

// 	file, _ := writer.CreateFormFile("file", "contohUpload.png") //please insert sample name of your image
// 	file.Write(image)
// 	writer.Close()

// 	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
// 	request.Header.Set("Content-Type", writer.FormDataContentType())
// 	recorder := httptest.NewRecorder()

// 	Upload(recorder, request)

// 	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
// 	fmt.Println(string(bodyResponse))
// }
