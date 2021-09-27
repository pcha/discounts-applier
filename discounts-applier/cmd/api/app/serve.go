package app

import (
	"log"

	"discounts-applier/cmd/api/app/dependencies"
)

func Serve() error {
	dep := &dependencies.RealDependencies{}
	r, err := setupRouter(dep)
	if err != nil {
		log.Fatal(err)
	}
	return r.Run(":8080")
}
