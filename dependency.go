package dependency

import (
	"fmt"
	"github.com/deckarep/golang-set"
)

type Dependency struct {
	Name     string
	Version  string
	Requires map[string]string
	Replaces map[string]string
}

func (d Dependency) ReplaceSelfVersion() {
	for k, v := range d.Requires {
		if v == "self.version" {
			d.Requires[k] = d.Version
		}
	}
}

func GetPackageNames(d Dependency) mapset.Set {
	packages := mapset.NewSet()

	for packageName, _ := range d.Requires {
		packages.Add(packageName)
	}

	return packages
}


type DependencyFetcher interface {
	Get(dependencyName string) (map[string]Dependency, error)
}

type DependencyRepo struct {
	DependencyNames mapset.Set
	Replaces mapset.Set
	Dependencies    map[string]map[string]Dependency
	Fetcher         DependencyFetcher
}

func (r DependencyRepo) GetAll(dependencies mapset.Set) {

	// check to see if there is 100% intersect if so bail
	f := dependencies.Intersect(r.DependencyNames)
	if f.Cardinality() == dependencies.Cardinality() {
		return
	}
	d := dependencies.Difference(r.DependencyNames)
	subDependencies := mapset.NewSet()
	for _, nameI := range d.ToSlice() {
		name := fmt.Sprintf("%s", nameI)

		r.DependencyNames.Add(name)
		dm, err := r.Fetcher.Get(name)

		if err != nil {
			panic(err)
		}

		r.Dependencies[name] = dm
		for _, sd := range dm {
			for packageName, _ := range sd.Requires {
				subDependencies.Add(packageName)
			}

		}
	}
	r.GetAll(subDependencies)
}

func GetNewRepo(fetcher DependencyFetcher) DependencyRepo {
	dr := DependencyRepo{

		DependencyNames: mapset.NewSet(),
		Replaces: mapset.NewSet(),
		Dependencies:    map[string]map[string]Dependency{},
		Fetcher:         fetcher,
	}

	return dr
}
