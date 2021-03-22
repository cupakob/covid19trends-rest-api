package router_test

import (
	"github.com/cupakob/covid19trends-rest-api/handler"
	"github.com/cupakob/covid19trends-rest-api/router"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRouter(t *testing.T) {
	t.Run("should create new router", func(t *testing.T) {
	    // given
		var h handler.Handler

		// when
		newRouter := router.NewRouter(h)

		// then
		assert.NotNil(t, newRouter)
	})
}
