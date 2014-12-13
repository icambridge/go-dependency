package dependency

import (
	"testing"
	"github.com/deckarep/golang-set"
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

	s := NewSolver(packages, mapset.NewSet())
	required, err := s.Solve(root)

	if err != nil {
		t.Errorf("%v", err)
	}

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

	s := NewSolver(packages, mapset.NewSet())
	required, err := s.Solve(root)

	if err != nil {
		t.Errorf("%v", err)
	}

	if minkExtV := "1.4.1"; minkExtV != required["behat/mink-ext"] {
		t.Errorf("Expected to require %v but got %v", minkExtV, required["behat/mink-ext"])
		return
	}

	if minkSymfonyV := "1.2.0"; minkSymfonyV != required["behat/mink-symfony"] {
		t.Errorf("Expected to require %v but got %v", minkSymfonyV, required["behat/mink-symfony"])
		return
	}
	if minkV := "1.6.0"; minkV != required["behat/mink"] {
		t.Errorf("Expected to require %v but got %v", minkV, required["behat/mink"])
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

	s := NewSolver(packages, mapset.NewSet())
	required, err := s.Solve(root)

	if err != nil {
		t.Errorf("%v", err)
	}

	if minkExtV := "1.4.1"; minkExtV != required["behat/mink-ext"] {
		t.Errorf("Expected to require %v but got %v", minkExtV, required["behat/mink-ext"])
		return
	}

	if minkSymfonyV := "1.2.0"; minkSymfonyV != required["behat/mink-symfony"] {
		t.Errorf("Expected to require %v but got %v", minkSymfonyV, required["behat/mink-symfony"])
		return
	}
	if minkV := "1.6.0"; minkV != required["behat/mink"] {
		t.Errorf("Expected to require %v but got %v", minkV, required["behat/mink"])
		return
	}
}

func Test_Gets_Correct_With_Sub_Dependency_Rules_Applied(t *testing.T) {

	mink := map[string]Dependency{
		"1.6.0": Dependency{
			Name:    "behat/mink",
			Version: "1.6.0",
			Requires: map[string]string{
				"behat/mink-symfony": "<1.2",
			},
		},
		"1.5.1": Dependency{
			Name:    "behat/mink",
			Version: "1.5.1",
			Requires: map[string]string{
				"behat/mink-symfony": "<1.2",
			},
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

	s := NewSolver(packages, mapset.NewSet())
	required, err := s.Solve(root)

	if err != nil {
		t.Errorf("%v", err)
	}

	if minkExtV := "1.4.1"; minkExtV != required["behat/mink-ext"] {
		t.Errorf("Expected to require %v but got %v", minkExtV, required["behat/mink-ext"])
		return
	}

	if minkSymfonyV := "1.1.0"; minkSymfonyV != required["behat/mink-symfony"] {
		t.Errorf("Expected to require %v but got %v", minkSymfonyV, required["behat/mink-symfony"])
		return
	}

	if minkV := "1.6.0"; minkV != required["behat/mink"] {
		t.Errorf("Expected to require %v but got %v", minkV, required["behat/mink"])
		return
	}
}

func Test_Gets_Errors(t *testing.T) {

	mink := map[string]Dependency{
		"1.6.0": Dependency{
			Name:    "behat/mink",
			Version: "1.6.0",
			Requires: map[string]string{
				"behat/mink-symfony": "<1.2",
			},
		},
		"1.5.1": Dependency{
			Name:    "behat/mink",
			Version: "1.5.1",
			Requires: map[string]string{
				"behat/mink-symfony": "<1.2",
			},
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
			"behat/mink-symfony": ">=1.2",
			"behat/mink-ext":     "~1.1",
		},
	}

	s := NewSolver(packages, mapset.NewSet())
	_, err := s.Solve(root)

	if err == nil {
		t.Errorf("Expected an error")
	}
}

func Test_Gets_Replaces(t *testing.T) {

	mink := map[string]Dependency{
		"1.6.0": Dependency{
			Name:    "behat/mink",
			Version: "1.6.0",
			Requires: map[string]string{
				"behat/mink-symfony": "<1.2",
			},
		},
		"1.5.1": Dependency{
			Name:    "behat/mink",
			Version: "1.5.1",
			Requires: map[string]string{
				"behat/mink-symfony": "<1.2",
			},
		},
	}

	replacedDepend := map[string]Dependency{
		"1.6.0": Dependency{
			Name:    "behat/mink-v2",
			Version: "1.6.0",
			Requires: map[string]string{
				"behat/mink-symfony": "<1.2",
			},
			Replaces: map[string]string{
				"behat/mink": "<1.2",
			},
		},
		"1.5.1": Dependency{
			Name:    "behat/mink-v2",
			Version: "1.5.1",
			Requires: map[string]string{
				"behat/mink-symfony": "<1.2",
			},
			Replaces: map[string]string{
				"behat/mink": "<1.2",
			},
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
				"behat/mink-v2": "~1.6",
			},
		},
		"1.1.0": Dependency{
			Name:    "behat/mink-symfony",
			Version: "1.2.5",
			Requires: map[string]string{
				"behat/mink-v2": "~1.6",
			},
		},
	}

	packages := map[string]map[string]Dependency{
		"behat/mink":         mink,
		"behat/mink-v2":      replacedDepend,
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
	replaced := mapset.NewSet()
	s := NewSolver(packages, replaced)
	required, err := s.Solve(root)

	if err != nil {
		t.Errorf("%v", err)
	}

	if minkExtV := "1.4.1"; minkExtV != required["behat/mink-ext"] {
		t.Errorf("Expected to require %v but got %v", minkExtV, required["behat/mink-ext"])
		return
	}

	if minkSymfonyV := "1.1.0"; minkSymfonyV != required["behat/mink-symfony"] {
		t.Errorf("Expected to require %v but got %v", minkSymfonyV, required["behat/mink-symfony"])
		return
	}

	if minkV := "1.6.0"; minkV != required["behat/mink-v2"] {
		t.Errorf("Expected to require %v but got %v", minkV, required["behat/mink"])
		return
	}

	_, notFound := required["behat/mink"]

	if notFound {
		t.Errorf("Didn't expect mink to be found")
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

		if packageName == "behat/mink-ext" && packageRules.Contains("~1.1") {
			foundCount++
		}

		if packageName == "behat/mink-symfony" && packageRules.Contains("~1.1") {
			foundCount++
		}

	}

	if expected := 2; foundCount != expected {
		t.Errorf("expected %v but got %v", expected, foundCount)
	}
}
