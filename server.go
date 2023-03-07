package main

import (
	"english_exam_go/application/middleware"
	"english_exam_go/application/routers"
	"english_exam_go/infrastructure/data/repositories"
	"english_exam_go/utils/app_logger"
	"english_exam_go/utils/di"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	mode := os.Getenv("MODE")

	apiPort := os.Getenv("API_PORT")
	dotEnvInit(mode)
	app_logger.Init()
	app_logger.Logger.Debug("Logger init succeeded")

	r := ginInit(mode)
	app_logger.Logger.Debug("Gin init succeeded")

	c := di.RegisterDIFunction()
	app_logger.Logger.Debug("Register all DI function succeeded")

	repositories.OpenDatabase()
	app_logger.Logger.Debug("Connect DB succeeded")

	if err := c.Invoke(func(s *routers.Service) {
		baseRoute := r.Group("/api")
		r.NoRoute(func(c *gin.Context) {
			c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
		})

		s.RouterInit(baseRoute)
		app_logger.Logger.Debug("All routers init succeeded")
	}); err != nil {
		app_logger.Logger.Panic("Failed to resolve dependency: " + err.Error())
	}

	// StartUp Server
	app_logger.Logger.Info("Listening and serving " + "HTTP on :" + apiPort + " , " + "MODE:" + mode)
	err = r.Run(":" + os.Getenv("API_PORT"))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func dotEnvInit(mode string) {
	if mode == "PROD" {
		if err := godotenv.Load(".prod_env"); err != nil {
			panic("Can not loading .prod_env. error:" + err.Error())
		}
	} else {
		if err := godotenv.Load(".env"); err != nil {
			panic("Can not loading .env. error:" + err.Error())
		}
	}
}

func ginInit(mode string) *gin.Engine {
	if mode == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// gin-zap middleware
	r.Use(ginzap.Ginzap(app_logger.Logger, time.RFC3339, true))
	// logging all panic to error log
	r.Use(ginzap.RecoveryWithZap(app_logger.Logger, true))
	// CORS middleware
	r.Use(middleware.SetCors())

	return r
}
