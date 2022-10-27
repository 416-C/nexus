package repositories

import (
	"errors"

	"github.com/416-C/nexus-client-go/repositorymanagement/repositories/apiv1"

	"github.com/go-resty/resty/v2"
)

type Repository struct {
	client *resty.Client
}

func NewRepositoryService(client *resty.Client) *Repository {
	return &Repository{
		client: client,
	}
}

func (s *Repository) InvalidateCache(repositoryName string) error {
	response, err := s.client.NewRequest().Post("service/rest/v1/repositories/" + repositoryName + "/invalidate-cache")
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}

func (s *Repository) RebuildIndex(repositoryName string) error {
	response, err := s.client.NewRequest().Post("service/rest/v1/repositories/" + repositoryName + "/rebuild-index")
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}

func (s *Repository) GetRepository(format, type_, repositoryName string) (*apiv1.Repository, error) {
	response, err := s.client.NewRequest().SetResult(&apiv1.Repository{}).Get("service/rest/v1/repositories/" + format + "/" + type_ + "/" + repositoryName)
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*apiv1.Repository), nil
}

func (s *Repository) DeleteRepository(repositoryName string) error {
	response, err := s.client.NewRequest().Delete("service/rest/v1/repositories/" + repositoryName)
	if err != nil {
		return err
	}
	if response.IsError() && response.StatusCode() != 404 {
		return errors.New(response.String())
	}

	return nil
}

func (s *Repository) ListRepository() ([]*apiv1.Repository, error) {
	response, err := s.client.NewRequest().SetResult([]*apiv1.Repository{}).Get("service/rest/v1/repositories")
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().([]*apiv1.Repository), nil
}

func (s *Repository) ListRepositorySettings() ([]*apiv1.RepositorySettings, error) {
	response, err := s.client.NewRequest().SetResult([]*apiv1.RepositorySettings{}).Get("service/rest/v1/repositorySettings")
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().([]*apiv1.RepositorySettings), nil
}
