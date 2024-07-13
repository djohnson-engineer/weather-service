package controllers

import (
	"net/http/httptest"
	"testing"
	"weather-server/src/testutils"

	"github.com/stretchr/testify/assert"
)

func Test_GetHealth(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := testutils.GetTestContext(w)

	testServer, _ := NewTestServer(t)
	testServer.GetHealth(ctx)

	assert.Equal(t, 200, w.Result().StatusCode)
}
