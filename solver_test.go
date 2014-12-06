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
		t.Errorf("Expected to  require %v but got %v", minkSymfonyV, required["behat/mink"])
        return
	}
}
