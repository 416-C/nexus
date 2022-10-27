package assets

import (
	"errors"

	"github.com/416-C/nexus-client-go/assets/apiv1"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type Assets struct {
	client *resty.Client
}

func NewAssetsService(client *resty.Client) *Assets {
	return &Assets{
		client: client,
	}
}

func (s *Assets) GetAsset(assetID string) (*apiv1.Asset, error) {
	response, err := s.client.NewRequest().SetResult(&apiv1.Asset{}).Get("service/rest/v1/assets/" + assetID)
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*apiv1.Asset), nil
}

func (s *Assets) DeleteAsset(assetID string) error {
	response, err := s.client.NewRequest().Delete("service/rest/v1/assets/" + assetID)
	if err != nil {
		return err
	}
	if response.IsError() {
		return errors.New(response.String())
	}

	return nil
}

func (s *Assets) ListAssets(params *apiv1.ListAssetsParameters) (*apiv1.AssetList, error) {
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	response, err := s.client.NewRequest().SetResult(&apiv1.AssetList{}).SetQueryParamsFromValues(v).Get("service/rest/v1/assets")
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*apiv1.AssetList), nil
}
