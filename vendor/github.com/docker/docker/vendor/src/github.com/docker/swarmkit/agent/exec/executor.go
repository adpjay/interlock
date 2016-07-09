package exec

import (
	"github.com/docker/swarmkit/api"
	"golang.org/x/net/context"
)

// Executor provides controllers for tasks.
type Executor interface {
	// Describe returns the underlying node description.
	Describe(ctx context.Context) (*api.NodeDescription, error)

<<<<<<< HEAD
	// Configure uses the node object state to propogate node
=======
	// Configure uses the node object state to propagate node
>>>>>>> 12a5469... start on swarm services; move to glade
	// state to the underlying executor.
	Configure(ctx context.Context, node *api.Node) error

	// Controller provides a controller for the given task.
	Controller(t *api.Task) (Controller, error)

	// SetNetworkBootstrapKeys passes the symmetric keys from the
	// manager to the executor.
	SetNetworkBootstrapKeys([]*api.EncryptionKey) error
}