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
	"github.com/kataras/iris/v12"
	"database/sql"
	"github.com/prinick96/ocrend_backend/ocrend/database"
	"github.com/prinick96/ocrend_backend/ocrend/helpers"
)



/**
	* Obtiene la información de un usuario por su TOKEN DE SESIÓN
*/
func Delete(ctx iris.Context) {
	// Obtenemos la información
	id_user := ctx.FormValue("id_user")

	// Respuesta por defecto
	res := new(Response)
	res.Success = false


	// Verificar que los datos no estén vacíos
	if helpers.Empty(id_user)  {
		res.Message = "El id no puede estar vacío."
		ctx.JSON(res)
		return
	}

	// Conectamos a la base de datos
	db := database.Init()
	defer db.Close()

	// Consulta de sekect para verificar que lu suario exista
	row := db.QueryRow("SELECT id_user FROM users WHERE id_user = ? LIMIT 1 ", id_user)

	// Verificar que la consulta se ejecuto correctamnete
	var id_db uint64
	switch err := row.Scan(&id_db); err {
		case sql.ErrNoRows:
			
			// No existe
			res.Message = "Este usario no existe."
			res.Success = false
			ctx.JSON(res)
		case nil: // el usuario existe asi que lo eliminamos
			
			// Eliminar usuario
			stmt, err := db.Prepare("DELETE FROM users WHERE id_user = ? LIMIT 1")
			if err != nil {
				log.Print(err)
				panic(err)
			}

			query, err_exec := stmt.Exec(id_user)
			if err_exec != nil {
				log.Print(err_exec)
				panic(err_exec)
			}

			log.Print(query)

			res.Message = "Usuario eliminado."
			res.Success = true
			ctx.JSON(res)
		default:
			log.Print(err)
	}
}