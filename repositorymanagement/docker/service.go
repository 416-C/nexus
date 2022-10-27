package docker

import (
	"errors"

	"github.com/416-C/nexus-client-go/repositorymanagement/docker/apiv1"

	"github.com/go-resty/resty/v2"
)

type Docker struct {
	client *resty.Client
}

func NewDockerService(client *resty.Client) *Docker {
	return &Docker{
		client: client,
	}
}

func (s *Docker) GetDockerLocalRepo(repositoryName string) (*apiv1.DockerLocalRepository, error) {
	response, err := s.client.NewRequest().SetResult(&apiv1.DockerLocalRepository{}).Get("service/rest/v1/repositories/docker/hosted/" + repositoryName)
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*apiv1.DockerLocalRepository), nil
}

func (s *Docker) UpdateDockerLocalRepo(repositoryName string, r *apiv1.DockerLocalApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Put("service/rest/v1/repositories/docker/hosted/" + repositoryName)
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}

func (s *Docker) CreateDockerLocalRepo(r *apiv1.DockerLocalApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Post("service/rest/v1/repositories/docker/hosted")
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}

func (s *Docker) GetDockerProxyRepo(repositoryName string) (*apiv1.DockerProxyRepository, error) {
	response, err := s.client.NewRequest().SetResult(&apiv1.DockerProxyRepository{}).Get("service/rest/v1/repositories/docker/proxy/" + repositoryName)
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*apiv1.DockerProxyRepository), nil
}

func (s *Docker) UpdateDockerProxyRepo(repositoryName string, r *apiv1.DockerProxyApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Put("service/rest/v1/repositories/docker/proxy/" + repositoryName)
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}

func (s *Docker) CreateDockerProxyRepo(r *apiv1.DockerProxyApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Post("service/rest/v1/repositories/docker/proxy")
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}
