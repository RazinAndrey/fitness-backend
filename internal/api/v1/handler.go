package v1

import (
	"github.com/RazinAndrey/fitness-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{}
}

// иницилизирует все наши end-points
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New() // иницилизация роутера

	// обявление методов и группировка их по маршрутам

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllList)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
		}
	}

	return router
}
