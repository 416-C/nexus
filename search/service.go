package search

import (
	"errors"

	assetsapiv1 "github.com/416-C/nexus-client-go/assets/apiv1"
	componentsapiv1 "github.com/416-C/nexus-client-go/components/apiv1"
	searchapiv1 "github.com/416-C/nexus-client-go/search/apiv1"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type Search struct {
	client *resty.Client
}

func NewSearchService(client *resty.Client) *Search {
	return &Search{
		client: client,
	}
}

func (s *Search) SearchComponents(params *searchapiv1.Components) (*componentsapiv1.ComponentList, error) {
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	response, err := s.client.NewRequest().SetResult(&componentsapiv1.ComponentList{}).SetQueryParamsFromValues(v).Post("service/rest/v1/search")
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*componentsapiv1.ComponentList), nil
}

func (s *Search) SearchAssets(params *searchapiv1.Assets) (*assetsapiv1.AssetList, error) {
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	response, err := s.client.NewRequest().SetResult(&assetsapiv1.AssetList{}).SetQueryParamsFromValues(v).Post("service/rest/v1/search/assets")
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*assetsapiv1.AssetList), nil
}
