package controllers

type MainController struct {
	webController WebController
}

func (mainController MainController) Start() {
	mainController.webController.Start()
}
