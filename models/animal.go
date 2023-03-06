package models

import(
	"desafio/database"
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
)

// info do animais
type Animal struct {
	Id		 		int 		`json:"id"`
	Sex				string		`json:"sex"`
    Datbirth 		int 		`json:"datbirth"`
	Birthweight		int			`json:"birthweight"`
	Flock  			int			`json:"flock"`
	Racename		string 		`json:"racename"`
	Weightnow		int 		`json:"weightnow"`
	Datesale 		int    		`json:"datesale"`
	Saleweight		int 		`json:"saleweight"`
	Situation		string		`json:"situation"`
}

var Animals []Animal;

//ação realizada por animal
type Action struct {
    Id                int    `json:"id"`
    Identifier        string `json:"identifier"`
    Actiondate        int    `json:"actiondate"`
    Flock             string    `json:"flock"`
    Descriptions      string `json:"descriptions"`
    Affectedanimals   string `json:"affectedanimals"`
    Cost              int    `json:"cost"`
    Vermicide         int    `json:"vermicide"`
    Supplementation   int    `json:"supplementation"`
}

var Actions []Action;

//açoes sendo tomada para rebanho com filtro de anos



//registros importantes por animal
type Report struct {
	Id				 	 	int 		`json:"id"`
	Datephysicalproblem		int			`json:"datephysicalproblem"`
	Descriptions 			string		`json:"descriptions"`
	Imageproblem	 	 	string		`json:"imageproblem"`
}

var Reports []Report;

func (a *Action) FindAllByFlockID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    var actions []Action
    database.DB.Where("identifier = ? AND flock_id = ?", a.Identifier, id).Find(&actions)
    json.NewEncoder(w).Encode(actions)
}