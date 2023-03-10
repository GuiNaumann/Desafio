package database

import (
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)
//gorm sendo utilizado para fazer fazer comunicação da database
var (
    DB  *gorm.DB
    err error
)

func ConectaBanco() {
    stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
    DB, err = gorm.Open(postgres.Open(stringDeConexao))
    if err != nil {
        log.Panic("Erro ao conectar com banco de dados")
    }
}