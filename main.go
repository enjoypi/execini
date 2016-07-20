package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/inis/", iniHandler)

	http.ListenAndServe(":8080", nil)
}

func iniHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r.Method)
	fmt.Println(r.PostForm)

	NewIni("template.ini", "newIni")
}
