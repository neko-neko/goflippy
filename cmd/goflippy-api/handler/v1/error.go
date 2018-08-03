package v1

import (
	"encoding/json"
	"net/http"

	validator "gopkg.in/go-playground/validator.v9"
)

// errorStruct has description error
type errorStruct struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

// errorResponse is a error response structure
// swagger:response errorResponse
type errorResponse struct {
	Message string        `json:"message"`
	Errors  []errorStruct `json:"errors"`
}

// newErrorResponse returns new ErrorResponse
func newErrorResponse(mes string) *errorResponse {
	return &errorResponse{
		Message: mes,
		Errors:  make([]errorStruct, 0),
	}
}

// ErrorHandler is handle server error func
func ErrorHandler(code int, err error, w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)

	res := newErrorResponse(err.Error())
	switch e := err.(type) {
	case validator.ValidationErrors:
		for _, ve := range e {
			res.Errors = append(res.Errors, errorStruct{
				Field: ve.Field(),
				Value: ve.Param(),
			})
		}
		break
	default:
		res.Errors = append(res.Errors, errorStruct{
			Field: "",
			Value: "",
		})
		break
	}

	json.NewEncoder(w).Encode(res)
}
