package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"taskMetrics/internal/task"
)

type TaskHandler struct {
	taskUseCase task.TaskUseCase
}

func NewTaskHandler(uc task.TaskUseCase) *TaskHandler {
	return &TaskHandler{
		taskUseCase: uc,
	}
}

// Create cria uma nova tarefa
func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID int64  `json:"user_id"`
		Title  string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.UserID <= 0 || req.Title == "" {
		http.Error(w, "user_id and title are required", http.StatusBadRequest)
		return
	}

	task, err := h.taskUseCase.CreateTask(r.Context(), req.UserID, req.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// UpdateStatus atualiza o status de uma tarefa
func (h *TaskHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "id parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id format", http.StatusBadRequest)
		return
	}

	var req struct {
		Status string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.Status == "" {
		http.Error(w, "status is required", http.StatusBadRequest)
		return
	}

	err = h.taskUseCase.UpdateTaskStatus(r.Context(), id, req.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetTask retorna uma tarefa pelo ID
func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "id parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id format", http.StatusBadRequest)
		return
	}

	task, err := h.taskUseCase.GetTask(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
