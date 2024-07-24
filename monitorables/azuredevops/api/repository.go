//go:generate mockery -name Connection|Repository

package api

import (
	"github.com/Vaelatern/monitoror/monitorables/azuredevops/api/models"

	"github.com/microsoft/azure-devops-go-api/azuredevops/build"
	"github.com/microsoft/azure-devops-go-api/azuredevops/release"
)

type (
	Connection interface {
		GetBuildConnection() (build.Client, error)
		GetReleaseConnection() (release.Client, error)
	}

	Repository interface {
		GetBuild(project string, definition int, branch *string) (*models.Build, error)
		GetRelease(project string, definition int) (*models.Release, error)
	}
)
