package maven

import (
	"errors"
	"github.com/416-C/nexus-client-go/repositorymanagement/maven/apiv1"
	"github.com/go-resty/resty/v2"
)

type Maven struct {
	client *resty.Client
}

func NewMavenService(client *resty.Client) *Maven {
	return &Maven{
		client: client,
	}
}

func (s *Maven) GetMavenLocalRepo(repositoryName string) (*apiv1.MavenLocalRepository, error) {
	response, err := s.client.NewRequest().SetResult(&apiv1.MavenLocalRepository{}).Get("service/rest/v1/repositories/maven/hosted/" + repositoryName)
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*apiv1.MavenLocalRepository), nil
}

func (s *Maven) UpdateMavenLocalRepo(repositoryName string, r *apiv1.MavenLocalApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Put("service/rest/v1/repositories/maven/hosted/" + repositoryName)
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}

func (s *Maven) CreateMavenLocalRepo(r *apiv1.MavenLocalApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Post("service/rest/v1/repositories/maven/hosted")
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}

func (s *Maven) GetMavenProxyRepo(repositoryName string) (*apiv1.MavenProxyRepository, error) {
	response, err := s.client.NewRequest().SetResult(&apiv1.MavenProxyRepository{}).Get("service/rest/v1/repositories/maven/proxy/" + repositoryName)
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*apiv1.MavenProxyRepository), nil
}

func (s *Maven) UpdateMavenProxyRepo(repositoryName string, r *apiv1.MavenProxyApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Put("service/rest/v1/repositories/maven/proxy/" + repositoryName)
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}

func (s *Maven) CreateMavenProxyRepo(r *apiv1.MavenProxyApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Post("service/rest/v1/repositories/maven/proxy")
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}
