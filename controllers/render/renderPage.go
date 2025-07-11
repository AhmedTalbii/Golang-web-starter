package render

import (
	"net/http"
	"text/template"

	"Templates/config"
)

func RenderPage(w http.ResponseWriter, r *http.Request, Path string) {
	layout := template.Must(template.ParseFiles(
		config.PathPages+"layout.html",
		config.PathSections+"navbar.html",
		Path,
	))
	errRendering := layout.ExecuteTemplate(w, "Layout", nil)
	if errRendering != nil {
		http.Error(w, "Template error: "+errRendering.Error(), http.StatusInternalServerError)
		return
	}
}
