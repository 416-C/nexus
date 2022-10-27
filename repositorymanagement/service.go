package repositorymanagement

import (
	"github.com/416-C/nexus-client-go/repositorymanagement/cocoapods"
	"github.com/416-C/nexus-client-go/repositorymanagement/docker"
	"github.com/416-C/nexus-client-go/repositorymanagement/gitlfs"
	"github.com/416-C/nexus-client-go/repositorymanagement/golang"
	"github.com/416-C/nexus-client-go/repositorymanagement/helm"
	"github.com/416-C/nexus-client-go/repositorymanagement/maven"
	"github.com/416-C/nexus-client-go/repositorymanagement/npm"
	"github.com/416-C/nexus-client-go/repositorymanagement/pypi"
	"github.com/416-C/nexus-client-go/repositorymanagement/raw"
	"github.com/416-C/nexus-client-go/repositorymanagement/repositories"

	"github.com/go-resty/resty/v2"
)

type RepositoryManagement struct {
	*docker.Docker
	*gitlfs.GitLFS
	*golang.Golang
	*helm.Helm
	*maven.Maven
	*npm.Npm
	*pypi.PyPI
	*raw.Raw
	*cocoapods.CocoaPods
	*repositories.Repository
}

func NewRepositoryManagementService(client *resty.Client) *RepositoryManagement {
	return &RepositoryManagement{
		Docker:     docker.NewDockerService(client),
		GitLFS:     gitlfs.NewGitLFSService(client),
		Golang:     golang.NewGolangService(client),
		Helm:       helm.NewHelmService(client),
		Maven:      maven.NewMavenService(client),
		Npm:        npm.NewNpmService(client),
		PyPI:       pypi.NewPypiService(client),
		Raw:        raw.NewRawService(client),
		CocoaPods:  cocoapods.NewCocoaPodsService(client),
		Repository: repositories.NewRepositoryService(client),
	}
}
