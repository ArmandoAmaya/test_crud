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
	* Hace el cambio de la contraseña y resetea los campos de recuperación
*/
func RestorePassword(ctx iris.Context) {
	// Obtenemos la información
	lostpass_token := ctx.FormValue("lostpass_token")

	// Respuesta por defecto
	res := new(Response)
	res.Success = false

	// Verificar que los datos no estén vacíos
	if !helpers.Empty(lostpass_token) {
		
		// Inicializamos la base de datos
		db := database.Init()
		defer db.Close()

		// Buscamos la coincidencia en la DB
		row := db.QueryRow("SELECT id_user, name FROM users WHERE lostpass_url_token = ? LIMIT 1", lostpass_token)
		var id_user, name string
		switch err := row.Scan(&id_user, &name); err {
			case sql.ErrNoRows: // Sin resultados en la DB
				res.Message = "El enlace es inválido o ha caducado."
				ctx.JSON(res)
			case nil: // Obtuvimos resultados
				// Actualizamos la base de datos
				_, err_update := db.Exec("UPDATE users SET password = lostpass_temporal_password, lostpass_temporal_password = NULL , lostpass_url_token = NULL WHERE id_user = ? LIMIT 1", id_user)
				if err_update != nil {
					log.Print(err_update)
				}

				res.Success = true
				res.Message = "Hola " + name + ", hemos realizado el cambio de la contraseña correctamente."
				ctx.JSON(res)
			default:
				log.Print(err)
		}
	}
}