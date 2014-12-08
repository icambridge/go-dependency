package dependency

import (
	"github.com/mcuadros/go-version"
)

type Solver struct {
	Packages map[string]map[string]Dependency
	Required map[string]string
	Found    map[string]string
}

func (s Solver) Solve(root Dependency) map[string]string {

	rules := GetRules([]Dependency{root})
	s.Inner(rules)

	return s.Found
}

func (s Solver) Inner(rules map[string][]string) {
//	required := []Dependency{}
	for packageName, packageRules := range rules {
		expectedTotal := len(packageRules)

		versionSet := GetVersionNumbers(s.Packages[packageName])
		versions := PrepVersionNumbers(versionSet)
		for _, versionNum := range versions {
			passes := 0
			for _, packageRule := range packageRules {
				// todo move to one creation
				cg := version.NewConstrainGroupFromString(packageRule)
				if cg.Match(versionNum) {
					passes++
				}
			}
			if passes == expectedTotal {
				s.Found[packageName] = versionNum
				break
			}
		}
	}
}

func GetRules(dependency []Dependency) map[string][]string {

	find := map[string][]string{}
	for _, root := range dependency {
		for requiredName, requiredRule := range root.Requires {

			_, ok := find[requiredName]

			if !ok {
				find[requiredName] = []string{}
			}
			find[requiredName] = append(find[requiredName], requiredRule)
		}
	}
	return find
}

func NewSolver(packages map[string]map[string]Dependency) Solver {
	return Solver{packages, map[string]string{}, map[string]string{}}
}
