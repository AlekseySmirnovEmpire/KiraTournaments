package handler

import (
	"github.com/AlekseySmirnovEmpire/CardMarketplace/package/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{services: s}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := r.Group("/api", h.userIdentity)
	{
		users := api.Group("/users")
		{
			users.GET("/", h.getUsersList)
			users.GET("/:id", h.getUser)
			users.POST("/", h.createUser)
			users.PUT("/:id", h.updateUser)
			users.DELETE("/:id", h.deleteUser)
		}
	}

	return r
}
