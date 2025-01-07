package Server

import (
	"fmt"
	"net/http"
)

func StartServer() {

	// Gestion des fichiers Static et configuration de la route pour y acceder
	fileServer := http.FileServer(http.Dir("Web/Static"))
	http.Handle("/Static/", http.StripPrefix("/Static/", fileServer))

	// Initialisation des routes
	Routes()

	// Lancement du serveur
	fmt.Println("LOG: Server started on : http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("LOG: ERROR while starting the server on port ':8080' ", err)
		return
	}
}
