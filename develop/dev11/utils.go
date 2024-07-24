package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func GetParams(r *http.Request) (map[string]string, error) {
	params := make(map[string]string)
	r.ParseForm()
	for key, value := range r.Form {
		params[key] = value[0]
	}
	return params, nil
}

func ValidateParams(params map[string]string, required []string) error {
	for _, param := range required {
		value, ok := params[param]
		if !ok || value == "" {
			return errors.New("missing or empty parameter: " + param)
		}
	}
	return nil
}


func ValidateDate(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}

func ValidateInt(intStr string) (int, error) {
	return strconv.Atoi(intStr)
}
