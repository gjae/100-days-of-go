package consul

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	consul "github.com/hashicorp/consul/api"
	"moviesexample.com/pkg/discovery"
)

// Registry defines a consul-based service registry
type Registry struct {
	client *consul.Client
}

// NewRegistry creates a new Consul-based service registry
// registry instance
func NewRegistry(addr string) (*Registry, error) {
	config := consul.DefaultConfig()
	config.Address = addr
	client, err := consul.NewClient(config)

	if err != nil {
		return nil, err
	}

	return &Registry{client: client}, nil
}

// Register creates a service record in the registry.
func (r *Registry) Register(ctx context.Context, instanceID discovery.InstanceID, serviceName string, hostPort string) error {
	parts := strings.Split(hostPort, ":")
	if len(parts) != 2 {
		return errors.New("hostPort must be in a form of <host>:<port>, example: localhost:8081")
	}
	port, err := strconv.Atoi(parts[1])
	if err != nil {
		return err
	}
	return r.client.Agent().ServiceRegister(&consul.AgentServiceRegistration{
		Address: parts[0],
		ID:      string(instanceID),
		Name:    serviceName,
		Port:    port,
		Check:   &consul.AgentServiceCheck{CheckID: string(instanceID), TTL: "5s"},
	})
}

// Deregister removes a service record from the
// registry.
func (r *Registry) Deregister(ctx context.Context, instanceID discovery.InstanceID, _ string) error {
	return r.client.Agent().ServiceDeregister(string(instanceID))
}

// ServiceAddresses returns the list of addresses of
// active instances of the given service.
func (r *Registry) ServiceAddresses(ctx context.Context, serviceName string) (discovery.ServiceAddresses, error) {
	entries, _, err := r.client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return nil, err
	} else if len(entries) == 0 {
		return nil, discovery.ErrNotFound
	}
	res := make(discovery.ServiceAddresses, 0)
	for _, e := range entries {
		res = append(res, fmt.Sprintf("%s:%d", e.Service.Address, e.Service.Port))
	}
	return res, nil
}

// ReportHealthyState is a push mechanism for
// reporting healthy state to the registry.
func (r *Registry) ReportHealthyState(instanceID discovery.InstanceID, _ string) error {
	return r.client.Agent().PassTTL(string(instanceID), "")
}
