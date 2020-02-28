package controllers

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
	"github.com/kataras/iris/v12"
	"github.com/prinick96/ocrend_backend/models/users"
	"github.com/prinick96/ocrend_backend/models/usuarios"
)

/**
	* Manejador de peticiones GET
*/
func HandlerGet(app *iris.Application) {

	// Endpoint principal
	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Success GET"})
	})

	// Consultar información de un usuario requiere Header:Authorization
	app.Get("/user", users.GetUser)
	

	// Endpoint listado de usuarios
	app.Get("/usuarios", usuarios.GetUsuarios)

	// Obtener un usuario
	app.Get("/usuarios/{id}", usuarios.GetUsuarios)

}