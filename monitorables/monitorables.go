package monitorables

import (
	"github.com/Vaelatern/monitoror/monitorables/azuredevops"
	"github.com/Vaelatern/monitoror/monitorables/github"
	"github.com/Vaelatern/monitoror/monitorables/gitlab"
	"github.com/Vaelatern/monitoror/monitorables/http"
	"github.com/Vaelatern/monitoror/monitorables/jenkins"
	"github.com/Vaelatern/monitoror/monitorables/ping"
	"github.com/Vaelatern/monitoror/monitorables/pingdom"
	"github.com/Vaelatern/monitoror/monitorables/port"
	"github.com/Vaelatern/monitoror/monitorables/travisci"
	"github.com/Vaelatern/monitoror/store"
)

func RegisterMonitorables(s *store.Store) {
	// ------------ AZURE DEVOPS ------------
	s.Registry.RegisterMonitorable(azuredevops.NewMonitorable(s))
	// ------------ GITHUB ------------
	s.Registry.RegisterMonitorable(github.NewMonitorable(s))
	// ------------ GITLAB ------------
	s.Registry.RegisterMonitorable(gitlab.NewMonitorable(s))
	// ------------ HTTP ------------
	s.Registry.RegisterMonitorable(http.NewMonitorable(s))
	// ------------ JENKINS ------------
	s.Registry.RegisterMonitorable(jenkins.NewMonitorable(s))
	// ------------ PING ------------
	s.Registry.RegisterMonitorable(ping.NewMonitorable(s))
	// ------------ PINGDOM ------------
	s.Registry.RegisterMonitorable(pingdom.NewMonitorable(s))
	// ------------ PORT ------------
	s.Registry.RegisterMonitorable(port.NewMonitorable(s))
	// ------------ TRAVIS CI ------------
	s.Registry.RegisterMonitorable(travisci.NewMonitorable(s))
}
