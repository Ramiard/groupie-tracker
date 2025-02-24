package Server

import (
	"fmt"
	"net/http"
)

func StartServer() {

	// 'Static' directory gestion and route configuration to access it
	fileServer := http.FileServer(http.Dir("Web/Static"))
	http.Handle("/Static/", http.StripPrefix("/Static/", fileServer))

	// Routes initialization
	Routes()

	// Starting the server
	fmt.Println("LOG: Server started on : http://localhost:8080")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Println("LOG: ERROR while starting the server on port ':8080' ", err)
		return
	}
}
