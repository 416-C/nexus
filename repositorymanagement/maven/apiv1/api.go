package apiv1

import (
	"github.com/416-C/nexus-client-go/components/apiv1"
	"github.com/416-C/nexus-client-go/repositorymanagement/common/attrs"
)

type Attributes struct {
	// What type of artifacts does this repository store?
	VersionPolicy string `json:"versionPolicy,omitempty"`
	// Validate that all paths are maven artifact or metadata paths
	LayoutPolicy string `json:"layoutPolicy,omitempty"`
	// Content Disposition
	ContentDisposition string `json:"contentDisposition,omitempty"`
}

type MavenLocalApiRequest struct {
	// A unique identifier for this repository
	Name string `json:"name"`
	// Whether this repository accepts incoming requests
	Online    bool                 `json:"online"`
	Storage   *attrs.LocalStorage  `json:"storage"`
	Cleanup   *attrs.CleanupPolicy `json:"cleanup,omitempty"`
	Component *apiv1.Attributes    `json:"component,omitempty"`
	Maven     *Attributes          `json:"maven"`
}

type MavenProxyApiRequest struct {
	// A unique identifier for this repository
	Name string `json:"name"`
	// Whether this repository accepts incoming requests
	Online        bool                                `json:"online"`
	Storage       *attrs.Storage                      `json:"storage"`
	Cleanup       *attrs.CleanupPolicy                `json:"cleanup,omitempty"`
	Proxy         *attrs.Proxy                        `json:"proxy"`
	NegativeCache *attrs.NegativeCache                `json:"negativeCache"`
	HttpClient    *attrs.HttpClientWithPreemptiveAuth `json:"httpClient"`
	RoutingRule   string                              `json:"routingRule,omitempty"`
	Replication   *attrs.Replication                  `json:"replication,omitempty"`
	Maven         *Attributes                         `json:"maven"`
}
