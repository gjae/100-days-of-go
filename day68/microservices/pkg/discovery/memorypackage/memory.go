package memory

import (
	"context"
	"errors"
	"sync"
	"time"

	"moviesexample.com/pkg/discovery"
)

type serviceName string
type instanceID string

// Registry defines an in-memory service registry.
type Registry struct {
	sync.RWMutex
	serviceAddrs map[serviceName]map[instanceID]*serviceInstance
}
type serviceInstance struct {
	hostPort   string
	lastActive time.Time
}

// NewRegistry creates a new in-memory service
// registry instance
func NewRegistry() *Registry {
	return &Registry{serviceAddrs: map[serviceName]map[instanceID]*serviceInstance{}}
}

// Register creates a service record in the registry.
func (r *Registry) Register(ctx context.Context, iID discovery.InstanceID, sName string, hostPort string) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.serviceAddrs[serviceName(sName)]; !ok {
		r.serviceAddrs[serviceName(sName)] = map[string]*serviceInstance{}
	}
	r.serviceAddrs[serviceName(sName)][instanceID(iID)] = &serviceInstance{hostPort: hostPort, lastActive: time.Now()}
	return nil
}

// Deregister removes a service record from the
// registry.
func (r *Registry) Deregister(ctx context.Context, iID discovery.InstanceID, sName string) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.serviceAddrs[serviceName(sName)]; !ok {
		return nil
	}
	delete(r.serviceAddrs[serviceName(sName)], instanceID(iID))
	return nil
}

// ReportHealthyState is a push mechanism for
// reporting healthy state to the registry.
func (r *Registry) ReportHealthyState(iID discovery.InstanceID, sName string) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.serviceAddrs[serviceName(sName)]; !ok {
		return errors.New("service is not registered yet")
	}
	if _, ok := r.serviceAddrs[serviceName(sName)][instanceID(iID)]; !ok {
		return errors.New("service instance is not registered yet")
	}
	r.serviceAddrs[serviceName(sName)][instanceID(iID)].lastActive = time.Now()
	return nil
}

// ServiceAddresses returns the list of addresses of
// active instances of the given service.
func (r *Registry) ServiceAddresses(ctx context.Context, sName string) (discovery.ServiceAddresses, error) {
	r.RLock()
	defer r.RUnlock()
	if len(r.serviceAddrs[serviceName(sName)]) == 0 {
		return nil, discovery.ErrNotFound
	}
	var res = make(discovery.ServiceAddresses, 0)
	for _, i := range r.serviceAddrs[serviceName(sName)] {
		if i.lastActive.Before(time.Now().Add(-5 * time.Second)) {
			continue
		}
		res = append(res, i.hostPort)
	}
	return res, nil
}
