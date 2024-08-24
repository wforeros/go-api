package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wforeros/go-gorm-restapi/db"
	"github.com/wforeros/go-gorm-restapi/models"
	"github.com/wforeros/go-gorm-restapi/routes"
)

func setup() {
	db.ConnectToDb()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})
}

func main() {
	setup()

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)

	PORT := ":3000"
	http.ListenAndServe(PORT, r)
	fmt.Println("Server running on port ", PORT)
}
