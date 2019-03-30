# Iro-iro Color Palettes
[![CodeFactor](https://www.codefactor.io/repository/github/sanisoclem/iro-iro/badge)](https://www.codefactor.io/repository/github/sanisoclem/iro-iro)
[![codecov](https://codecov.io/gh/sanisoclem/iro-iro/branch/master/graph/badge.svg)](https://codecov.io/gh/sanisoclem/iro-iro)

A color palette sharing site.

## Status

| Platform     | Go Ver | Build status                                                                                                                                                  |
|--------------|--------:|:--:|
| Windows      |      ? | ? |
| Linux        | `1.12` | [![Build Status](https://travis-ci.com/sanisoclem/iro-iro.svg?branch=master)](https://travis-ci.com/sanisoclem/iro-iro) |
| Linux        |  `tip` | ? |


## Prerequesites

- git
- [go](getgo)
- [sam cli](saminstall)
- make (optional)

## Getting Started

The easiest way to get started is to run `make all`.

### Install dependencies

You can install dependencies using `make deps` or:

```bash
$ go get -u github.com/aws/aws-lambda-go/...
github.com/aws/aws-lambda-go (download)
Fetching https://gopkg.in/urfave/cli.v1?go-get=1
Parsing meta tags from https://gopkg.in/urfave/cli.v1?go-get=1 (status code 200)
get "gopkg.in/urfave/cli.v1": found meta tag get.metaImport{Prefix:"gopkg.in/urfave/cli.v1", VCS:"git", RepoRoot:"https://gopkg.in/urfave/cli.v1"} at https://gopkg.in/urfave/cli.v1?go-get=1
gopkg.in/urfave/cli.v1 (download)
github.com/stretchr/testify (download)
...
```

### Build the app

To build all binaries, just run: `make build`. If make is not available, you can build each binary by:

```bash
# replace $name with the name of the function
$ GOOS=linux GOARCH=amd64 go build -o ./bin/function-$name ./functions/$name/main.go
```

### Run Tests

```bash
$ make test # or
$ go test -v ./...
=== RUN   TestHandler
=== RUN   TestHandler/Return_id_when_created
=== RUN   TestHandler/Return_error_if_exists
--- PASS: TestHandler (0.00s)
    --- PASS: TestHandler/Return_id_when_created (0.00s)
    --- PASS: TestHandler/Return_error_if_exists (0.00s)
PASS
ok      _/mnt/d/Mel/Workspace/src/github.com/sanisoclem/iro-iro/functions/new-palette   0.012s
```

## Deployment

TBD

[getgo]:https://golang.org/doc/install
[saminstall]:https://aws.amazon.com/serverless/sam/