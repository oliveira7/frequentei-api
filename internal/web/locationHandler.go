package web

import (
	"encoding/json"
	"frequentei-api/internal/usecase"
	"net/http"
)

type LocationHandler struct {
	CreateLocationUseCase *usecase.CreateLocationUseCase
	ListLocationsUseCase  *usecase.ListLocationUseCase
}

func NewLocationHandler(createLocationUseCase *usecase.CreateLocationUseCase, listLocationsUseCase *usecase.ListLocationUseCase) *LocationHandler {
	return &LocationHandler{
		CreateLocationUseCase: createLocationUseCase,
		ListLocationsUseCase:  listLocationsUseCase,
	}
}

func (h *LocationHandler) ListLocationsHandler(w http.ResponseWriter, r *http.Request) {
	output, err := h.ListLocationsUseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (h *LocationHandler) CreateLocationHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateLocationInputDto

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := h.CreateLocationUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}
