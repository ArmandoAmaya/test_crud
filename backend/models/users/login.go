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
	* Inicia sesión con las credenciales email y password
*/
func Login(ctx iris.Context) {
	// Obtenemos la información
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")
	
	// Respuesta por defecto
	res := new(ResponseWithUser)
	res.Token = ""
	res.Success = false

	// Verificar que los datos no estén vacíos
	if helpers.Empty(email) || helpers.Empty(password) {
		res.Message = "Todos los campos deben estar llenos."

		ctx.JSON(res)
		return
	}

	// Encriptamos la contraseña
	password = helpers.Encrypt(password)

	// Conectamos a la base de datos
	db := database.Init()
	defer db.Close()

	// Buscamos la coincidencia en la DB
	row := db.QueryRow("SELECT id_user, name, email, register_date FROM users WHERE email = ? AND password = ? LIMIT 1", email, password)

	switch err := row.Scan(&res.UserInformation.Id, &res.UserInformation.Name, &res.UserInformation.Email, &res.UserInformation.Registerdate); err {
		case sql.ErrNoRows: // Sin resultados en la DB
			res.Message = "Credenciales incorrectas."
			ctx.JSON(res)
		case nil: // Obtuvimos resultados
			// Generamos el token y marcamos como logeado
			res.Success, res.Token, res.Message = true , helpers.TokenGenerator(), "Conectado correctamente."
			
			// Actualizamos el token en la base de datos
			_, err_update := db.Exec("UPDATE users SET token = ? WHERE id_user = ? LIMIT 1", res.Token, res.UserInformation.Id)
			if err_update != nil {
				log.Print(err_update)
			}
			
			ctx.JSON(res)
		default:
			log.Print(err)
	}	

}