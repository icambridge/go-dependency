package dependency

import (
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

// func Test_GetAllDependencies(t *testing.T) {
//
//     root := Dependency{
//         Name: "App",
//         Requires: map[string]string{
//             "behat/mink-symfony": "~1.1",
//             "behat/mink-ext":     "~1.1",
//             },
//         }
//
// }
//
//
// type MockFetcher struct {
//
// }
//
// func (mf MockFetcher) Get(dependencyName string) (Dependency, error) {
//
// }
