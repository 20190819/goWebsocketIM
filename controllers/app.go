package controllers

type AppController struct {
	baseControler
}

func (_this *AppController) Get() {
	_this.TplName = "welcome.html"
}
