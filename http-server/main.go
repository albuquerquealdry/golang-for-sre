package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	auth "github.com/abbot/go-http-auth"
)

func Secret(user, realm string) string {
	if user == "kidboo" {
		return "$1$uBIqPv6k$NuRznkcYuLp.v.qxtVhUO."
	}
	return ""
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("##Errou Leoa## \n Uso: go run main.go <diretorio> <porta>")
		os.Exit(1)
	}
	httpDir := os.Args[1]
	port := os.Args[2]

	authenticator := auth.NewBasicAuthenticator("meuserver.com", Secret)
	http.HandleFunc("/", authenticator.Wrap(func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
		http.FileServer(http.Dir(httpDir)).ServeHTTP(w, &r.Request)
	}))

	fmt.Print("#######  KIDBOO DA FANAUTICO WEBSERVER FILE SYSTEM ####### \n" + "Share Directory:" + httpDir + "\nPort:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
