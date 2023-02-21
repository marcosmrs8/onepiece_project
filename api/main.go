package main

import (
	"api/src/routes"
	"fmt"
)

func main() {
	fmt.Println("Iniciando servidor...")
	routes.HandleRequest()
}
