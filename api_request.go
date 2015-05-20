package SprintlyAPI

import (
	"net/http"
)

type APIRequest struct {
	request  *http.Request
	response *http.Response
	result   interface{}
}

func (a *APIRequest) SetResponse(resp *http.Response) {
	a.response = resp
}
