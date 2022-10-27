package npm

import (
	"errors"

	"github.com/416-C/nexus-client-go/repositorymanagement/common/model"
	"github.com/416-C/nexus-client-go/repositorymanagement/npm/apiv1"

	"github.com/go-resty/resty/v2"
)

type Npm struct {
	client *resty.Client
}

func NewNpmService(client *resty.Client) *Npm {
	return &Npm{
		client: client,
	}
}

func (s *Npm) GetNpmLocalRepo(repositoryName string) (*model.LocalRepository, error) {
	response, err := s.client.NewRequest().SetResult(&model.LocalRepository{}).Get("service/rest/v1/repositories/npm/hosted/" + repositoryName)
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*model.LocalRepository), nil
}

func (s *Npm) UpdateNpmLocalRepo(repositoryName string, r *apiv1.NpmLocalApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Put("service/rest/v1/repositories/npm/hosted/" + repositoryName)
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}

func (s *Npm) CreateNpmLocalRepo(r *apiv1.NpmLocalApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Post("service/rest/v1/repositories/npm/hosted")
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}

func (s *Npm) GetNpmProxyRepo(repositoryName string) (*apiv1.NpmProxyRepository, error) {
	response, err := s.client.NewRequest().SetResult(&apiv1.NpmProxyRepository{}).Get("service/rest/v1/repositories/npm/proxy/" + repositoryName)
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*apiv1.NpmProxyRepository), nil
}

func (s *Npm) CreateNpmProxyRepo(repositoryName string, r *apiv1.NpmProxyRepository) error {
	response, err := s.client.NewRequest().SetBody(r).Post("service/rest/v1/repositories/npm/proxy/" + repositoryName)
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}

func (s *Npm) UpdateNpmProxyRepo(r *apiv1.NpmProxyRepository) error {
	response, err := s.client.NewRequest().SetBody(r).Put("service/rest/v1/repositories/npm/proxy/" + r.Name)
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}
