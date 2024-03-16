package main

import (
	"Glossika_interview/database"
	"Glossika_interview/database/models"
	"Glossika_interview/middleware"
	_ "Glossika_interview/myvalidators"
	"Glossika_interview/router"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	slog.New(slog.NewTextHandler(os.Stderr, nil))

	g := gin.Default()

	db := database.Connection()
	db.Set("gorm:table_options", "charset=utf8").AutoMigrate(&models.User{})

	g.Use(middleware.DBMiddleware(db))
	g.Use(middleware.ErrorHandler)
	router.Route(g)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: g,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("listen: \n", err)
			panic(err)
		}
	}()

	gracefulShutdown(srv)
}

// copy from https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-without-context/server.go
// gracefulShutdown is a function to handle the graceful shutdown of the server
func gracefulShutdown(srv *http.Server) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown: ", err)
		panic("Server forced to shutdown: " + err.Error())
	}

	slog.Info("Server exiting")
}
