package webapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) createTask(c *gin.Context) {
	c.String(http.StatusOK, "Hello there")
}

func (s *Server) deleteTask(c *gin.Context) {
	c.String(http.StatusOK, "Hello there")
}

func (s *Server) approveTask(c *gin.Context) {
	c.String(http.StatusOK, "Hello there")
}

func (s *Server) declineTask(c *gin.Context) {
	c.String(http.StatusOK, "Hello there")
}
