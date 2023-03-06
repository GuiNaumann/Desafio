package routes

import (
	"log"
    "net/http"
	"github.com/gorilla/mux"
	"desafio/controllers"
    "github.com/MadAppGang/httplog"
    "desafio/middleware"
)

//função para lidar com outras funcioses
func HandleRequests() {
	//roteamento ultilizado gorilla mux para as rotas http
    r := mux.NewRouter()
    //flock := r.URL.Query().Get("flock")
    //home page route
    r.HandleFunc("/", controllers.Home)

    //
    r.Use(middleware.ContentTypeMiddleware)

    // log em todas as requisições http ultilizando middleware
    r.Use(httplog.Logger)
    
    //rota para a crud de animals
    r.HandleFunc("/animals", controllers.DisplayAllAnimals).Methods("Get")
    r.HandleFunc("/animals/{id}", controllers.GetAnimals).Methods("Get")
    //preciso criar um get por raça tambemmm
    r.HandleFunc("/animals", controllers.CreateAnimal).Methods("Post")
    r.HandleFunc("/animals/{id}", controllers.DeleteAnimal).Methods("Delete")
    r.HandleFunc("/animals/{id}", controllers.UpdateAnimal).Methods("Put")

    //rota para a crud de actions
    r.HandleFunc("/actions", controllers.DisplayAllActions).Methods("Get")
    r.HandleFunc("/actions/{id}", controllers.GetActions).Methods("Get")
    r.HandleFunc("/actions/flock/{flock}", controllers.DisplayFlockActions).Methods("GET")
   // r.HandleFunc("/actions/flock/{flock}", controllers.GetActionsFlock).Methods("Get")
    //preciso criar um get por raça escolhendo bovino e suino
    r.HandleFunc("/actions", controllers.CreateActions).Methods("Post")
    r.HandleFunc("/actions/{id}", controllers.DeleteActions).Methods("Delete")
    r.HandleFunc("/actions/{id}", controllers.UpdateActions).Methods("Put")

    //relatorio de dados dos animais
    r.HandleFunc("/reports", controllers.DisplayAllReports).Methods("Get")
    r.HandleFunc("/reports/{id}", controllers.GetReports).Methods("Get")
    r.HandleFunc("/reports", controllers.CreateReports).Methods("Post")
    r.HandleFunc("/reports/{id}", controllers.DeleteReports).Methods("Delete")
    r.HandleFunc("/reports/{id}", controllers.UpdateReports).Methods("Put")

    //
    log.Fatal(http.ListenAndServe(":8080", r))
}