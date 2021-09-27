package app

import (
	"errors"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestServe(t *testing.T) {

	tests := []struct {
		name     string
		errSetup error
		errRun   error
	}{
		{
			name: "no errors",
		},
		{
			name:     "error on setup",
			errSetup: errors.New("setupErr"),
		},
		{
			name:   "error on run",
			errRun: errors.New("runErr"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.New()
			stopSrStub := stubSetupRouter(r, tt.errSetup)
			stopRStub := stubRun(tt.errRun)
			defer stopSrStub()
			defer stopRStub()

			err := Serve()
			if tt.errSetup != nil {
				assert.Same(t, tt.errSetup, err)
			} else if tt.errRun != nil {
				assert.Same(t, tt.errRun, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
