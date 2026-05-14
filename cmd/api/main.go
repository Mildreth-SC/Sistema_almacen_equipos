package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Iniciando servidor...")
	http.ListenAndServe(":8080", nil)
}
