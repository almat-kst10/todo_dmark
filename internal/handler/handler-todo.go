package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
	"todo_dmark/internal/models"

	"github.com/gorilla/mux"
)

func (h *Handler) InitTodoRouter(router *mux.Router) {
	router.HandleFunc("/api/v1/list", h.todoList).Methods("GET")
	router.HandleFunc("/api/v1/list", h.todoAdd).Methods("POST")
	router.HandleFunc("/api/v1/list/{id}", h.todoDetail).Methods("GET")
	router.HandleFunc("/api/v1/list/{id}", h.todoUpdate).Methods("PUT")
	router.HandleFunc("/api/v1/list/{id}", h.todoDelete).Methods("DELETE")
}

type TodoBody struct {
	Title       string `json:"title"`
	IsDone     	bool   `json:"is_done"`
}

type Response struct {
	Result interface{} `json:"result"`
}

func (h *Handler) todoList(w http.ResponseWriter, r *http.Request) {
	list, err := h.TodoService.List(r.Context())
	if err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}

	h.writeJSONResponse(w, http.StatusOK, list)
}

func (h *Handler) todoAdd(w http.ResponseWriter, r *http.Request) {
	var body TodoBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}

	todo := &models.TodoItem{
		Title:       body.Title,
		TimeAt:      time.Now(),
		IsDone: 	 false,
	}
	log.Println(todo, "todo")
	err = h.TodoService.Add(r.Context(), todo)
	if err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}
	
	responseMessage := &Response{Result: "Add success"}
	h.writeJSONResponse(w, http.StatusOK, responseMessage)
}

func (h *Handler) todoDetail(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}

	detail, err := h.TodoService.Detail(r.Context(), id)
	if err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}

	h.writeJSONResponse(w, http.StatusOK, detail)
}

func (h *Handler) todoUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}

	var body TodoBody
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}

	itemBody := &models.TodoItem{
		Id:          id,
		Title:       body.Title,
		TimeAt:      time.Now(),
		IsDone: 	 body.IsDone,
	}

	isDone, err := h.TodoService.Update(r.Context(), *itemBody)
	if err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}

	responseMessage := &Response{Result: isDone}
	h.writeJSONResponse(w, http.StatusOK, responseMessage)
}

func (h *Handler) todoDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}

	err = h.TodoService.Delete(r.Context(), id)
	if err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}

	responseMessage := &Response{Result: "Deleted success"}
	h.writeJSONResponse(w, http.StatusOK, responseMessage)
}
