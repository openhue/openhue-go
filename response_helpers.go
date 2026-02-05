package openhue

import (
	"net/http"
	"reflect"
)

// extractErrorFromResponse attempts to extract error details from any response type.
// It uses reflection to check for common error fields (JSON4xx, JSON5xx) present
// in the generated response structs.
func extractErrorFromResponse(resp any, statusCode int) *ErrorResponse {
	v := reflect.ValueOf(resp)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}

	// Map status codes to field names
	fieldName := statusCodeToFieldName(statusCode)
	if fieldName == "" {
		return nil
	}

	field := v.FieldByName(fieldName)
	if !field.IsValid() || field.IsNil() {
		return nil
	}

	// All error response types (Forbidden, NotFound, etc.) are type aliases for ErrorResponse.
	// We can safely convert via the underlying pointer type.
	if field.CanInterface() {
		// Use reflection to get the underlying value since all types are aliases
		if field.Type().Kind() == reflect.Ptr && field.Type().Elem().Name() == "ErrorResponse" {
			if errResp, ok := field.Interface().(*ErrorResponse); ok {
				return errResp
			}
		}

		// For type aliases, convert the underlying pointer
		if field.Elem().CanAddr() {
			ptr := field.Elem().Addr()
			if errResp, ok := ptr.Interface().(*ErrorResponse); ok {
				return errResp
			}
		}
	}

	return nil
}

// statusCodeToFieldName maps HTTP status codes to the corresponding field name
// in generated response structs.
func statusCodeToFieldName(statusCode int) string {
	switch statusCode {
	case http.StatusUnauthorized:
		return "JSON401"
	case http.StatusForbidden:
		return "JSON403"
	case http.StatusNotFound:
		return "JSON404"
	case http.StatusMethodNotAllowed:
		return "JSON405"
	case http.StatusNotAcceptable:
		return "JSON406"
	case http.StatusConflict:
		return "JSON409"
	case http.StatusTooManyRequests:
		return "JSON429"
	case http.StatusInternalServerError:
		return "JSON500"
	case http.StatusServiceUnavailable:
		return "JSON503"
	case http.StatusInsufficientStorage:
		return "JSON507"
	default:
		return ""
	}
}
