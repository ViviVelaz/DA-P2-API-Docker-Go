package routes

import (
	"encoding/json"
	"net/http"

	"github.com/ViviVelaz/api_go/db"
	"github.com/ViviVelaz/api_go/models"
	"github.com/gorilla/mux"
)

func GetIcesHandler(w http.ResponseWriter, r *http.Request) {
	var ices []models.IceCream
	db.DB.Find(&ices)
	json.NewEncoder(w).Encode(&ices)
}

func GetIceHandler(w http.ResponseWriter, r *http.Request) {
	var ice models.IceCream
	params := mux.Vars(r)

	db.DB.First(&ice, params["id"])

	if ice.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No existe el registro"))
		return
	}

	json.NewEncoder(w).Encode(&ice)
}

func PostIceHandler(w http.ResponseWriter, r *http.Request) {
	var ice models.IceCream

	json.NewDecoder(r.Body).Decode(&ice)

	createdIce := db.DB.Create(&ice)
	err := createdIce.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&ice)
}

func UpdateIceHandler(w http.ResponseWriter, r *http.Request) {
	var ice models.IceCream
	var iceup models.IceCream
	params := mux.Vars(r)

	db.DB.First(&ice, params["id"])

	if ice.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No existe el registro"))
		return
	}

	json.NewDecoder(r.Body).Decode(&iceup)

	db.DB.Model(&ice).Updates(&iceup)

	json.NewEncoder(w).Encode(&iceup)
}

func DeleteIceHandler(w http.ResponseWriter, r *http.Request) {
	var ice models.IceCream
	params := mux.Vars(r)

	db.DB.First(&ice, params["id"])

	if ice.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No existe el registro"))
		return
	}

	db.DB.Unscoped().Delete(&ice)
	w.WriteHeader(http.StatusOK)
}
