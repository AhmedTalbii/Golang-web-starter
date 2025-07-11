package controllers

import (
	"net/http"

	"Templates/config"
	"Templates/controllers/render"
)

func HelpController(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, r, config.PathPages+"help.html")
}
