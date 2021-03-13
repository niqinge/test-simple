package service

import (
	"fmt"
	"github.com/niqinge/test-simple/pkg/service/runtime"
)

type Runtime interface {
	Del(ctx runtime.BaseContext) error
	Create(ctx runtime.BaseContext) error
	Update(ctx runtime.BaseContext) error
	Query(ctx runtime.BaseContext) error
	QueryPage(ctx runtime.PageContext) error
}

type Manager struct {
	invoker map[string]Runtime
}

func (m *Manager) getRuntime(schema string) Runtime {
	if rt := m.invoker[schema]; rt != nil {
		return rt
	}
	panic(fmt.Sprintf("#Manager Not found runtime, schema:%s", schema))
}

func (m *Manager) Create(ctx runtime.BaseContext) error {
	return m.getRuntime(ctx.Schema()).Create(ctx)
}

func (m *Manager) Delete(ctx runtime.BaseContext) error {
	return m.getRuntime(ctx.Schema()).Del(ctx)
}

func (m *Manager) Update(ctx runtime.BaseContext) error {
	return m.getRuntime(ctx.Schema()).Update(ctx)
}

func (m *Manager) Query(ctx runtime.BaseContext) error {
	return m.getRuntime(ctx.Schema()).Query(ctx)
}

func (m *Manager) QueryListPage(ctx runtime.PageContext) error {
	return m.getRuntime(ctx.Schema()).QueryPage(ctx)
}
