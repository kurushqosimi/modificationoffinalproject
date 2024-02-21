package Kurush

import (
	"net/http"
	"one-day-job/internal/domain/Kurush"
	"one-day-job/pkg/utils"
)

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSONResponse(w, http.StatusOK, Kurush.Response{Status: http.StatusOK, Message: "Service Status OK"})
}
