package handlers

import (
	"encoding/json"

	"net/http"
	"time"

	"github.com/gorilla/mux"
	usecases "github.com/rodrinoblega/tallertechnologies/use_cases"
)

type EventHandler struct {
	uc *usecases.EventUseCase
}

func NewEventHandler(uc *usecases.EventUseCase) *EventHandler {
	return &EventHandler{uc: uc}
}

func (h *EventHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/events", h.CreateEvent).Methods("POST")
	r.HandleFunc("/events", h.ListEvents).Methods("GET")
	r.HandleFunc("/events/{id}", h.GetEventByID).Methods("GET")
}

func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		StartTime   time.Time `json:"start_time"`
		EndTime     time.Time `json:"end_time"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	e, err := h.uc.CreateEvent(r.Context(), req.Title, req.Description, req.StartTime, req.EndTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(e)
}

func (h *EventHandler) ListEvents(w http.ResponseWriter, r *http.Request) {
	events, err := h.uc.ListEvents(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(events)
}

func (h *EventHandler) GetEventByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	event, err := h.uc.GetEventByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(event)
}
