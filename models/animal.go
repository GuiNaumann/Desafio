package models

// info do animais
type Animal struct {
	Id		 		int 		`json:"id"`
	Sex				string		`json:"sex"`
    Datbirth 		string 		`json:"datbirth"`
	Birthweight		int			`json:"birthweight"`
	Flock  			string		`json:"flock"`
	Racename		string 		`json:"racename"`
	Weightnow		int 		`json:"weightnow"`
	Datesale 		string    	`json:"datesale"`
	Saleweight		int 		`json:"saleweight"`
	Situation		string		`json:"situation"`
//	Reports			[]Report	 `json:"report"`
}

var Animals []Animal;

//ação realizada por animal
type Action struct {
    Id                int    	`json:"id"`
    Actiondate        string   	`json:"actiondate"`
    Flock             string    `json:"flock"`
    Descriptions      string	`json:"descriptions"`
    Affectedanimals   string	`json:"affectedanimals"`
    Cost              int    	`json:"cost"`
    Vermicide         int    	`json:"vermicide"`
    Supplementation   int    	`json:"supplementation"`
}

var Actions []Action;

//açoes sendo tomada para rebanho com filtro de anos
type Report struct {
	Id				 	 	int 		`json:"id"`
	Problemcode				int 		`json:"code"`
	Datephysicalproblem		string		`json:"datephysicalproblem"`
	Descriptions 			string		`json:"descriptions"`
	Imageproblem	 	 	string		`json:"url"`
}

var Reports []Report;