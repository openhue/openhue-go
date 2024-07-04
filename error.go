package openhue

import (
	"fmt"
	"net/http"
)

type ApiError struct {
	error
	StatusCode int
}

func (a *ApiError) Error() string {

	if a.StatusCode == http.StatusForbidden {
		return "openhue api error: wrong API key"
	}

	return fmt.Sprintf("openhue api error: %d", a.StatusCode)
}

type apiResponse interface {
	Status() string
	StatusCode() int
}

func newApiError(resp apiResponse) error {
	return &ApiError{
		StatusCode: resp.StatusCode(),
	}
}
