package components

import (
	"errors"

	"github.com/416-C/nexus-client-go/components/apiv1"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type Components struct {
	client *resty.Client
}

func NewComponentsService(client *resty.Client) *Components {
	return &Components{
		client: client,
	}
}

func (s *Components) DeleteComponent(componentID string) error {
	response, err := s.client.NewRequest().Delete("service/rest/v1/components/" + componentID)
	if err != nil {
		return err
	}
	if response.IsError() && response.StatusCode() != 404 {
		return errors.New(response.String())
	}

	return nil
}

func (s *Components) GetComponent(componentID string) (*apiv1.Component, error) {
	response, err := s.client.NewRequest().SetResult(&apiv1.Component{}).Get("service/rest/v1/components/" + componentID)
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*apiv1.Component), nil
}

func (s *Components) ListComponents(params *apiv1.ListComponentsParameters) (*apiv1.ComponentList, error) {
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	response, err := s.client.NewRequest().SetResult(&apiv1.ComponentList{}).SetQueryParamsFromValues(v).Get("service/rest/v1/components")
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*apiv1.ComponentList), nil
}
