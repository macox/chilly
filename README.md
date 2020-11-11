# DX

[![codecov](https://codecov.io/gh/plumming/dx/branch/master/graph/badge.svg)](https://codecov.io/gh/plumming/dx)
[![Go Report Card](https://goreportcard.com/badge/github.com/plumming/dx)](https://goreportcard.com/report/github.com/plumming/dx)
![golangci-lint](https://github.com/plumming/dx/workflows/golangci-lint/badge.svg)
![Check PR can be merged](https://github.com/plumming/dx/workflows/Check%20PR%20can%20be%20merged/badge.svg)
![Set Label](https://github.com/plumming/dx/workflows/Set%20Label/badge.svg)
![Build and test Go](https://github.com/plumming/dx/workflows/Build%20and%20test%20Go/badge.svg)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

# Installation via Homebrew

To install `dx` using homebrew, run the following:

```
brew tap plumming/dx
brew install dx
```

## Basic Usage

### View a list of Open PRs

```
dx get prs
```

### Configure the list of repos to watch

```
dx edit config
```

Add more entries under the `repos:` block e.g.

```
repos:
- org/repo1
- org/repo2
```

### Exclude PRs based on labels

```
dx edit config
```

```
hiddenLabels:
- wip
- do not merge
```
