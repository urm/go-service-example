package webapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) auth(c *gin.Context) {
	c.String(http.StatusOK, "Hello there")
}
