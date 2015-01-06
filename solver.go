package dependency

import (
	"errors"
	"fmt"
	"github.com/deckarep/golang-set"
	"github.com/mcuadros/go-version"
	"strings"
)

type Solver struct {
	Packages        map[string]map[string]Dependency
	Required        map[string]string
	Found           map[string]string
	Replaced        mapset.Set
	Rules           map[string]mapset.Set
	RuleConstraints map[string]*version.ConstraintGroup
}

func (s Solver) Solve(root Dependency) (map[string]string, error) {

	rules := GetRules([]Dependency{root})
	err := s.Inner(rules)
	output := map[string]string{}

	for k, v := range s.Found {
		if !s.Replaced.Contains(k) {
			output[k] = v
		}
	}

	return output, err
}

func (s Solver) Inner(rules map[string]mapset.Set) error {
	if len(rules) == 0 {
		return nil
	}
	required := []Dependency{}

	for packageName, packageRules := range rules {

		if s.Replaced.Contains(packageName) {
			continue
		}

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

				packageRule := fmt.Sprintf("%s", packageRuleI)
				cg, found := s.RuleConstraints[packageRule]
				if !found {
					cg = version.NewConstrainGroupFromString(packageRule)
					s.RuleConstraints[packageRule] = cg
				}

				if cg.Match(versionNum) {
					passes++
				}
			}
			if passes == expectedTotal {
				found = true
				foundVersion, ok := s.Found[packageName]
				if !ok || foundVersion != versionNum {
					s.Found[packageName] = versionNum
					foundV := s.Packages[packageName][versionNum]
					for k, _ := range foundV.Replaces {
						s.Replaced.Add(k)
					}
					required = append(required, foundV)
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
		root.ReplaceSelfVersion()
		for requiredName, requiredRule := range root.Requires {

			requiredName = strings.ToLower(requiredName)
			_, ok := find[requiredName]

			if !ok {
				find[requiredName] = mapset.NewSet()
			}
			find[requiredName].Add(requiredRule)
		}
	}
	return find
}

func NewSolver(packages map[string]map[string]Dependency, replaces mapset.Set) Solver {

	return Solver{packages, map[string]string{}, map[string]string{}, replaces, map[string]mapset.Set{}, map[string]*version.ConstraintGroup{}}
}
