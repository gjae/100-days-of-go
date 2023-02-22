package discovery

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// InstanceID defines a InstanceID
type InstanceID string

// ServiceAddresses defines a list of services hosts addresses
type ServiceAddresses []string

// Registry defines a service registry
type Registry interface {
	//Register creates a service instance record in the registry
	Register(ctx context.Context, instanceID InstanceID, serviceName string, hostPort string) error

	// Deregister removes a service instance record from
	// the registry
	Deregister(ctx context.Context, instanceID InstanceID, serviceName string) error

	// ServiceAddresses returns the list of addresses of
	// active instances of the given service
	ServiceAddresses(ctx context.Context, serviceID string) (ServiceAddresses, error)

	// ReportHealthyState is a push mechanism for reporteng
	// healthy state of the registry
	ReportHealthyState(instanceID InstanceID, serviceName string) error
}

// ErrNotFound is returned when no service addresses are found
var ErrNotFound = errors.New("No service addresses found")

// GenerateInstanceID generates a psudo-random service
// instance identifier, using a service name
// Suffixed by dash and random number
func GenerateInstanceID(serviceName string) InstanceID {
	newID := fmt.Sprintf("%s-%d", serviceName, rand.New(rand.NewSource(time.Now().UnixNano())).Int())
	return InstanceID(newID)
}
