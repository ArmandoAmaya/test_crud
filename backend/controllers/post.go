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
	* Manejador de peticiones POST
*/
func HandlerPost(app *iris.Application) {

	// Endpoint principal
	app.Post("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Success POST"})
	})	

	// Login de usuario
	app.Post("/user/login", users.Login)

	// Registro de usuario
	app.Post("/user/register", users.Register)

	// Solicitud de recuperación de contraseña
	app.Post("/user/lostpass", users.GenerateLostpass)

	// Cambio de la contraseña una vez tenga el enlace con el token
	app.Post("/user/restore", users.RestorePassword)

	// crear un usuario
	app.Post("/usuarios/create", usuarios.Create)

	// Elkiminar un usuario
	app.Post("/usuarios/delete", usuarios.Delete)

	// edita un usuario
	app.Post("/usuarios/edit/{id}", usuarios.Edit)
}