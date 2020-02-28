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
	"time"
	"log"
	"database/sql"
	"github.com/kataras/iris/v12"
	"github.com/prinick96/ocrend_backend/ocrend/database"
	"github.com/prinick96/ocrend_backend/ocrend/helpers"
)

/**
	* Registra un nuevo usuario
*/
func Register(ctx iris.Context) {
	// Obtenemos la información
	name := ctx.FormValue("name")
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")
	password_repeat := ctx.FormValue("password_repeat")
	
	// Respuesta por defecto
	res := new(ResponseWithUser)
	res.Token = ""
	res.Success = false

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

	// Verificamos que no exista un usuario con el mismo correo en la DB
	var email_db string
	row := db.QueryRow("SELECT email  FROM users WHERE email = ? LIMIT 1", email)
	switch err := row.Scan(&email_db); err {
		case sql.ErrNoRows: // Sin resultados en la DB, puede continuar
			// Encriptamos la contraseña
			password = helpers.Encrypt(password)

			// Definimos la fecha actual de registro
			res.UserInformation.Registerdate = uint32(time.Now().Unix())

			// Creamos el token
			res.Token = helpers.TokenGenerator()

			// Insertamos al usuario en la base de datos
			stmt, err := db.Prepare("INSERT INTO users (name, email, password, token, register_date) VALUES (?, ?, ?, ?, ?)")
			if err != nil {
				log.Print(err)
				panic(err)
			}
			query, err_exec := stmt.Exec(name, email, password, res.Token, res.UserInformation.Registerdate)
			if err_exec != nil {
				log.Print(err_exec)
				panic(err_exec)
			}

			// Obtenemos el último id
			res.UserInformation.Id, _ = query.LastInsertId()

			// Llenamos la información del usuario
			res.UserInformation.Name = name
			res.UserInformation.Email = email

			// Entregamos exito
			res.Message = "Registrado con éxito."
			res.Success = true
			ctx.JSON(res)
		case nil: // Obtuvimos resultados
			res.Message = "El email introducido ya existe."
			ctx.JSON(res)
		default:
			log.Print(err)
	}	
}