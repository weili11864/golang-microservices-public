package main

import (
	"log"
	"net/http"

	"github.com/KernelGamut32/golang-microservices/demos/toyshop/internal/routes"
)

func main() {
	r := routes.Handlers()

	err := http.ListenAndServe(":5000", r)
	if err != nil {
		log.Fatal(err)
	}
}
