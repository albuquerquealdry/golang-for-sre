package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
)

func main() {
	from := "cassielalarmistic@gmail.com"
	password := os.Getenv("CASSIEL_PASSWORD")

	if password == "" {
		panic("GMAIL PASSWORD NAO CONFIGURADO")
		os.Exit(1)
	}
	to := []string{
		"albuquerquealdry@gmail.com",
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)
	t, _ := template.ParseFiles("template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Alerta: Servidor down \n%s\n\n", mimeHeaders)))
	t.Execute(&body, struct {
		Server  string
		Error   string
		Horario string
	}{
		Server:  "Google",
		Error:   "Erro ao acessar o servidor.",
		Horario: "10/12/2022 14:00",
	})
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Printf("Erro ao enviar o email: %s", err)
		return
	}
	fmt.Println("Email enviado com sucesso!")

}
