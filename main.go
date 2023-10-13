package main

import (
	"net/http"

	"github.com/ViviVelaz/api_go/db"
	"github.com/ViviVelaz/api_go/models"
	"github.com/ViviVelaz/api_go/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()

	db.DB.AutoMigrate(models.IceCream{})

	r := mux.NewRouter()

	r.HandleFunc("/icecreams", routes.GetIcesHandler).Methods("GET")
	r.HandleFunc("/icecreams/{id}", routes.GetIceHandler).Methods("GET")
	r.HandleFunc("/icecreams", routes.PostIceHandler).Methods("POST")
	r.HandleFunc("/icecreams/{id}", routes.DeleteIceHandler).Methods("DELETE")
	r.HandleFunc("/icecreams/{id}", routes.UpdateIceHandler).Methods("PUT")

	http.ListenAndServe(":8080", r)
}
