package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

var langTypes []string

func init() {
	// Initialize language type list.
	langTypes = strings.Split(beego.AppConfig.String("lang_types"), "|")

	// Load locale files according to language types.
	for _, lang := range langTypes {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file:", err)
			return
		}
	}
}

type baseControler struct {
	beego.Controller
	i18n.Locale
}

func (_this *baseControler) Prepare() {
	_this.Lang = ""
	acceptLanguage := _this.Ctx.Request.Header.Get("Accept-Language")
	if len(acceptLanguage) > 4 {
		acceptLanguage = acceptLanguage[:5]
		if i18n.IsExist(acceptLanguage) {
			_this.Lang = acceptLanguage
		}
	}

	if len(_this.Lang) == 0 {
		_this.Lang = "en-US"
	}

	_this.Data["Lang"] = _this.Lang

}
