package main

import "alert/email"

func main() {
	email.SendEmail([]string{"albuquerquealdry@gmail.com"},
		"Alerta de servidor down", "Google", "Erro ao conectar no servidor",
		"06/04/2022 15:04", "./email/template.html")
}
