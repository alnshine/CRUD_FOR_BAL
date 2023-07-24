package handler

import (
	_ "alnshine/CRUD_FOR_BAL/docs"
	"alnshine/CRUD_FOR_BAL/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		vacancy := api.Group("/vacancy")
		{
			vacancy.POST("/", h.createVacancy)
			vacancy.GET("/", h.getAllVacancy)
			vacancy.GET("/:id", h.getVacancyById)
			vacancy.PUT("/:id", h.updateVacancy)
			vacancy.DELETE("/:id", h.deleteVacancy)
		}
	}
	return router
}
