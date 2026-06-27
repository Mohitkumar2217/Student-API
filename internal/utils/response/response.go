package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/MohitKumar2217/Students-api/internal/types"
	"github.com/go-playground/validator"
)

const (
	StatusOK = "OK"
	StatusError = "Error"
)
func WriteJson(w http.ResponseWriter, status int, data interface {}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) types.Response {
	return types.Response{
		Status: StatusError,
		Error: err.Error(),
	}
}

func ValidatorError(errs validator.ValidationErrors) types.Response {
	var errMsgs []string
	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("%s is required field", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("%s invalid field", err.Field()))
		}
	}
	return types.Response{
		Status: StatusError,
		Error: strings.Join(errMsgs, ", "),
	}
}