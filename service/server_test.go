package service

import (
	"testing"

	"github.com/Vaelatern/monitoror/cli/debug"
	"github.com/Vaelatern/monitoror/config"
	"github.com/Vaelatern/monitoror/registry"
	"github.com/Vaelatern/monitoror/store"

	"github.com/GeertJohan/go.rice/embedded"
	"github.com/stretchr/testify/assert"
)

// /!\ this is an integration test /!\
// Note : It may be necessary to separate them from unit tests

func TestInit_Dev(t *testing.T) {
	s := &store.Store{
		CoreConfig: &config.CoreConfig{DisableUI: true},
		Registry:   registry.NewRegistry(),
	}
	debug.Enable()

	assert.NotPanics(t, func() {
		Init(s)
	})
}

func TestInit_Prod_WithoutRicebox(t *testing.T) {
	s := &store.Store{
		CoreConfig: &config.CoreConfig{DisableUI: false},
		Registry:   registry.NewRegistry(),
	}

	delete(embedded.EmbeddedBoxes, "../ui/dist")
	assert.Panics(t, func() {
		Init(s)
	})
}

func TestInit_Prod_WithRicebox(t *testing.T) {
	s := &store.Store{
		CoreConfig: &config.CoreConfig{DisableUI: false},
		Registry:   registry.NewRegistry(),
	}

	delete(embedded.EmbeddedBoxes, "../ui/dist")
	embedded.RegisterEmbeddedBox("../ui/dist", &embedded.EmbeddedBox{
		Name: "../ui/dist",
	})
	defer delete(embedded.EmbeddedBoxes, "../ui/dist")
	assert.NotPanics(t, func() {
		Init(s)
	})
}
