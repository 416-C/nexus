package helm

import (
	"errors"

	"github.com/416-C/nexus-client-go/repositorymanagement/common/model"
	"github.com/416-C/nexus-client-go/repositorymanagement/helm/apiv1"

	"github.com/go-resty/resty/v2"
)

type Helm struct {
	client *resty.Client
}

func NewHelmService(client *resty.Client) *Helm {
	return &Helm{
		client: client,
	}
}

func (s *Helm) GetHelmLocalRepo(repositoryName string) (*model.LocalRepository, error) {
	response, err := s.client.NewRequest().SetResult(&model.LocalRepository{}).Get("service/rest/v1/repositories/helm/hosted/" + repositoryName)
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*model.LocalRepository), nil
}

func (s *Helm) UpdateHelmLocalRepo(repositoryName string, r *apiv1.HelmLocalApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Put("service/rest/v1/repositories/helm/hosted/" + repositoryName)
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}

func (s *Helm) CreateHelmLocalRepo(r *apiv1.HelmLocalApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Post("service/rest/v1/repositories/helm/hosted")
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}

func (s *Helm) GetHelmProxyRepo(repositoryName string) (*model.ProxyRepository, error) {
	response, err := s.client.NewRequest().SetResult(&model.ProxyRepository{}).Get("service/rest/v1/repositories/helm/proxy/" + repositoryName)
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*model.ProxyRepository), nil
}

func (s *Helm) UpdateHelmProxyRepo(repositoryName string, r *apiv1.HelmProxyApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Put("service/rest/v1/repositories/helm/proxy/" + repositoryName)
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}

func (s *Helm) CreateHelmProxyRepo(r *apiv1.HelmProxyApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Post("service/rest/v1/repositories/helm/proxy")
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}
