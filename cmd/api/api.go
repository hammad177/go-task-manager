package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hammad177/task_management/modules/projects"
	"github.com/hammad177/task_management/modules/tasks"
	"github.com/hammad177/task_management/modules/users"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	subRouter := router.PathPrefix("/api/v1").Subrouter()

	projectsStore := projects.NewStore(s.db)
	projectsHandler := projects.NewHandler(projectsStore)
	projectsHandler.RegisterRoutes(subRouter)

	taskStore := tasks.NewStore(s.db)
	taskHandler := tasks.NewHandler(taskStore)
	taskHandler.RegisterRoutes(subRouter)

	userStore := users.NewStore(s.db)
	userHandler := users.NewHandler(userStore)
	userHandler.RegisterRoutes(subRouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
