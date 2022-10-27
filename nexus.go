package nexus

import (
	"time"

	"github.com/416-C/nexus-client-go/assets"
	"github.com/416-C/nexus-client-go/components"
	"github.com/416-C/nexus-client-go/repositorymanagement"
	"github.com/416-C/nexus-client-go/search"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	Assets               *assets.Assets
	Components           *components.Components
	RepositoryManagement *repositorymanagement.RepositoryManagement
	Search               *search.Search
}

func NewClient(baseURL, username, password string, opts ...Option) *Client {
	o := &Options{
		BaseURL:          baseURL,
		Username:         username,
		Password:         password,
		Timeout:          5 * time.Second,
		RetryCount:       3,
		RetryWaitTime:    5 * time.Second,
		RetryMaxWaitTime: 10 * time.Second,
	}

	for _, opt := range opts {
		opt(o)
	}

	client := resty.New().
		SetBaseURL(o.BaseURL).
		SetBasicAuth(o.Username, o.Password).
		SetTimeout(o.Timeout).
		SetRetryCount(o.RetryCount).
		SetRetryWaitTime(o.RetryWaitTime).
		SetRetryMaxWaitTime(o.RetryMaxWaitTime)

	return &Client{
		Assets:               assets.NewAssetsService(client),
		Components:           components.NewComponentsService(client),
		RepositoryManagement: repositorymanagement.NewRepositoryManagementService(client),
		Search:               search.NewSearchService(client),
	}
}
