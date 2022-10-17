package splitwise

import (
	"context"
	"encoding/json"
	"net/http"
)

// Category resources to access category information.
type Categories interface {
	// Categories returns list of available categories
	Categories(ctx context.Context) ([]Category, error)
}

type Category struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	IconTypes struct {
		Slim struct {
			Small string `json:"small"`
			Large string `json:"large"`
		} `json:"slim"`
		Square struct {
			Large  string `json:"large"`
			Xlarge string `json:"xlarge"`
		} `json:"square"`
	} `json:"icon_types"`
	Subcategories []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Icon      string `json:"icon"`
		IconTypes struct {
			Slim struct {
				Small string `json:"small"`
				Large string `json:"large"`
			} `json:"slim"`
			Square struct {
				Large  string `json:"large"`
				Xlarge string `json:"xlarge"`
			} `json:"square"`
		} `json:"icon_types"`
	} `json:"subcategories"`
}

type categoriesResponse struct {
	Categories []Category `json:"categories"`
}

func (c client) Categories(ctx context.Context) ([]Category, error) {
	url := c.baseURL + "/api/v3.0/get_categories"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	token, err := c.AuthProvider.Auth()
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	var response categoriesResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.Categories, nil
}
