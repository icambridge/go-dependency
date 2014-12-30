go-dependency
=============

[![Build Status](https://travis-ci.org/icambridge/go-dependency.svg)](https://travis-ci.org/icambridge/go-dependency)[![Coverage Status](https://img.shields.io/coveralls/icambridge/go-dependency.svg)](https://coveralls.io/r/icambridge/go-dependency)

A version dependency resolving library in Go.

Installation
------------

The recommended way to install go-dependency

```go
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


```go
type MockFetcher struct {
}

func (mf MockFetcher) Get(dependencyName string) (map[string]dependency.Dependency, error) {
    return map[string]dependency.Dependency{}, nil
}
```

Pass that to the Repo

```go
f := MockFetcher{}
r := dependency.GetNewRepo(f)
```

Then to use

```go
d := dependency.Dependency{Require: map[string]string{"behat/behat": "~1.3"}}

p := dependency.GetPackageNames(d)

s := dependency.NewSolver(repo.Dependencies, repo.Replaces)

required, err := s.Solve(d)
```