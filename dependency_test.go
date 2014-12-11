package dependency

import (
	"errors"
	"github.com/deckarep/golang-set"
	"testing"
)

func Test_GetPackageNames(t *testing.T) {
	root := Dependency{
		Name: "App",
		Requires: map[string]string{
			"behat/mink-symfony": "~1.1",
			"behat/mink-ext":     "~1.1",
		},
	}

	actual := GetPackageNames(root)
	expected := mapset.NewSet()
	expected.Add("behat/mink-symfony")
	expected.Add("behat/mink-ext")

	if !expected.Equal(actual) {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func Test_GetAllDependencies(t *testing.T) {
	packages := mapset.NewSet()
	packages.Add("behat/mink-ext")
	packages.Add("behat/mink-symfony")

	f := MockFetcher{}

	r := GetNewRepo(f)

	r.GetAll(packages)

	expected := mapset.NewSet()
	expected.Add("behat/mink-ext")
	expected.Add("behat/mink-symfony")
	expected.Add("behat/mink")

	if !r.DependencyNames.Equal(expected) {
		t.Errorf("Expected %s, but got %s", expected, r.DependencyNames)
	}
}

func Test_Replaces_SelfVersion(t *testing.T) {
	root := Dependency{
		Name: "App",
		Version: "1.2",
		Requires: map[string]string{
			"behat/mink-symfony": "self.version",
			"behat/mink-ext":     "~1.1",
		},
	}
	root.ReplaceSelfVersion()

	if versionNum := "1.2"; root.Requires["behat/mink-symfony"] != versionNum {
		t.Errorf("Expected %v but got %v", versionNum, root.Requires["behat/mink-symfony"])
	}
}

type MockFetcher struct {
}

func (mf MockFetcher) Get(dependencyName string) (map[string]Dependency, error) {

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


	minkSeleniumBrowser := map[string]Dependency{
		"1.2.0": Dependency{
			Name:    "behat/mink-selenium",
			Version: "1.2.5",
			Requires: map[string]string{
				"behat/mink-v2": "~1.6",
			},
		},
		"1.1.0": Dependency{
			Name:    "behat/mink-selenium",
			Version: "1.2.5",
			Requires: map[string]string{
				"behat/mink-v2": "~1.5",
			},
		},
	}


	minkV2 := map[string]Dependency{
		"1.6.0": Dependency{
			Name:    "behat/mink-v2",
			Version: "1.6.0",
			Requires: map[string]string{
				"behat/mink-ext": "<1.2",
			},
			Replaces: map[string]string{
				"behat/mink": "<1.2",
			},
		},
		"1.5.1": Dependency{
			Name:    "behat/mink-v2",
			Version: "1.5.1",
			Requires: map[string]string{
				"behat/mink-ext": "<1.2",
			},
			Replaces: map[string]string{
				"behat/mink": "<1.2",
			},
		},
	}

	switch dependencyName {
	case "behat/mink-ext":
		return minkExt, nil
	case "behat/mink-symfony":
		return minkSymfonyBrowser, nil
	case "behat/mink":
		return mink, nil
	case "behat/mink-v2":
		return minkV2, nil
	case "behat/mink-selenium":
		return minkSeleniumBrowser, nil
	}

	return nil, errors.New("Not found")
}
