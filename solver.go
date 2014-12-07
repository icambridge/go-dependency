package dependency

import (
	"github.com/mcuadros/go-version"
)

type Solver struct {
	Packages map[string]map[string]Dependency
}

func (s Solver) Solve(root Dependency) map[string]string {

	required := map[string]string{}

	for packageName, rule := range root.Requires {
		cg := version.NewConstrainGroupFromString(rule)
		versionSet := GetVersionNumbers(s.Packages[packageName])
		versions := PrepVersionNumbers(versionSet)
		for _, versionNum := range versions {
			if cg.Match(versionNum) {
				required[packageName] = versionNum
				break
			}
		}
	}

	return required
}
