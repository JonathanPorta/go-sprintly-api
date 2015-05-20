package SprintlyAPI

import (
	"fmt"
	"net/http"
)

type APIError struct {
	Response *http.Response
}

func (err *APIError) Error() string {
	req := err.Response.Request
	return fmt.Sprintf("%v %v -> %v", req.Method, req.URL, err.Response.Status)
}
