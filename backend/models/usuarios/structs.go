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


/**
	* Respuesta por defecto básica
*/
type Response struct {
	Message string `json:"message"`
	Success bool `json:"success"`
}