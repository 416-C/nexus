package pypi

import (
	"errors"

	"github.com/416-C/nexus-client-go/repositorymanagement/common/model"
	"github.com/416-C/nexus-client-go/repositorymanagement/pypi/apiv1"

	"github.com/go-resty/resty/v2"
)

type PyPI struct {
	client *resty.Client
}

func NewPypiService(client *resty.Client) *PyPI {
	return &PyPI{
		client: client,
	}
}

func (s *PyPI) GetPyPILocalRepo(repositoryName string) (*model.LocalRepository, error) {
	response, err := s.client.NewRequest().SetResult(&model.LocalRepository{}).Get("service/rest/v1/repositories/pypi/hosted/" + repositoryName)
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*model.LocalRepository), nil
}

func (s *PyPI) UpdatePyPILocalRepo(repositoryName string, r *apiv1.PyPILocalApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Put("service/rest/v1/repositories/pypi/hosted/" + repositoryName)
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}

func (s *PyPI) CreatePyPILocalRepo(r *apiv1.PyPILocalApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Post("service/rest/v1/repositories/pypi/hosted")
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}

func (s *PyPI) GetPyPIProxyRepo(repositoryName string) (*model.ProxyRepository, error) {
	response, err := s.client.NewRequest().SetResult(&model.ProxyRepository{}).Get("service/rest/v1/repositories/pypi/proxy/" + repositoryName)
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*model.ProxyRepository), nil
}

func (s *PyPI) UpdatePyPIProxyRepo(repositoryName string, r *apiv1.PyPIProxyApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Put("service/rest/v1/repositories/pypi/proxy/" + repositoryName)
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}

func (s *PyPI) CreatePyPIProxyRepo(r *apiv1.PyPIProxyApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Post("service/rest/v1/repositories/pypi/proxy")
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}
