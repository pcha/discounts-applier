// Package app contains the needed to start the application a d test it
package app

import (
	"discounts-applier/cmd/api/app/router"
	"discounts-applier/cmd/api/app/router/dependencies"

	"github.com/gin-gonic/gin"
)

func Serve() error {
	dep := &dependencies.RealDependencies{}
	r, err := setupRouter(dep)
	if err != nil {
		return err
	}
	return run(r)
}

type SetupRouterFunc func(dependencies.Dependencies) (*gin.Engine, error)

var setupRouter = router.SetupRouter

type RunFunc func(r *gin.Engine) error

var run RunFunc = func(r *gin.Engine) error {
	return r.Run(":8080")
}
