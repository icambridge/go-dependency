package dependency

type Solver struct {
	Packages map[string]map[string]Dependency
}

func (s Solver) Solve() map[string]string {

	required := map[string]string{}

    return required
}
