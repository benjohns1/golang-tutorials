package middleware

import (
	"net/http"
)

// RequestHandler logic for requests that returns the response data as a struct
type RequestHandler func(*http.Request) interface{}

// Formatter formats data and applies it to an http response
type Formatter func(w http.ResponseWriter, r *http.Request, data interface{})

// ResponseFormatter formats a response
type ResponseFormatter func(handler RequestHandler) http.HandlerFunc

// NewResponseFormatter creates a function to format responses
func NewResponseFormatter(formatter Formatter) ResponseFormatter {
	return func(handler RequestHandler) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			formatter(w, r, handler(r))
		})
	}
}
