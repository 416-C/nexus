package apiv1

import (
	"github.com/416-C/nexus-client-go/components/apiv1"
	"github.com/416-C/nexus-client-go/repositorymanagement/common/attrs"
)

type Attributes struct {
	// Remove Non-Cataloged Versions
	RemoveNonCataloged bool `json:"removeNonCataloged"`
	// Remove Quarantined Versions
	RemoveQuarantined bool `json:"removeQuarantined"`
}

type NpmLocalApiRequest struct {
	// A unique identifier for this repository
	Name string `json:"name"`
	// Whether this repository accepts incoming requests
	Online    bool                 `json:"online"`
	Storage   *attrs.LocalStorage  `json:"storage"`
	Cleanup   *attrs.CleanupPolicy `json:"cleanup,omitempty"`
	Component *apiv1.Attributes    `json:"component,omitempty"`
}
