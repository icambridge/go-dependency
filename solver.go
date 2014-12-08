package dependency

import (
	"errors"
	"fmt"
	"github.com/deckarep/golang-set"
	"github.com/mcuadros/go-version"
)

type Solver struct {
	Packages map[string]map[string]Dependency
	Required map[string]string
	Found    map[string]string
	Rules    map[string]mapset.Set
}

func (s Solver) Solve(root Dependency) (map[string]string, error) {

	rules := GetRules([]Dependency{root})
	err := s.Inner(rules)

	return s.Found, err
}

func (s Solver) Inner(rules map[string]mapset.Set) error {
	if len(rules) == 0 {
		return nil
	}
	required := []Dependency{}

	for packageName, packageRules := range rules {

		_, ok := s.Rules[packageName]

		if !ok {
			s.Rules[packageName] = mapset.NewSet()
		}

		s.Rules[packageName] = s.Rules[packageName].Union(packageRules)

		expectedTotal := s.Rules[packageName].Cardinality()
		found := false

		versionSet := GetVersionNumbers(s.Packages[packageName])
		versions := PrepVersionNumbers(versionSet)

		for _, versionNum := range versions {
			passes := 0

			for _, packageRuleI := range s.Rules[packageName].ToSlice() {
				// todo move to one creation
				packageRule := fmt.Sprintf("%s", packageRuleI)
				cg := version.NewConstrainGroupFromString(packageRule)
				if cg.Match(versionNum) {
					passes++
				}
			}
			if passes == expectedTotal {
				// TODO log rules
				found = true
				foundVersion, ok := s.Found[packageName]
				if !ok || foundVersion != versionNum {
					s.Found[packageName] = versionNum
					required = append(required, s.Packages[packageName][versionNum])
				}
				break
			}
		}
		if !found {
			return errors.New(fmt.Sprintf("Couldn't find a package for %s", packageName))
		}
	}

	newRules := GetRules(required)
	return s.Inner(newRules)
}

func GetRules(dependency []Dependency) map[string]mapset.Set {

	find := map[string]mapset.Set{}
	for _, root := range dependency {
		for requiredName, requiredRule := range root.Requires {

			_, ok := find[requiredName]

			if !ok {
				find[requiredName] = mapset.NewSet()
			}
			find[requiredName].Add(requiredRule)
		}
	}
	return find
}

func NewSolver(packages map[string]map[string]Dependency) Solver {

	return Solver{packages, map[string]string{}, map[string]string{}, map[string]mapset.Set{}}
}
