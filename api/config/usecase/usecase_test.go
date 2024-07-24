package usecase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"github.com/Vaelatern/monitoror/api/config"
	"github.com/Vaelatern/monitoror/api/config/models"
	"github.com/Vaelatern/monitoror/api/config/repository"
	"github.com/Vaelatern/monitoror/api/config/versions"
	coreConfig "github.com/Vaelatern/monitoror/config"
	coreModels "github.com/Vaelatern/monitoror/models"
	jenkinsApi "github.com/Vaelatern/monitoror/monitorables/jenkins/api"
	jenkinsModels "github.com/Vaelatern/monitoror/monitorables/jenkins/api/models"
	pingApi "github.com/Vaelatern/monitoror/monitorables/ping/api"
	pingModels "github.com/Vaelatern/monitoror/monitorables/ping/api/models"
	pingdomApi "github.com/Vaelatern/monitoror/monitorables/pingdom/api"
	pindomModels "github.com/Vaelatern/monitoror/monitorables/pingdom/api/models"
	portApi "github.com/Vaelatern/monitoror/monitorables/port/api"
	portModels "github.com/Vaelatern/monitoror/monitorables/port/api/models"
	"github.com/Vaelatern/monitoror/registry"
	"github.com/Vaelatern/monitoror/store"

	"github.com/jsdidierlaurent/echo-middleware/cache"
	"github.com/stretchr/testify/assert"
)

func initConfigUsecase(repository config.Repository) *configUsecase {
	s := &store.Store{
		CoreConfig: &coreConfig.CoreConfig{InitialMaxDelay: 1000},
		CacheStore: cache.NewGoCacheStore(time.Second, time.Second),
		Registry:   registry.NewRegistry(),
	}

	usecase := NewConfigUsecase(repository, s).(*configUsecase)

	usecase.registry.RegisterTile(pingApi.PingTileType, versions.MinimalVersion, []coreModels.VariantName{coreModels.DefaultVariantName}).
		Enable(coreModels.DefaultVariantName, &pingModels.PingParams{}, "/ping/default/ping")
	usecase.registry.RegisterTile(portApi.PortTileType, versions.MinimalVersion, []coreModels.VariantName{coreModels.DefaultVariantName}).
		Enable(coreModels.DefaultVariantName, &portModels.PortParams{}, "/port/default/port")
	usecase.registry.RegisterTile(jenkinsApi.JenkinsBuildTileType, versions.MinimalVersion, []coreModels.VariantName{coreModels.DefaultVariantName, "disabledVariant"}).
		Enable(coreModels.DefaultVariantName, &jenkinsModels.BuildParams{}, "/jenkins/default/build")
	usecase.registry.RegisterTile(pingdomApi.PingdomCheckTileType, versions.MinimalVersion, []coreModels.VariantName{coreModels.DefaultVariantName}).
		Enable(coreModels.DefaultVariantName, &pindomModels.CheckParams{}, "/pingdom/default/check")

	return usecase
}

func readConfig(input string) (configBag *models.ConfigBag, err error) {
	configBag = &models.ConfigBag{}
	reader := ioutil.NopCloser(strings.NewReader(input))
	configBag.Config, err = repository.ReadConfig(reader)
	return
}

func TestUsecase_Global_Success(t *testing.T) {
	rawConfig := fmt.Sprintf(`
{
	"version" : %q,
  "columns": 4,
  "tiles": [
		{ "type": "PORT", "label": "Monitoror", "params": {"hostname": "localhost", "port": 8080} }
  ]
}
`, versions.CurrentVersion)

	expectRawConfig := fmt.Sprintf(`{"config":{"version":%q,"columns":4,"tiles":[{"type":"PORT","label":"Monitoror","url":"/port/default/port?hostname=localhost\u0026port=8080","initialMaxDelay":1000}]}}`, versions.CurrentVersion)

	config, err := readConfig(rawConfig)
	if assert.NoError(t, err) {
		usecase := initConfigUsecase(nil)
		usecase.Verify(config)
		usecase.Hydrate(config)

		assert.Len(t, config.Errors, 0)

		marshal, err := json.Marshal(config)
		assert.NoError(t, err)
		assert.Equal(t, expectRawConfig, string(marshal))
	}
}
