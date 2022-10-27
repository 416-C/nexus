package cocoapods

import (
	"errors"

	"github.com/416-C/nexus-client-go/repositorymanagement/cocoapods/apiv1"
	"github.com/416-C/nexus-client-go/repositorymanagement/common/model"

	"github.com/go-resty/resty/v2"
)

type CocoaPods struct {
	client *resty.Client
}

func NewCocoaPodsService(client *resty.Client) *CocoaPods {
	return &CocoaPods{
		client: client,
	}
}

func (s *CocoaPods) GetCocoaPodsProxyRepo(repositoryName string) (*model.ProxyRepository, error) {
	response, err := s.client.NewRequest().SetResult(&model.ProxyRepository{}).Get("service/rest/v1/repositories/cocoapods/proxy/" + repositoryName)
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*model.ProxyRepository), nil
}

func (s *CocoaPods) UpdateCocoaPodsProxyRepo(repositoryName string, r *apiv1.CocoapodsProxyApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Put("service/rest/v1/repositories/cocoapods/proxy/" + repositoryName)
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}

func (s *CocoaPods) CreateCocoaPodsProxyRepo(r *apiv1.CocoapodsProxyApiRequest) error {
	response, err := s.client.NewRequest().SetBody(r).Post("service/rest/v1/repositories/cocoapods/proxy/")
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}
