package controllers

import (
	"net/http"

	"Templates/config"
	"Templates/controllers/render"
)

func AboutController(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, r, config.PathPages+"about.html")
}
