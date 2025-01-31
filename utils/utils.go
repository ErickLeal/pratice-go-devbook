package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func JSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

var validate = validator.New()

func ValidateStuctRequest(s interface{}) map[string]string {
	errors := make(map[string]string)

	err := validate.Struct(s)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errors[e.Field()] = e.Tag()
		}
	}

	return errors
}
