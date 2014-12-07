package dependency

import (
	"github.com/deckarep/golang-set"
)

func GetVersionNumbers(versions map[string]Dependency) mapset.Set {
	versionNumbers := mapset.NewSet()

    for versionNumber, _ := range versions {
        versionNumbers.Add(versionNumber)
    }

	return versionNumbers
}
