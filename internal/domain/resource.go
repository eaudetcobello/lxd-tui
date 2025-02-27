package domain

type ResourceType string

const (
	TypeContainer = "container"
	TypeProject   = "project"
)

type Resource interface {
	GetName() string
	GetType() ResourceType
}
