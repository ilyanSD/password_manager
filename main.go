package main

import (
	"fmt"
	"flag"
	"github.com/ilyanSD/password_manager/client"
	"github.com/ilyanSD/password_manager/server"
)

func main() {
	var option string

	flag.StringVar(&option, "type", "s", "Especifica si quieres ejecutar el servidor (s) o el cliente (c)") // Por defecto ejecutamos el servidor

	flag.Parse()

	if option == "s" {
		server.Server()
	} else if option == "c" {
		client.Client()
	} else {
		fmt.Println("Opcion no valida! Prueba a utilizar '-type s' para ejecutar el servidor o '-type c' para ejecutar el cliente.")
	}
}