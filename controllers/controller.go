package controllers

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
    "desafio/database"
    "desafio/models"
    
)

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Home Page")
}

func DisplayAllAnimals(w http.ResponseWriter, r *http.Request) {
    var p []models.Animal
    database.DB.Find(&p)
    json.NewEncoder(w).Encode(p)
}

func GetAnimals(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    var a models.Animal
    database.DB.First(&a, id)
    json.NewEncoder(w).Encode(a)
}

func CreateAnimal(w http.ResponseWriter, r *http.Request) {
	var newAnimal models.Animal
	json.NewDecoder(r.Body).Decode(&newAnimal)
	database.DB.Create(&newAnimal)
	json.NewEncoder(w).Encode(newAnimal)
}

func DeleteAnimal(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    id := vars["id"]
    var a models.Animal    
	database.DB.Delete(&a, id)
	json.NewEncoder(w).Encode(a)
}

func UpdateAnimal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    id := vars["id"]
    var a models.Animal
    database.DB.First(&a, id)
    json.NewDecoder(r.Body).Decode(&a)
    database.DB.Save(&a)
    json.NewEncoder(w).Encode(a)
}

//PARA BAIXO É ACTIONS

func DisplayAllActions(w http.ResponseWriter, r *http.Request) {
    var p []models.Action
    database.DB.Find(&p)
    json.NewEncoder(w).Encode(p)
}

//------------------------------

// controllers/ActionController.go
func DisplayFlockActions(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["flock"]
    var actions []models.Action
    resultDb := database.DB.Where("flock = ?", id).Find(&actions)
    fmt.Println(resultDb, actions, id)
    json.NewEncoder(w).Encode(actions)
}

//-----------------------------

func GetActions(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    var a models.Action
    database.DB.First(&a, id)
    json.NewEncoder(w).Encode(a)
}

func GetActionsFlock(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.Header().Add("Content-Type", "application/json")
    id := vars["flock"]
    var a models.Action
    database.DB.First(&a, id)
    json.NewEncoder(w).Encode(a)
}

func CreateActions(w http.ResponseWriter, r *http.Request) {
	var newActions models.Action
	json.NewDecoder(r.Body).Decode(&newActions)
	database.DB.Create(&newActions)
	json.NewEncoder(w).Encode(newActions)
}

func DeleteActions(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    id := vars["id"]
    var a models.Action    
	database.DB.Delete(&a, id)
	json.NewEncoder(w).Encode(a)
}

func UpdateActions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    id := vars["id"]
    var b models.Action
    database.DB.First(&b, id)
    json.NewDecoder(r.Body).Decode(&b)
    database.DB.Save(&b)
    json.NewEncoder(w).Encode(b)
}

// PARA BAIXO É REPORT 

func DisplayAllReports(w http.ResponseWriter, r *http.Request) {
    var p []models.Report
    database.DB.Find(&p)
    json.NewEncoder(w).Encode(p)
}

func GetReports(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    var a models.Report
    database.DB.First(&a, id)
    json.NewEncoder(w).Encode(a)
}

func CreateReports(w http.ResponseWriter, r *http.Request) {
	var newReports models.Report
	json.NewDecoder(r.Body).Decode(&newReports)
	database.DB.Create(&newReports)
	json.NewEncoder(w).Encode(newReports)
}

func DeleteReports(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    id := vars["id"]
    var a models.Report    
	database.DB.Delete(&a, id)
	json.NewEncoder(w).Encode(a)
}

func UpdateReports(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    id := vars["id"]
    var c models.Report
    database.DB.First(&c, id)
    json.NewDecoder(r.Body).Decode(&c)
    database.DB.Save(&c)
    json.NewEncoder(w).Encode(c)
}