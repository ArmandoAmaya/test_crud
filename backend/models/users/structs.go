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

/**
	* Estructura de un usuario
*/
type UserInformation struct {
	Id int64 `json:"id_user"`
	Name string `json:"name"`
	Email string `json:"email"`
	Registerdate uint32 `json:"register_date"`
}

/**
	* Respuesta completa con Usuario y Token
*/
type ResponseWithUser struct {
	UserInformation `json:"user"`
	Token string `json:"token"`
	Message string `json:"message"`
	Success bool `json:"success"`
}

/**
	* Respuesta por defecto básica
*/
type Response struct {
	Message string `json:"message"`
	Success bool `json:"success"`
}