package SprintlyAPI

import (
	"net/http"
	"time"
)

type Product struct {
	ID        int        `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Admin     bool       `json:"admin,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Webhook   string     `json:"webhook,omitempty"`
	Archived  bool       `json:"archived,omitempty"`
}

type ProductService struct {
	api *API
}

func (service *ProductService) List() ([]Product, *http.Response, error) {
	endpoint := "products.json"

	apiRequest, err := service.api.NewRequest("GET", endpoint)
	if err != nil {
		return nil, nil, err
	}

	var product []Product
	apiRequest.result = &product
	resp, err := service.api.Do(apiRequest)
	if err != nil {
		return nil, resp, err
	}

	return product, resp, nil
}
