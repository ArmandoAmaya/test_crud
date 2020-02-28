package usuarios

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
	"time"
	"github.com/prinick96/ocrend_backend/ocrend/helpers"
)

/**
	Obtiene la información pública de un usuario
*/
type PublicUser struct {
	Id uint64 `json:"id_user"`
	Name string `json:"name"`
	Email string `json:"email"`
	Registerdate int64 `json:"register_date"`
	RegisterStrDate string `json:"str_register_date"`
	Success bool `json:"success"`
}


/**
	* Obtiene la información de un usuario por su TOKEN DE SESIÓN
*/
func GetUsuarios(ctx iris.Context) {
	// Conectamos a la base de datos
	db := database.Init()
	defer db.Close()
	id_user := ctx.Params().Get("id")

	
	if helpers.Empty(id_user) {
		data := []PublicUser{}

		row, error := db.Query("SELECT id_user, name, email, register_date FROM users")
	
		if nil == error {
			for row.Next() {
				var r PublicUser
				err := row.Scan(&r.Id, &r.Name, &r.Email, &r.Registerdate);

				if nil == err {
					tm := time.Unix(r.Registerdate, 0).UTC()
					r.RegisterStrDate = tm.Format("02/01/2006")
					data = append(data, r)
				}
			}
		}

		// Damos los resultados posibles
		switch error {
		case sql.ErrNoRows: // Sin resultados en la DB
			ctx.JSON("[]")
		case nil: // Obtuvimos resultados
			ctx.JSON(data)
		default:
			log.Print(error)
		}
	}else{
		data := new(PublicUser)
		row := db.QueryRow("SELECT id_user, name, email FROM users WHERE id_user = ? LIMIT 1", id_user)
		error := row.Scan(&data.Id, &data.Name, &data.Email)

		// Damos los resultados posibles
		switch error {
		case sql.ErrNoRows: // Sin resultados en la DB
			ctx.JSON(data)
		case nil: // Obtuvimos resultados
			data.Success = true
			ctx.JSON(data)
		default:
			log.Print(error)
		}
	}
	
	
}