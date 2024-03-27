package rest

import (
	"Narcolepsick1d/mini-twitter/internal/middleware"
	"Narcolepsick1d/mini-twitter/internal/scope"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// swaggo/swag requires import of docs for side effect.
	_ "Narcolepsick1d/mini-twitter/docs"
)

type HandlerConfig struct {
	Dependencies *scope.Dependencies
}

// NewHandler initializes REST API Handlers.
func NewHandler(cfg HandlerConfig) http.Handler {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/sign-up", cfg.signUp)
	router.POST("/sign-in", cfg.signIn)
	router.Use(middleware.AuthMiddleware())
	router.POST("/follow", cfg.follow)
	router.POST("/unfollow", cfg.unfollow)
	router.POST("/tweet", cfg.tweet)
	return router
}
