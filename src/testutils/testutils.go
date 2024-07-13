package testutils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTestContext(w http.ResponseWriter) *gin.Context {
	gin.SetMode("test")
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	return c
}
