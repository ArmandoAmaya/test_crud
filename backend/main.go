package main

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
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/prinick96/ocrend_backend/controllers"
	"github.com/prinick96/ocrend_backend/ocrend/config"
)

func main() {
	app := iris.New()

	// Si está en modo producción el debug estará deshabilitado
	if !config.Data["site_production"].(bool) {
		app.Logger().SetLevel("debug")
		app.Use(recover.New())
		app.Use(logger.New())
	}

	crs := cors.New(cors.Options{
        AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Access-Control-Allow-Origin","Origin","X-Requested-With","Content-Type","Content-Length","Accept","Authorization"},
        AllowCredentials: true,
    })
    app.Use(crs)
	app.AllowMethods(iris.MethodOptions)

	// Manejadores HTTP
	controllers.HandlerGet(app)
	controllers.HandlerPost(app)
	
	// Arranca la aplicación en el puerto definido
	app.Run(iris.Addr(config.Data["site_backend_port"].(string)), iris.WithoutServerError(iris.ErrServerClosed))
}