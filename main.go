package main

import (
	"fmt"
	"flag"
	"github.com/ilyanSD/password_manager/client"
	"github.com/ilyanSD/password_manager/server"
)

func main() {
	var option string

	flag.StringVar(&option, "type", "s", "Defines if the program will be executed as a server (s) or a client (c).")

	flag.Parse()

	if option == "s" {
		server.Server()
	} else if option == "c" {
		client.Client()
	} else {
		fmt.Println("Opcion no valida! Prueba a utilizar '-type s' para ejecutar el servidor o '-type c' para ejecutar el cliente.")
	}
}