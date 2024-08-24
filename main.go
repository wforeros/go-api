package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wforeros/go-gorm-restapi/routes"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	PORT := ":3000"
	http.ListenAndServe(PORT, r)
	fmt.Println("Server running on port ", PORT)
}
