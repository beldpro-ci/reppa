package router_test

import (
	repparouter "github.com/beldpro-ci/reppa/reppa/router"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInspectRouteWithoutContainers(t *testing.T) {
	router := repparouter.New()
	writer := httptest.NewRecorder()

	request, err := http.NewRequest("GET", "/ping", nil)
	assert.NoError(t, err)

	router.ServeHTTP(writer, request)
	assert.Equal(t, writer.Body.String(), "PONG")

}
