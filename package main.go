package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/mini/", serveFiles)
	http.HandleFunc("/avalon/", serveFiles2)
	http.HandleFunc("/greeno/", serveFiles3)
	http.ListenAndServe(":80", nil)

}
func serveFiles(w http.ResponseWriter, r *http.Request) {
	p := "." + r.URL.Path
	if p == "./" {
		p = "./mini/index.html"
	}
	fmt.Println(p)
	http.ServeFile(w, r, p)
}
func serveFiles2(w http.ResponseWriter, r *http.Request) {
	p := "." + r.URL.Path
	if p == "./" {
		p = "./avalon/index.html"
	}
	fmt.Println(p)
	http.ServeFile(w, r, p)
}
func serveFiles3(w http.ResponseWriter, r *http.Request) {
	p := "." + r.URL.Path
	if p == "./" {
		p = "./greeno/index.html"
	}
	fmt.Println(p)
	http.ServeFile(w, r, p)
}
