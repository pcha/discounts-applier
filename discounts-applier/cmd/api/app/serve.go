package app

import (
	"log"

	"discounts-applier/cmd/api/app/dependencies"
)

func Serve() {
	dep := &dependencies.RealDependencies{}
	r, err := setupRouter(dep)
	if err != nil {
		log.Fatal(err)
	}
	err = r.Run(":8080")
	log.Fatal(err)
}
