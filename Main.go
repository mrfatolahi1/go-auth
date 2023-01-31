package main

import "go-auth/controllers"

func main() {
	mainController := controllers.MainController{}
	mainController.Start()
}
