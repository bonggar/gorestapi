package router

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestInitializeRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	routeHandler := Make()
	if routeHandler != nil {
		t.Logf("Route Initialization Success")
	} else {
		t.Fail()
	}
}
