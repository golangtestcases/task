package server

import (
	"net/http"

	"https://github.com/golangtestcases/task/internal/controller/task"
)

type Server struct {
	taskHandler *task.Handler
}

func NewServer(th *task.Handler) *Server {
	return &Server{taskHandler: th}
}

func (s *Server) SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/tasks", s.taskHandler.CreateTask)
	mux.HandleFunc("/api/v1/tasks/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			s.taskHandler.GetTask(w, r)
		} else if r.Method == http.MethodDelete {
			s.taskHandler.DeleteTask(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}

func (s *Server) Run(addr string) error {
	return http.ListenAndServe(addr, s.SetupRoutes())
}
