package dependency

import (
	"testing"
)

func Test_Gets_Correct(t *testing.T) {

	mink := map[string]Dependency{
		"1.6.0": Dependency{
			Name:    "behat/mink",
			Version: "1.6.0",
		},
		"1.5.1": Dependency{
			Name:    "behat/mink",
			Version: "1.5.1",
		},
	}

	minkExt := map[string]Dependency{
		"1.4.1": Dependency{
			Name:    "behat/mink-ext",
			Version: "1.2.5",
			Requires: map[string]string{
				"behat/mink": "~1.5",
			},
		},
	}

	minkSymfonyBrowser := map[string]Dependency{
		"1.2.0": Dependency{
			Name:    "behat/mink-symfony",
			Version: "1.2.5",
			Requires: map[string]string{
				"behat/mink": "~1.6",
			},
		},
		"1.1.0": Dependency{
			Name:    "behat/mink-symfony",
			Version: "1.2.5",
			Requires: map[string]string{
				"behat/mink": "~1.5",
			},
		},
	}

	packages := map[string]map[string]Dependency{
		"behat/mink":         mink,
		"behat/mink-ext":     minkExt,
		"behat/mink-symfony": minkSymfonyBrowser,
	}

	root := Dependency{
		Name: "App",
		Requires: map[string]string{
			"behat/mink-symfony": "~1.1",
			"behat/mink-ext":     "~1.1",
		},
	}

	s := NewSolver(packages)
	required := s.Solve(root)

	if minkExtV := "1.4.1"; minkExtV != required["behat/mink-ext"] {
		t.Errorf("Expected to require %v but got %v", minkExtV, required["behat/mink-ext"])
		return
	}

	if minkSymfonyV := "1.2.0"; minkSymfonyV != required["behat/mink-symfony"] {
		t.Errorf("Expected to require %v but got %v", minkSymfonyV, required["behat/mink-symfony"])
		return
	}
}


func Test_Gets_Correct_Including_Second_Layer(t *testing.T) {

	mink := map[string]Dependency{
		"1.6.0": Dependency{
			Name:    "behat/mink",
			Version: "1.6.0",
		},
		"1.5.1": Dependency{
			Name:    "behat/mink",
			Version: "1.5.1",
		},
	}

	minkExt := map[string]Dependency{
		"1.4.1": Dependency{
			Name:    "behat/mink-ext",
			Version: "1.2.5",
			Requires: map[string]string{
				"behat/mink": "~1.5",
			},
		},
	}

	minkSymfonyBrowser := map[string]Dependency{
		"1.2.0": Dependency{
			Name:    "behat/mink-symfony",
			Version: "1.2.5",
			Requires: map[string]string{
				"behat/mink": "~1.6",
			},
		},
		"1.1.0": Dependency{
			Name:    "behat/mink-symfony",
			Version: "1.2.5",
			Requires: map[string]string{
				"behat/mink": "~1.5",
			},
		},
	}

	packages := map[string]map[string]Dependency{
		"behat/mink":         mink,
		"behat/mink-ext":     minkExt,
		"behat/mink-symfony": minkSymfonyBrowser,
	}

	root := Dependency{
		Name: "App",
		Requires: map[string]string{
			"behat/mink-symfony": "~1.1",
			"behat/mink-ext":     "~1.1",
		},
	}

	s := NewSolver(packages)
	required := s.Solve(root)

	if minkExtV := "1.4.1"; minkExtV != required["behat/mink-ext"] {
		t.Errorf("Expected to require %v but got %v", minkExtV, required["behat/mink-ext"])
		return
	}

	if minkSymfonyV := "1.2.0"; minkSymfonyV != required["behat/mink-symfony"] {
		t.Errorf("Expected to require %v but got %v", minkSymfonyV, required["behat/mink-symfony"])
		return
	}
	if minkSymfonyV := "1.6.0"; minkSymfonyV != required["behat/mink"] {
		t.Errorf("Expected to require %v but got %v", minkSymfonyV, required["behat/mink"])
		return
	}
}

func Test_Gets_Correct_Without_Infinite_Loop(t *testing.T) {

	mink := map[string]Dependency{
		"1.6.0": Dependency{
			Name:    "behat/mink",
			Version: "1.6.0",
			Requires: map[string]string{
				"behat/mink-symfony": "~1.1",
			},
		},
		"1.5.1": Dependency{
			Name:    "behat/mink",
			Version: "1.5.1",
		},
	}

	minkExt := map[string]Dependency{
		"1.4.1": Dependency{
			Name:    "behat/mink-ext",
			Version: "1.2.5",
			Requires: map[string]string{
				"behat/mink": "~1.5",
			},
		},
	}

	minkSymfonyBrowser := map[string]Dependency{
		"1.2.0": Dependency{
			Name:    "behat/mink-symfony",
			Version: "1.2.5",
			Requires: map[string]string{
				"behat/mink": "~1.6",
			},
		},
		"1.1.0": Dependency{
			Name:    "behat/mink-symfony",
			Version: "1.2.5",
			Requires: map[string]string{
				"behat/mink": "~1.5",
			},
		},
	}

	packages := map[string]map[string]Dependency{
		"behat/mink":         mink,
		"behat/mink-ext":     minkExt,
		"behat/mink-symfony": minkSymfonyBrowser,
	}

	root := Dependency{
		Name: "App",
		Requires: map[string]string{
			"behat/mink-symfony": "~1.1",
			"behat/mink-ext":     "~1.1",
		},
	}

	s := NewSolver(packages)
	required := s.Solve(root)

	if minkExtV := "1.4.1"; minkExtV != required["behat/mink-ext"] {
		t.Errorf("Expected to require %v but got %v", minkExtV, required["behat/mink-ext"])
		return
	}

	if minkSymfonyV := "1.2.0"; minkSymfonyV != required["behat/mink-symfony"] {
		t.Errorf("Expected to require %v but got %v", minkSymfonyV, required["behat/mink-symfony"])
		return
	}
	if minkSymfonyV := "1.6.0"; minkSymfonyV != required["behat/mink"] {
		t.Errorf("Expected to require %v but got %v", minkSymfonyV, required["behat/mink"])
		return
	}
}
func Test_GetRules_Returns_Rules(t *testing.T) {

	root := Dependency{
		Name: "App",
		Requires: map[string]string{
			"behat/mink-symfony": "~1.1",
			"behat/mink-ext":     "~1.1",
		},
	}

	rules := GetRules([]Dependency{root})
	foundCount := 0
	for packageName, packageRules := range rules {

		if packageName == "behat/mink-ext" && packageRules[0] == "~1.1" {
			foundCount++
		}

		if packageName == "behat/mink-symfony" && packageRules[0] == "~1.1" {
			foundCount++
		}

	}

	if expected := 2; foundCount != expected {
		t.Errorf("expected %v but got %v", expected, foundCount)
	}
}
