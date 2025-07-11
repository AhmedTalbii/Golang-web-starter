package controllers

import (
	"net/http"

	"Templates/config"
	"Templates/controllers/render"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		render.RenderPage(w, r, config.PathPages+"error.html")
		return
	}
	render.RenderPage(w, r, config.PathPages+"home.html")
}
