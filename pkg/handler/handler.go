package handler

import "github.com/gin-gonic/gin"

type Handler struct {
	
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up")
		auth.POST("/sign-in")
	}
	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.GET("/")
			lists.GET("/:id")
			lists.POST("/")
			lists.PUT("/:id")
			lists.DELETE("/:id")

			items := lists.Group(":id/items")
			{
				items.GET("/")
				items.GET("/:item_id")
				items.POST("/")
				items.PUT("/:item_id")
				items.DELETE("/:item_id")
			}
		}
	}
	return router
}
