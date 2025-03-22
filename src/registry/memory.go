package registry

import (
	"fmt"
	"github.com/kercylan98/vivid-stateful-common/src/application"
	"github.com/kercylan98/vivid/src/vivid"
)

var _ application.Registry = (*MemoryRegistry)(nil)

func NewMemoryRegistry() *MemoryRegistry {
	return &MemoryRegistry{}
}

// MemoryRegistry 是一个内存注册中心，用于存储 ActorRef
type MemoryRegistry struct {
	services map[string]map[string]vivid.ActorRef
}

func (m *MemoryRegistry) Register(serviceType string, ref vivid.ActorRef) {
	if m.services == nil {
		m.services = make(map[string]map[string]vivid.ActorRef)
	}

	if _, ok := m.services[serviceType]; !ok {
		m.services[serviceType] = make(map[string]vivid.ActorRef)
	}

	m.services[serviceType][ref.Address()+ref.Path()] = ref
}

func (m *MemoryRegistry) Select(serviceType string) (vivid.ActorRef, error) {
	for _, ref := range m.services[serviceType] {
		return ref, nil
	}
	return nil, fmt.Errorf("%s %w", serviceType, ErrorNoServiceAvailable)
}

func (m *MemoryRegistry) SelectAll(serviceType string) []vivid.ActorRef {
	var refs []vivid.ActorRef
	for _, ref := range m.services[serviceType] {
		refs = append(refs, ref)
	}
	return refs
}
