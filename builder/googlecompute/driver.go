package googlecompute

// Driver is the interface that has to be implemented to communicate
// with GCE. The Driver interface exists mostly to allow a mock implementation
// to be used to test the steps.
type Driver interface {
	// DeleteInstance deletes the given instance.
	DeleteInstance(zone, name string) (<-chan error, error)

	// GetNatIP gets the NAT IP address for the instance.
	GetNatIP(zone, name string) (string, error)

	// RunInstance takes the given config and launches an instance.
	RunInstance(*InstanceConfig) (<-chan error, error)

	// WaitForInstance waits for an instance to reach the given state.
	WaitForInstance(state, zone, name string) <-chan error
}

type InstanceConfig struct {
	Description string
	Image       string
	MachineType string
	Metadata    map[string]string
	Name        string
	Network     string
	Tags        []string
	Zone        string
}
