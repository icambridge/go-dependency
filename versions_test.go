package dependency

import (
	"github.com/deckarep/golang-set"
	"reflect"
	"testing"
)

func Test_Get_Versions_ReturnsSetWith(t *testing.T) {

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

	actual := GetVersionNumbers(minkSymfonyBrowser)

	expected := mapset.NewSet()
	expected.Add("1.2.0")
	expected.Add("1.1.0")

	if !expected.Equal(actual) {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func Test_Prep_ReturnsSortedSlice(t *testing.T) {
	input := mapset.NewSet()
	input.Add("1.2.0")
	input.Add("1.1.0")
	input.Add("1.1.1")
	input.Add("1.2.1")

	actual := PrepVersionNumbers(input)
	expected := []string{"1.2.1", "1.2.0", "1.1.1", "1.1.0"}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
