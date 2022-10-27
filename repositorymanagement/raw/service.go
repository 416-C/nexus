package raw

import (
	"errors"

	"github.com/416-C/nexus-client-go/repositorymanagement/common/model"
	"github.com/416-C/nexus-client-go/repositorymanagement/raw/apiv1"

	"github.com/go-resty/resty/v2"
)

type Raw struct {
	client *resty.Client
}

func NewRawService(client *resty.Client) *Raw {
	return &Raw{
		client: client,
	}
}

func (s *Raw) GetRawLocalRepo(repositoryName string) (*apiv1.RawLocalRepository, error) {
	response, err := s.client.NewRequest().SetResult(&apiv1.RawLocalRepository{}).Get("service/rest/v1/repositories/raw/hosted/" + repositoryName)
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*apiv1.RawLocalRepository), nil
}

func (s *Raw) UpdateRawLocalRepo(repositoryName string, r *apiv1.RawLocalApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Put("service/rest/v1/repositories/raw/hosted/" + repositoryName)
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}

func (s *Raw) CreateRawLocalRepo(r *apiv1.RawLocalApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Post("service/rest/v1/repositories/raw/hosted")
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}

func (s *Raw) GetRawProxyRepo(repositoryName string) (*model.ProxyRepository, error) {
	response, err := s.client.NewRequest().SetResult(&model.ProxyRepository{}).Get("service/rest/v1/repositories/raw/proxy/" + repositoryName)
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*model.ProxyRepository), nil
}

func (s *Raw) UpdateRawProxyRepo(repositoryName string, r *apiv1.RawProxyApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Put("service/rest/v1/repositories/raw/proxy/" + repositoryName)
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}

func (s *Raw) CreateRawProxyRepo(r *apiv1.RawProxyApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Post("service/rest/v1/repositories/raw/proxy")
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}
