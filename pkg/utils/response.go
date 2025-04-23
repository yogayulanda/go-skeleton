package utils

import (
	"encoding/json"
	"net/http"
)

// Status tipe untuk response
type ResponseStatus string

const (
	Success ResponseStatus = "success"
	Error   ResponseStatus = "error"
)

// ResponseCode bisa menggunakan type alias string atau enum
type ResponseCode string

// Response adalah format standar
type Response struct {
	Status  ResponseStatus    `json:"status,omitempty"`
	Code    ResponseCode      `json:"code,omitempty"`
	Message string            `json:"message,omitempty"`
	Data    interface{}       `json:"data,omitempty"`
	Detail  map[string]string `json:"detail,omitempty"` // Optional untuk field error yang detail
}

// SetSuccess mengatur response sukses
func (r *Response) SetSuccess(code ResponseCode, data interface{}) {
	r.Status = Success
	r.Code = code
	r.Message = GetMessageForCode(code)
	r.Data = data
}

// SetError mengatur response error
func (r *Response) SetError(code ResponseCode, err error, detail map[string]string) {
	r.Status = Error
	r.Code = code
	r.Message = GetMessageForCode(code)
	r.Detail = detail
}

// WriteJSON untuk kirim response
func WriteJSON(w http.ResponseWriter, statusCode int, response *Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// GetMessageForCode mengembalikan pesan untuk kode tertentu
func GetMessageForCode(code ResponseCode) string {
	//TODO: Buat mapping dari database
	switch code {
	case "200":
		return "OK"
	case "400":
		return "Bad Request"
	case "401":
		return "Unauthorized"
	case "403":
		return "Forbidden"
	case "404":
		return "Not Found"
	case "500":
		return "Internal Server Error"
	default:
		return "Unknown Error"
	}
}
