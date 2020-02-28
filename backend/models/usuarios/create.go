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
	"time"
	"log"
	"database/sql"
	"github.com/kataras/iris/v12"
	"github.com/prinick96/ocrend_backend/ocrend/database"
	"github.com/prinick96/ocrend_backend/ocrend/helpers"
)


/**
	* crea un nuevo usuario
*/
func Create(ctx iris.Context) {

	// Obtenemos la información
	name := ctx.FormValue("name")
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")
	password_repeat := ctx.FormValue("password_repeat")
	
	res := new(Response)
	res.Success = false
	res.Message = ""

	// Verificar que los datos no estén vacíos
	if helpers.Empty(name) || helpers.Empty(email) || helpers.Empty(password) || helpers.Empty(password_repeat)  {
		res.Message = "Todos los campos deben estar llenos."

		ctx.JSON(res)
		return
	}

	// Verificamos que las contraseñas sean iguales
	if password != password_repeat {
		res.Message = "Las contraseñas deben ser iguales."

		ctx.JSON(res)
		return
	}

	// Conectamos a la base de datos
	db := database.Init()
	defer db.Close()


	var email_db string
	row := db.QueryRow("SELECT email  FROM users WHERE email = ? LIMIT 1", email)
	switch err := row.Scan(&email_db); err {
		case sql.ErrNoRows:

			// Encriptamos la contraseña
			password = helpers.Encrypt(password)

			// Definimos la fecha actual de registro
			register_date := uint32(time.Now().Unix())
			// Creamos el token
			token := helpers.TokenGenerator()

			// Insertamos al usuario en la base de datos
			stmt, err := db.Prepare("INSERT INTO users (name, email, password, token, register_date) VALUES (?, ?, ?, ?, ?)")
			if err != nil {
				log.Print(err)
				panic(err)
			}

			query, err_exec := stmt.Exec(name, email, password, token, register_date)
			if err_exec != nil {
				log.Print(err_exec)
				panic(err_exec)
			}

			log.Print(query)
			
			// Entregamos exito
			res.Message = "Guardado con éxito."
			res.Success = true
			ctx.JSON(res)
		case nil:
			res.Message = "El email introducido ya existe."
			ctx.JSON(res)
		default:
			log.Print(err)
	}			
}