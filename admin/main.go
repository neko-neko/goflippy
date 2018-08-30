package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/neko-neko/goflippy/admin/handler"
	"github.com/neko-neko/goflippy/admin/service"
	baseHandler "github.com/neko-neko/goflippy/pkg/handler"
	"github.com/neko-neko/goflippy/pkg/log"
	"github.com/neko-neko/goflippy/pkg/middleware"
	"github.com/neko-neko/goflippy/pkg/repository"
	"github.com/neko-neko/goflippy/pkg/store"
	"github.com/neko-neko/goflippy/pkg/util"
)

// expose port
const port = 9001

// run application
func run() int {
	// load env
	err := EnvInit()
	if err != nil {
		log.ErrorWithErr(err, "message", "failed to load environment variables")
		return 1
	}
	log.Debug("message", fmt.Sprintf("loaded environment %#v", Spec))

	log.Info("message", "initializing goflippy Admin server...")
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
	projectService := service.NewProjectService(projectRepo)
	featureService := service.NewFeatureService(featureRepo, userRepo)
	userService := service.NewUserService(userRepo)

	// initialize handler
	projectHandler := handler.NewProjectHandler(projectService)
	projectFeatureHandler := handler.NewProjectFeatureHandler(featureService)
	projectUserHandler := handler.NewProjectUserHandler(userService)

	// initialize router
	r := mux.NewRouter()
	r.Handle("/v1/projects", baseHandler.Handler(projectHandler.GetProjects, handler.ErrorHandler)).Methods("GET")
	r.Handle("/v1/projects", baseHandler.Handler(projectHandler.PostProjects, handler.ErrorHandler)).Methods("POST")
	r.Handle("/v1/projects/{id}", baseHandler.Handler(projectHandler.GetProject, handler.ErrorHandler)).Methods("GET")
	r.Handle("/v1/projects/{id}", baseHandler.Handler(projectHandler.PatchProjects, handler.ErrorHandler)).Methods("PATCH")
	r.Handle("/v1/projects/{id}/features", baseHandler.Handler(projectFeatureHandler.GetFeatures, handler.ErrorHandler)).Methods("GET")
	r.Handle("/v1/projects/{id}/features/{key}", baseHandler.Handler(projectFeatureHandler.GetFeature, handler.ErrorHandler)).Methods("GET")
	r.Handle("/v1/projects/{id}/features", baseHandler.Handler(projectFeatureHandler.PostFeature, handler.ErrorHandler)).Methods("POST")
	r.Handle("/v1/projects/{id}/users", baseHandler.Handler(projectUserHandler.GetUsers, handler.ErrorHandler)).Methods("GET")

	// register middlewares
	r.Use(middleware.NewRecoverMiddleware(handler.RecoverErrorHandler).Middleware)

	srv := &http.Server{
		Handler: handlers.CORS(
			handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}),
			handlers.ExposedHeaders([]string{baseHandler.HTTPHeaderXApiKey}),
			handlers.AllowedHeaders([]string{"accept", "accept-language", "content-type", baseHandler.HTTPHeaderXApiKey}),
			handlers.AllowedOrigins(Spec.AllowOrigins),
		)(r),
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
	log.Info("message", fmt.Sprintf("goflippy Admin server started listen %s", srv.Addr))

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