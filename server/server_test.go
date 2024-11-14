package server_test

import (
	"canvas/integrationtest"
	"github.com/matryer/is"
	"net/http"
	"testing"
)

func TestServer_Start(t *testing.T) {
	integrationtest.SkipIfShort(t)

	t.Run("starts the server and listens for requests", func(t *testing.T) {
		assert := is.New(t)

		cleanup := integrationtest.CreateServer()
		defer cleanup()

		resp, err := http.Get("http://localhost:8081/")
		assert.NoErr(err)
		assert.Equal(http.StatusNotFound, resp.StatusCode)
	})
}
