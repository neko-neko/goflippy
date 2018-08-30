package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/neko-neko/goflippy/api/ctx"
	"github.com/neko-neko/goflippy/api/handler"
	"github.com/neko-neko/goflippy/api/handler/v1"
	"github.com/neko-neko/goflippy/api/service"
	baseHandler "github.com/neko-neko/goflippy/pkg/handler"
	"github.com/neko-neko/goflippy/pkg/log"
	"github.com/neko-neko/goflippy/pkg/middleware"
	"github.com/neko-neko/goflippy/pkg/repository"
	"github.com/neko-neko/goflippy/pkg/store"
	"github.com/neko-neko/goflippy/pkg/util"
)

// expose port
const port = 9000

// run application
func run() int {
	// load env
	err := EnvInit()
	if err != nil {
		log.ErrorWithErr(err, "message", "failed to load environment variables")
		return 1
	}
	log.Debug("message", fmt.Sprintf("loaded environment %#v", Spec))

	log.Info("message", "initializing goflippy API server...")
	// connect DB
	log.Info("message", "connecting store...")
	store := store.Init(store.Configuration{
		TimeoutSeconds: Spec.StoreTimeoutSeconds,
		Addrs:          Spec.StoreAddrs,
		DB:             Spec.StoreDB,
		User:           Spec.StoreUser,
		Password:       Spec.StorePassword,
		Source:         Spec.StoreSource,
	})
	err = util.Retry(5, store.Connect)
	if err != nil {
		log.ErrorWithErr(err, "message", "failed to connect data store")
		return 1
	}
	defer store.Close()
	log.Info("message", "connected store!")

	// initialize repository
	userRepo := repository.NewUserRepositoryMongoDB(store)
	projectRepo := repository.NewProjectRepositoryMongoDB(store)
	featureRepo := repository.NewFeatureRepositoryMongoDB(store)

	// initialize service
	userService := service.NewUserService(userRepo)
	featureService := service.NewFeature(userRepo, featureRepo)

	// initialize handler
	featureHandler := v1.NewFeatureHandler(featureService)
	userHandler := v1.NewUserHandler(userService)
	userFeatureHandler := v1.NewUserFeatureHandler(featureService)

	// initialize router
	r := mux.NewRouter()
	r.Handle("/v1/features/{key}", baseHandler.Handler(featureHandler.GetFeature, v1.ErrorHandler)).Methods("GET")
	r.Handle("/v1/users", baseHandler.Handler(userHandler.PostUsers, v1.ErrorHandler)).Methods("POST")
	r.Handle("/v1/users/{uuid}", baseHandler.Handler(userHandler.PatchUsers, v1.ErrorHandler)).Methods("PATCH")
	r.Handle("/v1/users/{uuid}", baseHandler.Handler(userHandler.DeleteUsers, v1.ErrorHandler)).Methods("DELETE")
	r.Handle("/v1/users/{uuid}/features/{key}", baseHandler.Handler(userFeatureHandler.GetFeature, v1.ErrorHandler)).Methods("GET")

	// register middlewares
	r.Use(middleware.NewRecoverMiddleware(handler.RecoverErrorHandler).Middleware)
	r.Use(middleware.NewKeyAuthMiddleware(projectRepo.FindProjectIDByAPIKey, ctx.CreateRequestWithContext, handler.AuthErrorHandler).Middleware)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("0.0.0.0:%d", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.ErrorWithErr(err, "message", "server got an error")
		}
	}()
	log.Info("message", fmt.Sprintf("goflippy API server started listen %s", srv.Addr))

	// signal handler
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	log.Info("message", "shutting down")

	return 0
}

func main() {
	os.Exit(run())
}
