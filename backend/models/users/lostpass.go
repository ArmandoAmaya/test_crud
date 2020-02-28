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
	"github.com/prinick96/ocrend_backend/ocrend/config"
)

/**
	* Genera el enlace de recuperar contraseña y almacena una nueva cotraseña temporal
*/
func GenerateLostpass(ctx iris.Context) {
	// Obtenemos la información
	email := ctx.FormValue("email")

	// Respuesta por defecto
	res := new(Response)
	res.Success = false

	// Verificar que los datos no estén vacíos
	if helpers.Empty(email) {
		res.Message = "El email no puede estar vacío."

		ctx.JSON(res)
		return
	}

	// Inicializamos la base de datos
	db := database.Init()
	defer db.Close()

	// Buscamos la coincidencia en la DB
	row := db.QueryRow("SELECT id_user, name FROM users WHERE email = ? LIMIT 1", email)
	var id_user, name string
	switch err := row.Scan(&id_user, &name); err {
		case sql.ErrNoRows: // Sin resultados en la DB
			res.Message = "El email introducido no existe en el sistema."
			ctx.JSON(res)
		case nil: // Obtuvimos resultados
			// Generamos un token para la url
			token := helpers.TokenGenerator()

			// Generamos una nueva contraseña
			new_password := helpers.RandomPassword(15)
			
			// Actualizamos en la DB
			_, err_update := db.Exec("UPDATE users SET lostpass_temporal_password = ?, lostpass_url_token = ? WHERE id_user = ? LIMIT 1", helpers.Encrypt(new_password), token,id_user)
			if err_update != nil {
				log.Print(err_update)
			}

			// Generamos la URL
			url := config.Data["site_frontend"].(string) + "restore?token=" + token
			message := "Hola " + name + 
			", hemos recibido una solicitud de recuperación de contraseña. <br /><br /> Para proceder a cambiar tu contraseña por <b>" + new_password + 
			"</b>, debes hacer clic en el siguiente enlace: <br /><br /><a href=\""+ url +"\">" + url + 
			"</a> <br /><br /> Si no has solicitado este cambio, no debes hacer nada y aconsejamos cambiar tu contraseña actual."

			// Enviar correo
			go helpers.SendEmailMessage([]string{email}, "Recuperación de Contraseña - " + config.Data["site_name"].(string), message)
			res.Message = "Se ha enviado un enlace de recuperación y la nueva contraseña a su correo electrónico."
			res.Success = true
			
			ctx.JSON(res)
		default:
			log.Print(err)
	}	

}

