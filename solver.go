package dependency

type Solver struct {
	Packages map[string]map[string]Dependency
}

func (s Solver) Solve() map[string]string {

	required := map[string]string{}

	return required
}

func (s Solver) findUniqueRules(dependencies map[string]Dependency) map[string]map[string][]string {

	rules := map[string]map[string][]string{}

	for version, dependency := range dependencies {
		for name, rule := range dependency.Requires {
			_, nameExists := rules[name]

			if !nameExists {
				rules[name] = map[string][]string{}
			}

			_, ruleExists := rules[name][rule]

			if !ruleExists {
				rules[name][rule] = []string{}
			}

			rules[name][rule] = append(rules[name][rule], version)

		}
	}

	return rules
}

func (s Solver) mergeRules(rules map[string]map[string]map[string][]string) map[string]map[string][]string {

	mergedRules := map[string]map[string][]string{}

	for name, dRules := range rules {

		for dependencyName, pdRules := range dRules {

			_, nameExists := mergedRules[dependencyName]

			if !nameExists {
				mergedRules[dependencyName] = map[string][]string{}
			}
			for rule, _ := range pdRules {

				_, ruleExists := mergedRules[dependencyName][rule]

				if !ruleExists {
					mergedRules[dependencyName][rule] = []string{}
				}

				mergedRules[dependencyName][rule] = append(mergedRules[dependencyName][rule], name)
			}
		}
	}

	return mergedRules
}
