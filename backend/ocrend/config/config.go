package config

/*
 * This file is part of the Ocrend Framewok package.
 *
 * (c) Ocrend Software <info@ocrend.com>
 * @author Brayan Narváez <prinick@ocrend.com>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
*/

var Data = map[string]interface{}{
	"db_host" : "localhost",
	"db_user" : "root",
	"db_password" : "",
	"db_name" : "proyecto_go",

	"mail_host" : "smtp.gmail.com",
	"mail_user" : "ocrendsoftware@gmail.com",
	"mail_password" : "ocrendsoftware++demo",
	"mail_port" : "587",
	"mail_from" : "noreply@ocrend.com",

	"site_frontend" : "http://localhost:8080/#/",
	"site_backend_port" : ":9000",
	"site_production" : false,
	"site_name" : "Proyecto GO",
	"site_timezone" : "America/Bogota",
}