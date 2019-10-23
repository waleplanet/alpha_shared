package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type (
	ErrorObj struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HttpStatus int    `json:"status"`
	}

	HttpResponse struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	}
)

func GetHeaders(r *http.Request, headers ...string) map[string]interface{} {
	var requestHeaders map[string]interface{}

	requestHeaders = make(map[string]interface{})

	var value string
	for _, headerName := range headers {
		value = r.Header.Get(headerName)
		value = strings.TrimSpace(value)
		if len(value) == 0 {
			requestHeaders[headerName] = nil
			continue
		}
		requestHeaders[headerName] = value
	}

	return requestHeaders
}
func DisplayAppError(w http.ResponseWriter, err error, message string, code int) {
	errObj := ErrorObj{
		Error:      err.Error(),
		Message:    message,
		HttpStatus: code,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errObj); err == nil {
		w.Write(j)
	}
}

func WriteJsonResponse(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	j, err := json.Marshal(data)
	if err != nil {
		DisplayAppError(w, err, "Json marshal error", http.StatusForbidden)
	}
	w.Write(j)
}

func DecodeRequestData(w http.ResponseWriter, r *http.Request, payload interface{}) error {
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		fmt.Fprintf(w, "err decoding request :%s \n", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	return err
}
func DisplayApiError(w http.ResponseWriter, message, response_code string, data interface{}, code int) {
	errObj := struct {
		Message      string
		Data         interface{}
		ResponseCode string
	}{
		message,
		data,
		response_code,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errObj); err == nil {
		w.Write(j)
	}
}
