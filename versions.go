package dependency

import (
	"fmt"
	"github.com/deckarep/golang-set"
	"github.com/mcuadros/go-version"
	"sort"
)

func GetVersionNumbers(versions map[string]Dependency) mapset.Set {
	versionNumbers := mapset.NewSet()

	for versionNumber, _ := range versions {
		versionNumbers.Add(versionNumber)
	}

	return versionNumbers
}

func PrepVersionNumbers(versionNumbers mapset.Set) []string {
	versions := []string{}

	for _, value := range versionNumbers.ToSlice() {
		versions = append(versions, fmt.Sprintf("%s", value))
	}
	sort.Sort(sort.Reverse(VersionSlice(versions)))
	return versions
}

type VersionSlice []string

func (p VersionSlice) Len() int           { return len(p) }
func (p VersionSlice) Less(i, j int) bool { return version.Compare(p[i], p[j], "<") }
func (p VersionSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
