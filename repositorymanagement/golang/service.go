package golang

import (
	"errors"

	"github.com/416-C/nexus-client-go/repositorymanagement/common/model"
	"github.com/416-C/nexus-client-go/repositorymanagement/golang/apiv1"

	"github.com/go-resty/resty/v2"
)

type Golang struct {
	client *resty.Client
}

func NewGolangService(client *resty.Client) *Golang {
	return &Golang{
		client: client,
	}
}

func (s *Golang) GetGolangProxyRepo(repositoryName string) (*model.ProxyRepository, error) {
	response, err := s.client.NewRequest().SetResult(&model.ProxyRepository{}).Get("service/rest/v1/repositories/golang/proxy/" + repositoryName)
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*model.ProxyRepository), nil
}

func (s *Golang) UpdateGolangProxyRepo(repositoryName string, r *apiv1.GolangProxyApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Put("service/rest/v1/repositories/golang/proxy/" + repositoryName)
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}

func (s *Golang) CreateGolangProxyRepo(r *apiv1.GolangProxyApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Post("service/rest/v1/repositories/golang/proxy")
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}
