package users

/*
 * This file is part of the Ocrend Framewok package.
 *
 * (c) Ocrend Software <info@ocrend.com>
 * @author Brayan Narváez <prinick@ocrend.com>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
*/

import (
	"log"
	"database/sql"
	"github.com/kataras/iris/v12"
	"github.com/prinick96/ocrend_backend/ocrend/database"
	"github.com/prinick96/ocrend_backend/ocrend/helpers"
)

/**
	Obtiene la información pública de un usuario
*/
type PublicUser struct {
	Id uint64 `json:"id_user"`
	Name string `json:"name"`
	Email string `json:"email"`
	Registerdate int `json:"register_date"`
}

/**
	* Obtiene la información de un usuario por su TOKEN DE SESIÓN
*/
func GetUser(ctx iris.Context) {
	// Conectamos a la base de datos
	db := database.Init()
	defer db.Close()

	// Obtenemos el token del usuario solicitado
	token := ctx.GetHeader("Authorization")

	// Si es inválido retornamos un JSON vacío
	if helpers.Empty(token) {
		ctx.JSON("[]")
		return
	}
	
	// Hacemos la consulta y llenamos la estructura
	u := new(PublicUser)
	row := db.QueryRow("SELECT id_user, name, email, register_date FROM users WHERE token = ? LIMIT 1", token)

	// Damos los resultados posibles
	switch err := row.Scan(&u.Id, &u.Name, &u.Email, &u.Registerdate); err {
	case sql.ErrNoRows: // Sin resultados en la DB
		ctx.JSON("[]")
	case nil: // Obtuvimos resultados
		ctx.JSON(u)
	default:
		log.Print(err)
	}	
}