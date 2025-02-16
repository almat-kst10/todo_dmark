package handler

import (
	"net/http"
	"todo_dmark/internal/service"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Handler struct {
	TodoService service.ITodoService
}

func NewHandler(TodoService service.ITodoService) *Handler {
	return &Handler{TodoService: TodoService}
}

func (h *Handler) InitRouter(router *mux.Router) {
	h.InitTodoRouter(router)
}


func (h *Handler) Init() http.Handler {
	router := mux.NewRouter()
	h.InitRouter(router)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Разрешенный фронтенд
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// return router
	return c.Handler(router)
}
