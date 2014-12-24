go-dependency
=============

[![Build Status](https://travis-ci.org/icambridge/go-dependency.svg)](https://travis-ci.org/icambridge/go-dependency)

A version dependency resolving library in Go.

Installation
------------

The recommended way to install go-dependency

```
    get github.com/icambridge/go-dependency
```

Examples
--------

How import the package

```go
import (
    "github.com/icambridge/go-dependency"
)
```

Create your own DependencyFetcher


```type MockFetcher struct {
}

func (mf MockFetcher) Get(dependencyName string) (map[string]dependency.Dependency, error) {
    return map[string]dependency.Dependency{}, nil
}
```

Pass that to the Repo

```
f := MockFetcher{}
r := dependency.GetNewRepo(f)
```


```d := dependency.Dependency{Require: map[string]string{"behat/behat": "~1.3"}}

p := dependency.GetPackageNames(d)

s := dependency.NewSolver(repo.Dependencies, repo.Replaces)

required, err := s.Solve(d)
```