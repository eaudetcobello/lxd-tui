package interfaces

type Instance struct {
	Name   string
	Status string
}

type InstanceType string

const (
	InstanceTypeContainer InstanceType = "container"
	InstanceTypeVM        InstanceType = "vm"
	InstanceTypeAny       InstanceType = "any"
)

// InstanceDAO is the interface that wraps the basic instance methods.
// DAO allows to abstract the data layer from the business logic.
// For example, the UI layer can call the GetInstances method without knowing
// how the data is retrieved. This allows to easily switch between different
// data sources (e.g. LXD, Incus, etc.) without changing the UI code.
type InstanceDAO interface {
	// GetInstances returns a list of instances
	GetInstances(instanceType InstanceType, projectName string) ([]Instance, error)
	DeleteInstance(instanceName string, projectName string) error
	StopInstance(instanceName string) error
}
