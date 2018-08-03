package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"
	"github.com/neko-neko/goflippy"
	"github.com/neko-neko/goflippy/cmd/goflippy-api/ctx"
	"github.com/neko-neko/goflippy/cmd/goflippy-api/handler"
	"github.com/neko-neko/goflippy/cmd/goflippy-api/handler/v1"
	"github.com/neko-neko/goflippy/cmd/goflippy-api/service"
	"github.com/neko-neko/goflippy/log"
	"github.com/neko-neko/goflippy/middleware"
	"github.com/neko-neko/goflippy/repository"
	"github.com/neko-neko/goflippy/store"
	"github.com/neko-neko/goflippy/util"
)

// run application
func run() int {
	// load env
	err := EnvInit()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load environment variables %v", err)
		return 1
	}
	level.Debug(log.Logger).Log("message", fmt.Sprintf("loaded environment %#v", Spec))

	level.Info(log.Logger).Log("message", "initializing goflippy API server...")
	// connect DB
	level.Info(log.Logger).Log("message", "connecting store...")
	store := store.Init(store.Configuration{
		URL: Spec.StoreURL,
		DB:  Spec.DB,
	})
	err = util.Retry(5, store.Connect)
	if err != nil {
		level.Error(log.Logger).Log("message", "failed to connect data store", "err", err)
		return 1
	}
	defer store.Close()
	level.Info(log.Logger).Log("message", "connected store!")

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
	r.Handle("/v1/features", goflippy.Handler(featureHandler.GetFeatures, v1.ErrorHandler)).Methods("GET")
	r.Handle("/v1/features/{key}", goflippy.Handler(featureHandler.GetFeature, v1.ErrorHandler)).Methods("GET")
	r.Handle("/v1/users", goflippy.Handler(userHandler.PostUsers, v1.ErrorHandler)).Methods("POST")
	r.Handle("/v1/users/{uuid}", goflippy.Handler(userHandler.PatchUsers, v1.ErrorHandler)).Methods("PATCH")
	r.Handle("/v1/users/{uuid}", goflippy.Handler(userHandler.DeleteUsers, v1.ErrorHandler)).Methods("DELETE")
	r.Handle("/v1/users/{uuid}/features/{key}", goflippy.Handler(userFeatureHandler.GetFeatures, v1.ErrorHandler)).Methods("GET")

	// register middlewares
	r.Use(middleware.NewRecoverMiddleware(handler.RecoverErrorHandler).Middleware)
	r.Use(middleware.NewKeyAuthMiddleware(projectRepo.FindProjectIDByAPIKey, ctx.CreateRequestWithContext, handler.AuthErrorHandler).Middleware)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("0.0.0.0:%d", Spec.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			level.Error(log.Logger).Log("message", "server got an error", "err", err)
		}
	}()
	level.Info(log.Logger).Log("message", fmt.Sprintf("goflippy API server started listen %s", srv.Addr))

	// signal handler
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	level.Info(log.Logger).Log("message", "shutting down")

	return 0
}

func main() {
	os.Exit(run())
}
