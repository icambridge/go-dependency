package dependency

type Dependency struct {
	Name     string
	Version  string
	Requires map[string]string
}

func SolveDependencies(packages map[string]map[string]Dependency) map[string]string {

	required := map[string]string{}



    return required
}
