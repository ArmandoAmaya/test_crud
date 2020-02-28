package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/prinick96/ocrend_backend/ocrend/config"
)

/**
	* Inicia la conexión con la base de datos
	* Devuelve la base de datos instanciada
*/
func Init() *sql.DB {
	db, err := sql.Open("mysql", config.Data["db_user"].(string) + ":" + config.Data["db_password"].(string) + "@tcp(" + config.Data["db_host"].(string) + ")/" + config.Data["db_name"].(string))
	
	if err != nil {
        panic(err.Error())
	}
	
	return db
}