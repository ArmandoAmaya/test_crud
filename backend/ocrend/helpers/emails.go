package helpers

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
	"bytes"
	"log"
	"html"
	"html/template"
	"net/smtp"
	"github.com/prinick96/ocrend_backend/ocrend/config"
)

var auth smtp.Auth

/**
	* Estructura del REQUEST
*/
type EmailRequest struct {
	From    string
	To      []string
	Subject string
	Body    string
}

/**
	* Construye un Request
*/
func newRequest(to []string, subject string, body string) *EmailRequest {
	return &EmailRequest{
		To:      to,
		Subject: subject,
		Body:    body,
	}
}

/**
	* Envía el correo electrónico
*/
func (r *EmailRequest) sendEmail() (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.Subject + "!\n"
	msg := []byte(subject + mime + "\n" + html.UnescapeString(r.Body))

	if err := smtp.SendMail(config.Data["mail_host"].(string) + ":" + config.Data["mail_port"].(string), auth, config.Data["mail_from"].(string), r.To, msg); err != nil {
		return false, err
	}
	return true, nil
}

/**
	* Construye la plantilla HTML
*/
func (r *EmailRequest) parseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.Body = buf.String()
	return nil
}

/**
	* Envía un correo electrónico con formato HTML
*/
func SendEmailMessage(to []string, subject string, body string) bool {
	auth = smtp.PlainAuth("", config.Data["mail_user"].(string), config.Data["mail_password"].(string), config.Data["mail_host"].(string))
	
	templateData := struct {
		UrlLogo string
		Logo  string
		Content string
		Copyright string 
	}{
		UrlLogo: config.Data["site_frontend"].(string),
		Logo:  "https://ocrend.com/dist/logo-footer-min.png",
		Content : body,
		Copyright : config.Data["site_name"].(string) + " &copy;. Todos los derechos reservados.",
	}
	
	r := newRequest(to, subject, body)
	err := r.parseTemplate("content/email_template.html", templateData)
	if err == nil {
		ok, error_enviando := r.sendEmail()
		if !ok {
			log.Print(error_enviando)
		} else {
			log.Print("Envio de recuperación de contraseña")
			log.Print(to)
		}

		return true
	} else {
		log.Print(err)
	}

	return false
}