package server

import (
	"Templates/config"
	"Templates/routes"
	"fmt"
	"net/http"
)

func ServerStart() {
	mux := http.NewServeMux()

	routes.RoutesHandle(mux)

	serv := &http.Server{
		Addr: config.Port,
		Handler: mux,
	}

	fmt.Println("Server is running at http://localhost:3000 (:")
	errListening := serv.ListenAndServe()
	if errListening != nil {
		fmt.Println("Error Listening (:")
		return
	}
}