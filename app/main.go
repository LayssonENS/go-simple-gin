package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	BooksHttpDelivery "github.com/LayssonENS/go-simple-gin/books/delivery/http"
	BooksRepository "github.com/LayssonENS/go-simple-gin/books/repository"
	BooksUCase "github.com/LayssonENS/go-simple-gin/books/usecase"
	"github.com/LayssonENS/go-simple-gin/config"
	"github.com/LayssonENS/go-simple-gin/config/database"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title go-simple-gin API
// @version 1.0
// @description This is go-simple-gin API in Go.

func main() {
	ctx := context.Background()
	log := logrus.New()

	dbInstance, err := database.NewPostgresConnection()
	if err != nil {
		log.WithError(err).Fatal("failed connection database")
		return
	}

	//err = database.DBMigrate(dbInstance, config.GetEnv().DbConfig)
	//if err != nil {
	//	log.WithError(err).Fatal("failed to migrate")
	//	return
	//}

	router := gin.Default()

	booksRepository := BooksRepository.NewPostgresBooksRepository(dbInstance)
	booksService := BooksUCase.NewBooksUseCase(booksRepository)

	BooksHttpDelivery.NewBooksHandler(router, booksService)
	router.GET("/go-simple-gin/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	gin.SetMode(gin.ReleaseMode)
	if config.GetEnv().Debug {
		gin.SetMode(gin.DebugMode)
	}

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.GetEnv().Port),
		Handler: router,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down API...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("API Server forced to shutdown:", err)
	}

	log.Println("API Server exiting")
}
