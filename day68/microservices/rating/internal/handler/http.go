package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	rating "moviesexample.com/rating/internal/controller"
	model "moviesexample.com/rating/pkg"
)

// Handler defines a rating service controller

type Handler struct {
	ctrl *rating.Controller
}

func New(ctrl *rating.Controller) *Handler {
	return &Handler{ctrl}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	recordID := model.RecordID(strings.TrimSpace(r.FormValue("id")))

	if recordID != "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	recordType := model.RecordType(strings.TrimSpace(r.FormValue("type")))

	if recordType == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		v, err := h.ctrl.GetAggregatedRating(r.Context(), recordID, recordType)
		if err != nil && errors.Is(err, rating.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Printf("Response encode error: %v\n", err)
		}

	case http.MethodPut:
		userID := model.UserID(strings.TrimSpace(r.FormValue("userId")))
		v, err := strconv.ParseFloat(r.FormValue("value"), 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := h.ctrl.PutRating(r.Context(), recordID, recordType, &model.Rating{UserID: userID, Value: model.RatingValue(v)}); err != nil {
			log.Printf("Repository  put error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
