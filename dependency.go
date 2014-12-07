package dependency

import (
	"github.com/deckarep/golang-set"
)

type Dependency struct {
	Name     string
	Version  string
	Requires map[string]string
}

func GetPackageNames(d Dependency) mapset.Set {
	packages := mapset.NewSet()

	for packageName, _ := range d.Requires {
		packages.Add(packageName)
	}


	return packages
}

type DependencyFetcher interface {
 	Get(dependencyName string) (Dependency, error)
}

type DependencyRepo struct {

}
//
// func (r DependencyRepo) GetAll(d Dependency) mapset.Set {
//
//
//
//
// 	return input := mapset.NewSet()
// }
