package SprintlyAPI

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	// DefaultBaseURL is the default base URL for the Sprint.ly API.
	DefaultBaseURL = "https://sprint.ly/api/"
)

type API struct {
	username string
	token    string
	client   *http.Client
	baseURL  *url.URL
}

func Create(username string, token string) *API {
	baseURL, _ := url.Parse(DefaultBaseURL)

	api := &API{
		username: username,
		token:    token,
		client:   http.DefaultClient,
		baseURL:  baseURL,
	}

	return api
}

func (a *API) Product() *ProductService {
	return &ProductService{api: a}
}

func (a *API) NewRequest(method string, endpoint string) (*APIRequest, error) {
	path, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	url := a.baseURL.ResolveReference(path)

	req, err := http.NewRequest(method, url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(a.username, a.token)

	return &APIRequest{request: req}, nil
}

func (a *API) Do(apiRequest *APIRequest) (*http.Response, error) {
	req := apiRequest.request
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return resp, &APIError{
			Response: resp,
		}
	}

	apiRequest.SetResponse(resp)

	decoder := json.NewDecoder(resp.Body)
	parseErr := decoder.Decode(&apiRequest.result)

	if parseErr != nil {
		return resp, parseErr
	}

	return resp, err
}
