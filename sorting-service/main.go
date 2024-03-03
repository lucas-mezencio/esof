package main

import (
	"log"
	"sorting-service/routes"
)

func main() {
	r := routes.SetupRoutes()
	log.Fatal(r.Run(":8080"))
}
