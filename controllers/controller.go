package controllers

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
    "desafio/database"
    "desafio/models"
    "encoding/base64"
    "io/ioutil"
   )
// Home page
func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Home Page")
}

// função para GET de todos os animals
func DisplayAllAnimals(w http.ResponseWriter, r *http.Request) {
    var p []models.Animal
    database.DB.Find(&p)
    json.NewEncoder(w).Encode(p)
}
//----------------------------------
// filtro por tipo bovino/suino
func DisplayFlockAnimal(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["flock"]
    var animals []models.Animal
    database.DB.Where("flock = ?", id).Find(&animals)
    json.NewEncoder(w).Encode(animals)
}
//----------------------------------

// filtro por ID do animal
func GetAnimals(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    var a models.Animal
    database.DB.First(&a, id)
    json.NewEncoder(w).Encode(a)
}

// função para POST criar animal novo
func CreateAnimal(w http.ResponseWriter, r *http.Request) {
	var newAnimal models.Animal
	json.NewDecoder(r.Body).Decode(&newAnimal)
	database.DB.Create(&newAnimal)
	json.NewEncoder(w).Encode(newAnimal)
}

//função de DELETE produto animal
func DeleteAnimal(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    id := vars["id"]
    var a models.Animal    
	database.DB.Delete(&a, id)
	json.NewEncoder(w).Encode(a)
}

// função para atualizar/editar animal
func UpdateAnimal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    id := vars["id"]
    var a models.Animal
    database.DB.First(&a, id)
    json.NewDecoder(r.Body).Decode(&a)
    database.DB.Save(&a)
    json.NewEncoder(w).Encode(a)
}

//PARA BAIXO É TODAS AS FUNÇÕES ACTIONS
// Função para mostrar/GET todos as açoes 
func DisplayAllActions(w http.ResponseWriter, r *http.Request) {
    var p []models.Action
    database.DB.Find(&p)
    json.NewEncoder(w).Encode(p)
}
//------------------------------
// filtro por tipo do animal como bovino/suino
func DisplayFlockActions(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["flock"]
    var actions []models.Action
    database.DB.Where("flock = ?", id).Find(&actions)
    json.NewEncoder(w).Encode(actions)
}
//-----------------------------
//função para filtrar por ID
func GetActions(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    var a models.Action
    database.DB.First(&a, id)
    json.NewEncoder(w).Encode(a)
}
//função para criar action
func CreateActions(w http.ResponseWriter, r *http.Request) {
	var newActions models.Action
	json.NewDecoder(r.Body).Decode(&newActions)
	database.DB.Create(&newActions)
	json.NewEncoder(w).Encode(newActions)
}
//função para deletar action
func DeleteActions(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    id := vars["id"]
    var a models.Action    
	database.DB.Delete(&a, id)
	json.NewEncoder(w).Encode(a)
}
//função para editar action
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
//função para listar todos os reports sem filtro
func DisplayAllReports(w http.ResponseWriter, r *http.Request) {
    var p []models.Report
    database.DB.Find(&p)
    json.NewEncoder(w).Encode(p)
}
//----------------------------------------------------------------
// função para filtrar o tipo com bovino/suino
func DisplayFlockReports(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["code"]
    var reports []models.Report
    database.DB.Where("problemcode = ?", id).Find(&reports)
    json.NewEncoder(w).Encode(reports)
}
//----------------------------------------------------------------
//função para filtrar todos os produtos por ID
func GetReports(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    var a models.Report
    database.DB.First(&a, id)
    json.NewEncoder(w).Encode(a)
}
// função para criar
func CreateReports(w http.ResponseWriter, r *http.Request) {
	var tt models.Report
    json.NewDecoder(r.Body).Decode(&tt)
    database.DB.Create(&tt)
    json.NewEncoder(w).Encode(tt)
}
//função para deletar report
func DeleteReports(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    id := vars["id"]
    var a models.Report    
	database.DB.Delete(&a, id)
	json.NewEncoder(w).Encode(a)
}
//função para editar report
func UpdateReports(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    id := vars["id"]
    var c models.Report
    database.DB.First(&c, id)
    json.NewDecoder(r.Body).Decode(&c)
    database.DB.Save(&c)
    json.NewEncoder(w).Encode(c)
}
//test
func Testimages(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["imageproblem"]
    fmt.Printf(id)
    var actions []models.Report
    database.DB.Where("flock = ?", id).Find(&actions)
    json.NewEncoder(w).Encode(actions)
}

func Teste(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["url"]
    var t []models.Report
    database.DB.Where("imageproblem = ?", id).Find(&t)
   // fmt.Printf(id, t)

    //dencoded := base64.StdEncoding.EncodeToString([]byte(id))
    dencoded, _ := base64.StdEncoding.DecodeString(id)
        // 
        err := ioutil.WriteFile("C:/teste/" + "https://" + string(dencoded) + ".jpg", []byte(dencoded), 0644)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "Failed to save encoded URL: %v", err)
            return
        }

    json.NewEncoder(w).Encode(t)
}

func Testeee(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["url"]
    var t []models.Report
    database.DB.Where("imageproblem = ?", id).Find(&t)
   // fmt.Printf(id, t)

    encoded := base64.StdEncoding.EncodeToString([]byte(id))
        // 
        err := ioutil.WriteFile("C:/teste/" + id + ".jpg", []byte(dencoded), 0644)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "Failed to save encoded URL: %v", err)
            return
        }

    json.NewEncoder(w).Encode(t)
}
