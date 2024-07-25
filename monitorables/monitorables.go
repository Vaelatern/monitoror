package monitorables

import (
	"github.com/Vaelatern/monitoror/monitorables/github"
	"github.com/Vaelatern/monitoror/monitorables/http"
	"github.com/Vaelatern/monitoror/monitorables/jenkins"
	"github.com/Vaelatern/monitoror/monitorables/ping"
	"github.com/Vaelatern/monitoror/monitorables/port"
	"github.com/Vaelatern/monitoror/store"
)

func RegisterMonitorables(s *store.Store) {
	// ------------ GITHUB ------------
	s.Registry.RegisterMonitorable(github.NewMonitorable(s))
	// ------------ HTTP ------------
	s.Registry.RegisterMonitorable(http.NewMonitorable(s))
	// ------------ JENKINS ------------
	s.Registry.RegisterMonitorable(jenkins.NewMonitorable(s))
	// ------------ PING ------------
	s.Registry.RegisterMonitorable(ping.NewMonitorable(s))
	// ------------ PORT ------------
	s.Registry.RegisterMonitorable(port.NewMonitorable(s))
}
