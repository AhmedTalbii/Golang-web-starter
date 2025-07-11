package controllers

import (
	"net/http"

	"Templates/config"
	"Templates/controllers/render"
)

func ErrorController(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, r, config.PathPages+"error.html")
}
