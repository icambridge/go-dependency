package dependency

import (
	"github.com/mcuadros/go-version"
)

type Solver struct {
	Packages map[string]map[string]Dependency
}

func (s Solver) Solve(root Dependency) map[string]string {

	required := map[string]string{}
	// found :=
	for packageName, rule := range root.Requires {
		cg := version.NewConstrainGroupFromString(rule)
		versions := s.Packages[packageName]
		for versionNum, _ := range versions {
			if cg.Match(versionNum) {
				required[packageName] = versionNum
				break
			}
		}
	}

	return required
}
