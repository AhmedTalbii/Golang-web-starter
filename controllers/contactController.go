package controllers

import (
	"net/http"

	"Templates/config"
	"Templates/controllers/render"
)

func ContactController(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, r, config.PathPages+"contact.html")
}
