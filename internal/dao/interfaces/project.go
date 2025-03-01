package interfaces

type Project struct {
	Name string
}

type ProjectDAO interface {
	// GetProjects returns a list of projects
	GetProjects() ([]Project, error)
}
