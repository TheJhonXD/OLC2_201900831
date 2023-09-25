package main

import (
	"Server/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(server.CorsMiddleware)
	router.HandleFunc("/", server.Index)
	router.HandleFunc("/run", server.InputOutput)
	router.HandleFunc("/tree", server.TreeParse)

	/* Muestra el la ruta del servidor */
	fmt.Println("*****************************************************************")
	fmt.Println("*\t\t\t\t\t\t\t\t*")
	fmt.Println("*\tServidor corriendo en http://localhost:3000/ \t\t*")
	fmt.Println("*\t\t\t\t\t\t\t\t*")
	fmt.Println("*****************************************************************")
	/* Inicia el servidor */
	log.Fatal(http.ListenAndServe(":3000", router))
}
