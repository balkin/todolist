package main

import (
	"context"
	"github.com/balkin/todolist/controllers"
	"github.com/balkin/todolist/db"
	_ "github.com/balkin/todolist/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//Sigterm: graceful (finish all requests and quit)
//Sigint: fast quit
//Postgresql
//Unit tests
//docker-compose up -d
//Features:
// * Add TODO
// * See all todos (count only)
// * See specific todo (with children)
// * Add subitem
// * Delete specific

var HttpDaemon *http.Server

// @host localhost:8000
// @BasePath /api/v1
func main() {
	log.SetFlags(log.LstdFlags)
	router := SetupRouter()
	db.ConnectToDatabase()
	defer db.DisconnectDatabase()

	HttpDaemon = &http.Server{Addr: ":8000", Handler: router, ReadTimeout: 5 * time.Second, WriteTimeout: 5 * time.Second}
	go func() {
		if err := HttpDaemon.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to run server %v\n", err)
		}
	}()

	// Handle SIGINT and SIGTERM.
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	select {
	case sig := <-ch:
		switch sig {
		case syscall.SIGINT:
			log.Println("Immediately stopping HTTP server")
			return
		case syscall.SIGTERM:
			log.Println("Gracefully stopping HTTP server")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := HttpDaemon.Shutdown(ctx); err != nil {
				log.Println("Error during HTTP server shutdown:", err)
				return
			}
			<-ctx.Done()
		}
	}
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", controllers.IndexController)
	router.GET("/ping", controllers.PingController)
	v1 := router.Group("/api/v1")
	{
		// Add some middleware to process JWT tokens or similar stuff
		todoApi := v1.Group("/todo")
		{
			todoApi.GET("count", controllers.TodoCountItems)
			todoApi.GET("countall", controllers.TodoCountAllItems)
			todoApi.GET("item/", controllers.TodoListItems)
			todoApi.POST("item/", controllers.TodoAddItem)
			todoApi.GET("item/:id", controllers.TodoShowItem)
			todoApi.POST("item/:id", controllers.TodoAddSubItem)
			todoApi.DELETE("item/:id", controllers.TodoDeleteItem)
		}
	}
	// not in production
	router.GET("/swagger/*any", controllers.GinModeReleaseSwagger(swaggerFiles.Handler))
	return router
}
