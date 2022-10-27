package gitlfs

import (
	"errors"

	"github.com/416-C/nexus-client-go/repositorymanagement/common/model"
	"github.com/416-C/nexus-client-go/repositorymanagement/gitlfs/apiv1"

	"github.com/go-resty/resty/v2"
)

type GitLFS struct {
	client *resty.Client
}

func NewGitLFSService(client *resty.Client) *GitLFS {
	return &GitLFS{
		client: client,
	}
}

func (s *GitLFS) GetGitLFSLocalRepo(repositoryName string) (*model.LocalRepository, error) {
	response, err := s.client.NewRequest().SetResult(&model.LocalRepository{}).Get("service/rest/v1/repositories/gitlfs/hosted/" + repositoryName)
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*model.LocalRepository), nil
}

func (s *GitLFS) UpdateGitLFSLocalRepo(repositoryName string, r *apiv1.GitLFSLocalApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Put("service/rest/v1/repositories/gitlfs/hosted/" + repositoryName)
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}

func (s *GitLFS) CreateGitLFSLocalRepo(r *apiv1.GitLFSLocalApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Post("service/rest/v1/repositories/gitlfs/hosted")
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}
