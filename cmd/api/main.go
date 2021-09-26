package main

import (
	"log"

	"discounts-applier/cmd/api/dependencies"
)

func main() {
	dep := &dependencies.RealDependencies{}
	r, err := setupRouter(dep)
	if err != nil {
		log.Fatal(err)
	}
	err = r.Run(":8080")
	log.Fatal(err)
}
