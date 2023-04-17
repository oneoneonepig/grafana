// This file is autogenerated. DO NOT EDIT.
//
// Generated by pkg/framework/coremodel/gen.go
//
// Run `make gen-cue` from repository root to regenerate.

package registry

import (
	"fmt"

	"github.com/grafana/grafana/pkg/coremodel/dashboard"
	"github.com/grafana/grafana/pkg/coremodel/pluginmeta"
	"github.com/grafana/grafana/pkg/framework/coremodel"
	"github.com/grafana/thema"
)

// Base is a registry of coremodel.Interface. It provides two modes for accessing
// coremodels: individually via literal named methods, or as a slice returned from All().
//
// Prefer the individual named methods for use cases where the particular coremodel(s) that
// are needed are known to the caller. For example, a dashboard linter can know that it
// specifically wants the dashboard coremodel.
//
// Prefer All() when performing operations generically across all coremodels. For example,
// a validation HTTP middleware for any coremodel-schematized object type.
type Base struct {
	all        []coremodel.Interface
	dashboard  *dashboard.Coremodel
	pluginmeta *pluginmeta.Coremodel
}

// type guards
var (
	_ coremodel.Interface = &dashboard.Coremodel{}
	_ coremodel.Interface = &pluginmeta.Coremodel{}
)

// Dashboard returns the dashboard coremodel. The return value is guaranteed to
// implement coremodel.Interface.
func (b *Base) Dashboard() *dashboard.Coremodel {
	return b.dashboard
}

// Pluginmeta returns the pluginmeta coremodel. The return value is guaranteed to
// implement coremodel.Interface.
func (b *Base) Pluginmeta() *pluginmeta.Coremodel {
	return b.pluginmeta
}

func doProvideBase(lib thema.Library) *Base {
	var err error
	reg := &Base{}

	reg.dashboard, err = dashboard.New(lib)
	if err != nil {
		panic(fmt.Sprintf("error while initializing dashboard coremodel: %s", err))
	}
	reg.all = append(reg.all, reg.dashboard)

	reg.pluginmeta, err = pluginmeta.New(lib)
	if err != nil {
		panic(fmt.Sprintf("error while initializing pluginmeta coremodel: %s", err))
	}
	reg.all = append(reg.all, reg.pluginmeta)

	return reg
}