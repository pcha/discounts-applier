package app

import (
	"discounts-applier/cmd/api/app/router/dependencies"

	"github.com/gin-gonic/gin"
)

type StopStub func()

func stubSetupRouter(r *gin.Engine, err error) StopStub {
	origVal := setupRouter
	setupRouter = func(_ dependencies.Dependencies) (*gin.Engine, error) {
		return r, err
	}
	return func() {
		setupRouter = origVal
	}
}

func stubRun(err error) StopStub {
	origVal := run
	run = func(r *gin.Engine) error {
		return err
	}
	return func() {
		run = origVal
	}
}
