package routes

import (
	"encoding/json"
	"net/http"
)

type ready struct {
	Data string
}

// ReadinessRoute - readiness check
func ReadinessRoute(w http.ResponseWriter, r *http.Request) {

	ready := ready{"Ok"}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&ready)
}
