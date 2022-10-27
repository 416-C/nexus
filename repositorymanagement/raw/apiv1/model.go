package apiv1

import (
	componentsapiv1 "github.com/416-C/nexus-client-go/components/apiv1"
	"github.com/416-C/nexus-client-go/repositorymanagement/common/attrs"
)

type RawLocalRepository struct {
	// A unique identifier for this repository
	Name string `json:"name,omitempty"`
	// Whether this repository accepts incoming requests
	Online    bool                        `json:"online"`
	Storage   *attrs.LocalStorage         `json:"storage"`
	Cleanup   *attrs.CleanupPolicy        `json:"cleanup,omitempty"`
	Component *componentsapiv1.Attributes `json:"component,omitempty"`
}
