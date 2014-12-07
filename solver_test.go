package dependency

import (
	"reflect"
	"testing"
	set "github.com/deckarep/golang-set"
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

	s := Solver{packages}
	required := s.Solve()

	if minkV := "1.5.1"; minkV != required["behat/mink"] {
		t.Errorf("Expected to require %v but got %v", minkV, required["behat/mink"])
		return
	}

	if minkExtV := "1.4.1"; minkExtV != required["behat/mink-ext"] {
		t.Errorf("Expected to require %v but got %v", minkExtV, required["behat/mink"])
		return
	}

	if minkSymfonyV := "1.1.0"; minkSymfonyV != required["behat/mink-symfony"] {
		t.Errorf("Expected to require %v but got %v", minkSymfonyV, required["behat/mink"])
		return
	}
}

func Test_Find_Unique_Rules(t *testing.T) {

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
	s := Solver{}
	actual := s.findUniqueRules(minkSymfonyBrowser)

	expected := map[string]map[string][]string{
		"behat/mink": map[string][]string{
			"~1.5": []string{"1.1.0"},
			"~1.6": []string{"1.2.0"},
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected to %v, got %v", actual, expected)
	}
}

func Test_Merge_Rules(t *testing.T) {

	name := "behat/mink"
	symfony := map[string]map[string][]string{
		name: map[string][]string{
			"~1.5": []string{"1.1.0"},
			"~1.6": []string{"1.2.0"},
		},
	}
	minkExt := map[string]map[string][]string{
		name: map[string][]string{
			"~1.5": []string{"2.1.0"},
		},
	}

	allRules := map[string]map[string]map[string][]string{
		"behat/mink-ext":     minkExt,
		"behat/mink-symfony": symfony,
	}

	s := Solver{}
	actual := s.mergeRules(allRules)

	if len(actual) != 1 {
		t.Errorf("Expected to be one got %v", len(actual))
	}

	if len(actual[name]) != 2 {
		t.Errorf("Expected to be two")
	}

	if len(actual[name]["~1.5"]) != 2 {
		t.Errorf("Expected to be two")
	}
	expected15 := []string{"behat/mink-ext", "behat/symfony"}
	Sx := NewSetFromStringSlice(actual[name]["~1.5"])
	Sy := NewSetFromStringSlice(expected15)

	if Sx.Equal(Sy) {
		t.Errorf("Got %v, expected %v", actual[name]["~1.5"], expected15)
	}
	expected16 := []string{"behat/symfony"}
	Sx = NewSetFromStringSlice(actual[name]["~1.6"])
	Sy = NewSetFromStringSlice(expected16)
	if len(actual[name]["~1.6"]) != 1 {
		t.Errorf("Expected to be one")
	}

	if Sx.Equal(Sy) {
		t.Errorf("Got %v, expected %v", actual[name]["~1.6"], expected16)
	}
}


func NewSetFromStringSlice(s []string) set.Set {
	a := set.NewSet()
	for _, item := range s {
		a.Add(item)
	}
	return a
}
