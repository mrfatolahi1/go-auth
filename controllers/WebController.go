package controllers

import (
	"fmt"
	"net/http"
)

type WebController struct {
}

func handleFunc(writer http.ResponseWriter, request *http.Request) {
	// Setting the Content-Type header
	writer.Header().Set("Content-Type", "text/html")

	// Writing Response according to the routes
	if request.URL.Path == "/" || request.URL.Path == "/SignUp" {
		fmt.Println("Salam!")
	} else if request.URL.Path == "/SignIn" {

	} else if request.URL.Path == "/SignOut" {

	} else if request.URL.Path == "/UserInfo" {

	} else {

	}

	fmt.Fprintf(writer, "</center>")
}

func (webController WebController) Start() {

	mux := &http.ServeMux{}

	mux.HandleFunc("/", handleFunc)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}
