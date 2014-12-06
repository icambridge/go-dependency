package dependency

type Dependency struct {
	Name     string
	Version  string
	Requires map[string]string
}
