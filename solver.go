package dependency

import (
	"github.com/mcuadros/go-version"
)

type Solver struct {
	Packages map[string]map[string]Dependency
	Required map[string]string
}

func (s Solver) Solve(root Dependency) map[string]string {

	for packageName, rule := range root.Requires {
		cg := version.NewConstrainGroupFromString(rule)
		versionSet := GetVersionNumbers(s.Packages[packageName])
		versions := PrepVersionNumbers(versionSet)
		for _, versionNum := range versions {
			if cg.Match(versionNum) {
				s.Required[packageName] = versionNum
				break
			}
		}
	}

	return s.Required
}
