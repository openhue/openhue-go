package openhue

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// ApiError represents an error returned by the Philips Hue API.
// It includes the HTTP status code and any error details from the response body.
type ApiError struct {
	StatusCode  int
	Status      string
	Description string
}

func (a *ApiError) Error() string {
	if a.Description != "" {
		return fmt.Sprintf("openhue api error (%d): %s", a.StatusCode, a.Description)
	}

	// Fallback messages for common status codes
	switch a.StatusCode {
	case http.StatusUnauthorized:
		return "openhue api error (401): unauthorized - invalid or missing API key"
	case http.StatusForbidden:
		return "openhue api error (403): forbidden - wrong API key"
	case http.StatusNotFound:
		return "openhue api error (404): resource not found"
	case http.StatusConflict:
		return "openhue api error (409): conflict"
	case http.StatusTooManyRequests:
		return "openhue api error (429): too many requests - rate limited"
	case http.StatusInternalServerError:
		return "openhue api error (500): internal server error"
	case http.StatusServiceUnavailable:
		return "openhue api error (503): service unavailable"
	default:
		return fmt.Sprintf("openhue api error (%d): %s", a.StatusCode, a.Status)
	}
}

// Sentinel errors for common API error conditions.
// Use errors.Is() to check for these specific error types.
var (
	ErrUnauthorized        = errors.New("unauthorized")
	ErrForbidden           = errors.New("forbidden")
	ErrNotFound            = errors.New("not found")
	ErrConflict            = errors.New("conflict")
	ErrTooManyRequests     = errors.New("too many requests")
	ErrInternalServerError = errors.New("internal server error")
	ErrServiceUnavailable  = errors.New("service unavailable")
	ErrEmptyResponse       = errors.New("no data returned from API")
)

// Is implements errors.Is for ApiError, allowing checks like:
//
//	if errors.Is(err, openhue.ErrNotFound) { ... }
func (a *ApiError) Is(target error) bool {
	switch a.StatusCode {
	case http.StatusUnauthorized:
		return errors.Is(target, ErrUnauthorized)
	case http.StatusForbidden:
		return errors.Is(target, ErrForbidden)
	case http.StatusNotFound:
		return errors.Is(target, ErrNotFound)
	case http.StatusConflict:
		return errors.Is(target, ErrConflict)
	case http.StatusTooManyRequests:
		return errors.Is(target, ErrTooManyRequests)
	case http.StatusInternalServerError:
		return errors.Is(target, ErrInternalServerError)
	case http.StatusServiceUnavailable:
		return errors.Is(target, ErrServiceUnavailable)
	default:
		return false
	}
}

// apiResponse is the interface that all generated response types implement.
type apiResponse interface {
	Status() string
	StatusCode() int
}

// newApiError creates an ApiError from an API response, extracting error details if available.
func newApiError(resp apiResponse) error {
	statusCode := resp.StatusCode()
	apiErr := &ApiError{
		StatusCode: statusCode,
		Status:     resp.Status(),
	}

	// Try to extract error description from the response using reflection
	if errResp := extractErrorFromResponse(resp, statusCode); errResp != nil && errResp.Errors != nil {
		apiErr.Description = extractErrorDescriptions(*errResp.Errors)
	}

	return apiErr
}

// extractErrorDescriptions joins all error descriptions into a single string.
func extractErrorDescriptions(errs []Error) string {
	if len(errs) == 0 {
		return ""
	}

	var descriptions []string
	for _, e := range errs {
		if e.Description != nil && *e.Description != "" {
			descriptions = append(descriptions, *e.Description)
		}
	}

	return strings.Join(descriptions, "; ")
}
