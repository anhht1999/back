package server

import (
	"fmt"
	"net/http"

	"../handlers"
)

func RunServer() {
	fmt.Println("Starting server. Please open http://localhost:8080")
	defer func() {
		fmt.Println("Server is stopped")
	}()

	http.HandleFunc("/sum", handlers.Sum)

	http.HandleFunc("/minus", handlers.Minus)

	http.HandleFunc("/multi", handlers.Multi)

	http.HandleFunc("/div", handlers.Div)

	http.HandleFunc("/", handlers.Handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("Error when running server")
	}
}
