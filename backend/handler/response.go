package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func RespondJSON(ctx context.Context, w http.ResponseWriter, body any, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		rsp := ErrorResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		}
		err := json.NewEncoder(w).Encode(rsp)
		if err != nil {
			fmt.Printf("write error response error: %v", err)
		}
		return
	}

	w.WriteHeader(status)
	_, err = fmt.Fprintf(w, "%s", bodyBytes)
}
