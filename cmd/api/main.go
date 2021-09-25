package main

import (
	"log"

	"discounts-applier/cmd/api/dependencies"
)

func main() {
	dep := &dependencies.RealDependencies{}
	r := setupRouter(dep)
	err := r.Run(":8080")
	log.Fatal(err)
}
