package routes

import (
	"net/http"

	"Templates/controllers"
)

func RoutesHandle(mux *http.ServeMux) {
	mux.HandleFunc("/", controllers.HomeController)
	mux.HandleFunc("/about", controllers.AboutController)
	mux.HandleFunc("/contact", controllers.ContactController)
	mux.HandleFunc("/help", controllers.HelpController)
	mux.HandleFunc("/error", controllers.ErrorController)
}
