package apiv1

import (
	componentsapiv1 "github.com/416-C/nexus-client-go/components/apiv1"
	"github.com/416-C/nexus-client-go/repositorymanagement/common/attrs"
)

type PyPILocalApiRequest struct {
	// A unique identifier for this repository
	Name string `json:"name"`
	// Whether this repository accepts incoming requests
	Online    bool                        `json:"online"`
	Storage   *attrs.LocalStorage         `json:"storage"`
	Cleanup   *attrs.CleanupPolicy        `json:"cleanup,omitempty"`
	Component *componentsapiv1.Attributes `json:"component,omitempty"`
}

type PyPIProxyApiRequest struct {
	// A unique identifier for this repository
	Name string `json:"name"`
	// Whether this repository accepts incoming requests
	Online        bool                 `json:"online"`
	Storage       *attrs.Storage       `json:"storage"`
	Cleanup       *attrs.CleanupPolicy `json:"cleanup,omitempty"`
	Proxy         *attrs.Proxy         `json:"proxy"`
	NegativeCache *attrs.NegativeCache `json:"negativeCache"`
	HttpClient    *attrs.HttpClient    `json:"httpClient"`
	RoutingRule   string               `json:"routingRule,omitempty"`
	Replication   *attrs.Replication   `json:"replication,omitempty"`
}
