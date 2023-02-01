package controllers

import (
	"fmt"
	"net/http"
)

type MainController struct {
}

func handleFunc(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")

	if request.URL.Path == "/" || request.URL.Path == "/SignUp" {
		fmt.Println("Salam!")
	} else if request.URL.Path == "/SignIn" {

	} else if request.URL.Path == "/SignOut" {

	} else if request.URL.Path == "/UserInfo" {

	} else {

	}

	fmt.Fprintf(writer, "</center>")
}

//func signUp

func (mainController MainController) Start() {
	mux := &http.ServeMux{}

	mux.HandleFunc("/", handleFunc)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}
