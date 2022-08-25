package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vicfntm/splitshit2/src/services"
)

type Handler struct {
	services *services.Services
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/assign-role", h.UserIdentity, h.userSuperAdminMiddleware, h.assignRole)
		auth.POST("sign-in", h.signIn)
	}

	api := router.Group("/public")
	{
		api.Use(h.UserIdentity)
		api.POST("/create-group", h.createGroup)
		api.GET("/available-groups", h.availableGroups)
		api.POST("/join-the-group", h.joinGroup)
		api.POST("/grab-the-shit", h.grabShit)
	}

	return router
}
func NewHandler(s *services.Services) *Handler {
	return &Handler{
		services: s,
	}
}
