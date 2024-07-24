package monitorables

import (
	"testing"

	"github.com/Vaelatern/monitoror/internal/pkg/monitorable/test"
	coreModels "github.com/Vaelatern/monitoror/models"
	azureDevOpsApi "github.com/Vaelatern/monitoror/monitorables/azuredevops/api"
	githubApi "github.com/Vaelatern/monitoror/monitorables/github/api"
	gitlabApi "github.com/Vaelatern/monitoror/monitorables/gitlab/api"
	httpApi "github.com/Vaelatern/monitoror/monitorables/http/api"
	jenkinsApi "github.com/Vaelatern/monitoror/monitorables/jenkins/api"
	pingApi "github.com/Vaelatern/monitoror/monitorables/ping/api"
	pingdomApi "github.com/Vaelatern/monitoror/monitorables/pingdom/api"
	portApi "github.com/Vaelatern/monitoror/monitorables/port/api"
	travisCIApi "github.com/Vaelatern/monitoror/monitorables/travisci/api"
	"github.com/Vaelatern/monitoror/registry"

	"github.com/stretchr/testify/assert"
)

func TestManager_RegisterMonitorables(t *testing.T) {
	// init Store
	store, _ := test.InitMockAndStore()
	store.Registry = registry.NewRegistry()

	RegisterMonitorables(store)

	mr := store.Registry.(*registry.MetadataRegistry)

	// ------------ AZURE DEVOPS ------------
	assert.NotNil(t, mr.TileMetadata[azureDevOpsApi.AzureDevOpsBuildTileType])
	assert.NotNil(t, mr.TileMetadata[azureDevOpsApi.AzureDevOpsReleaseTileType])
	// ------------ GITHUB ------------
	assert.NotNil(t, mr.TileMetadata[githubApi.GithubCountTileType])
	assert.NotNil(t, mr.TileMetadata[githubApi.GithubChecksTileType])
	assert.NotNil(t, mr.TileMetadata[githubApi.GithubPullRequestTileType])
	assert.NotNil(t, mr.GeneratorMetadata[coreModels.NewGeneratorTileType(githubApi.GithubPullRequestTileType)])
	// ------------ GITLAB ------------
	assert.NotNil(t, mr.TileMetadata[gitlabApi.GitlabPipelineTileType])
	assert.NotNil(t, mr.TileMetadata[gitlabApi.GitlabMergeRequestTileType])
	assert.NotNil(t, mr.GeneratorMetadata[coreModels.NewGeneratorTileType(gitlabApi.GitlabMergeRequestTileType)])
	// ------------ HTTP ------------
	assert.NotNil(t, mr.TileMetadata[httpApi.HTTPStatusTileType])
	assert.NotNil(t, mr.TileMetadata[httpApi.HTTPRawTileType])
	assert.NotNil(t, mr.TileMetadata[httpApi.HTTPFormattedTileType])
	// ------------ JENKINS ------------
	assert.NotNil(t, mr.TileMetadata[jenkinsApi.JenkinsBuildTileType])
	assert.NotNil(t, mr.GeneratorMetadata[coreModels.NewGeneratorTileType(jenkinsApi.JenkinsBuildTileType)])
	// ------------ PING ------------
	assert.NotNil(t, mr.TileMetadata[pingApi.PingTileType])
	// ------------ PINGDOM ------------
	assert.NotNil(t, mr.TileMetadata[pingdomApi.PingdomCheckTileType])
	assert.NotNil(t, mr.TileMetadata[pingdomApi.PingdomTransactionCheckTileType])
	assert.NotNil(t, mr.GeneratorMetadata[coreModels.NewGeneratorTileType(pingdomApi.PingdomCheckTileType)])
	assert.NotNil(t, mr.GeneratorMetadata[coreModels.NewGeneratorTileType(pingdomApi.PingdomTransactionCheckTileType)])
	// ------------ PORT ------------
	assert.NotNil(t, mr.TileMetadata[portApi.PortTileType])
	// ------------ TRAVIS CI ------------
	assert.NotNil(t, mr.TileMetadata[travisCIApi.TravisCIBuildTileType])
}
