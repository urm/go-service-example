package webapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/urm/go-service-example/internal/ports"
)

type Server struct {
	Tasks  ports.Tasks
	Router http.Handler
}

func New(tasks ports.Tasks) *Server {
	s := new(Server)
	s.Tasks = tasks

	router := gin.Default()

	router.POST("/login", s.auth)
	tasksGroup := router.Group("/tasks")
	{
		tasksGroup.POST("/run", s.createTask)
		tasksGroup.DELETE("/:id", s.deleteTask)
		tasksGroup.GET("/approve/:id/:user", s.approveTask)
		tasksGroup.GET("/decline/:id/:user", s.declineTask)
	}

	s.Router = router
	return s
}
