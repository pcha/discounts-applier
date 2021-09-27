package app

import (
	"errors"
	"testing"
	"time"

	"discounts-applier/cmd/api/app/router"
	"discounts-applier/cmd/api/app/router/dependencies"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_stubSetupRouter(t *testing.T) {
	mr := gin.New()
	merr := errors.New("")
	stop := stubSetupRouter(mr, merr)
	dep := &dependencies.RealDependencies{}
	r, err := setupRouter(dep)
	assert.Same(t, mr, r)
	assert.Same(t, merr, err)
	stop()
	r, err = setupRouter(dep)
	r, err = router.SetupRouter(dep)
}

func Test_stubRun(t *testing.T) {
	r := gin.New()
	merr := errors.New("")

	stop := stubRun(merr)
	err := run(r)

	assert.Same(t, merr, err)
	stop()
	err = nil
	go func(e error) {
		e = r.Run()
	}(err)
	time.Sleep(500 * time.Millisecond)
	assert.Nil(t, err)
}
