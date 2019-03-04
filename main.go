package main

import (
	"context"
	"github.com/balkin/todolist/controllers"
	"github.com/balkin/todolist/todo"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//HTTP Server
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

func main() {
	log.SetFlags(log.LstdFlags)
	router := gin.Default()
	router.GET("/", controllers.IndexController)
	todo.ConnectToDatabase()

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
			os.Exit(0)
		case syscall.SIGTERM:
		case syscall.SIGHUP:
			log.Println("Gracefully stopping HTTP server")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := HttpDaemon.Shutdown(ctx); err != nil {
				log.Fatal("Error during HTTP server shutdown:", err)
			}
			go todo.DisconnectDatabase()
			select {
			case <-ctx.Done():
				log.Println("Stopped HTTP server after waiting 5 seconds")
			}
		}
	}
}
