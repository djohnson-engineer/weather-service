package controllers

import (
	"github.com/gin-gonic/gin"
)

// Health godoc
//
//	@Summary		Service health
//	@Description	Check if the service is up
//	@Tags			health
//	@Security		Authorization
//	@Accept			json
//	@Produce		json
//	@Success		200	string		Ok
//	@Failure		400	{object}	Response
//	@Failure		404	{object}	Response
//	@Router			/v1/health [get]
func (*Server) GetHealth(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Ok",
	})
}
