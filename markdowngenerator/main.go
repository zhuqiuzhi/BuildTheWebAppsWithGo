package main

import (
	"net/http"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"
)

func main() {
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("./public")))

	http.ListenAndServe(":8080", nil)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown:= r.PostFormValue("body")
	if markdown== "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Markdown 内容为空"))
	}

	unsafe := blackfriday.Run([]byte(markdown), blackfriday.WithExtensions(blackfriday.CommonExtensions))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	rw.Write(html)
}