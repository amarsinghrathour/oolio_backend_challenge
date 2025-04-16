// cmd/main.go
package main

import (
	"context"
	migratecmd "github.com/amarsinghrathour/oolio_backend_challenge/cmd/migrate"
	"github.com/amarsinghrathour/oolio_backend_challenge/config"
	"github.com/amarsinghrathour/oolio_backend_challenge/database"
	_ "github.com/amarsinghrathour/oolio_backend_challenge/docs"
	"github.com/amarsinghrathour/oolio_backend_challenge/middlewares"
	"github.com/amarsinghrathour/oolio_backend_challenge/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Order Food API
// @version 1.0
// @description This is a sample food ordering API.
// @contact.name Amar Singh Rathour
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name api_key
func main() {

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			migratecmd.RunMigration()
			return
		default:
			log.Fatalf("❌ Unknown command: %s", os.Args[1])
		}
	}

	cfg := config.LoadConfig("./.env")

	dbInstance, err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatal("DB init failed:", err)
	}
	database.SeedProducts(dbInstance.DB)
	log.Println("🔥 Starting Gin API server...")
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	// Swagger docs — skip API auth for swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Middlewares
	r.Use(middlewares.LoggerMiddleware())
	r.Use(func(c *gin.Context) {
		if c.FullPath() != "/swagger/*any" {
			middlewares.APIKeyAuthMiddleware(cfg.ApiKey)(c)
		}
	})

	// Setup Routes
	routes.SetupRoutes(r, dbInstance, cfg)

	// Server
	srv := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: r,
	}

	// Start server in goroutine
	go func() {
		log.Printf("🚀 Server started at :%s", cfg.AppPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Listen error: %s\n", err)
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🛑 Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Server forced to shutdown: %v", err)
	}

	log.Println("✅ Server exited properly")
}
