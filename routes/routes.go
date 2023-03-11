package routes

// todos os imports usado no arquivo routes
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
    r.HandleFunc("/", controllers.Home)

    //
    r.Use(middleware.ContentTypeMiddleware)

    // log em todas as requisições http ultilizando middleware
    r.Use(httplog.Logger)
    
    //rota para a crud de animals
    r.HandleFunc("/animals", controllers.DisplayAllAnimals).Methods("Get")
    r.HandleFunc("/animals/{id}", controllers.GetAnimals).Methods("Get")
    //filtro por tipo como bovino/suino
    r.HandleFunc("/animals/flock/{flock}", controllers.DisplayFlockAnimal).Methods("GET")
    r.HandleFunc("/animals", controllers.CreateAnimal).Methods("Post")
    r.HandleFunc("/animals/{id}", controllers.DeleteAnimal).Methods("Delete")
    r.HandleFunc("/animals/{id}", controllers.UpdateAnimal).Methods("Put")

    //rota para a crud de actions
    r.HandleFunc("/actions", controllers.DisplayAllActions).Methods("Get")
    r.HandleFunc("/actions/{id}", controllers.GetActions).Methods("Get")
    //filtro por tipo como bovino/suino
    r.HandleFunc("/actions/flock/{flock}", controllers.DisplayFlockActions).Methods("GET")
    r.HandleFunc("/actions", controllers.CreateActions).Methods("Post")
    r.HandleFunc("/actions/{id}", controllers.DeleteActions).Methods("Delete")
    r.HandleFunc("/actions/{id}", controllers.UpdateActions).Methods("Put")

    //relatorio de dados dos animais
    r.HandleFunc("/reports", controllers.DisplayAllReports).Methods("Get")
    r.HandleFunc("/reports/{id}", controllers.GetReports).Methods("Get")
    //filtro por tipo como numero do problema
    r.HandleFunc("/reports/code/{code}", controllers.DisplayFlockReports).Methods("GET")


    r.HandleFunc("/reports", controllers.CreateReports).Methods("Post")
    r.HandleFunc("/reports/save/{url}", controllers.Teste).Methods("Get")
    r.HandleFunc("/reports/test/{url}", controllers.Testeee).Methods("GET")


    r.HandleFunc("/reports/{id}", controllers.DeleteReports).Methods("Delete")
    r.HandleFunc("/reports/{id}", controllers.UpdateReports).Methods("Put")

    //
    r.HandleFunc("/reports/teste/{id}", controllers.Testimages).Methods("get")
    
    //http://localhost:8080/
    log.Fatal(http.ListenAndServe(":8080", r))

}