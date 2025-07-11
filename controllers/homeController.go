package controllers

import (
	"Templates/config"
	"Templates/controllers/render"
	"net/http"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w,r,config.PathPages+"home.html")
}
