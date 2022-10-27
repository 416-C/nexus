package apiv1

import (
	"github.com/416-C/nexus-client-go/assets/apiv1"
)

type Attributes struct {
	// Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	ProprietaryComponents bool `json:"proprietaryComponents"`
}

type Component struct {
	Id         string        `json:"id"`
	Repository string        `json:"repository"`
	Format     string        `json:"format"`
	Group      string        `json:"group"`
	Name       string        `json:"name"`
	Version    string        `json:"version"`
	Assets     []apiv1.Asset `json:"assets"`
}

type ComponentList struct {
	Items             []Component `json:"items"`
	ContinuationToken string      `json:"continuationToken"`
}
