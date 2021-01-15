# form3api
![form3api-client](https://github.com/screwyprof/form3api/workflows/form3api-client/badge.svg)
[![codecov](https://codecov.io/gh/screwyprof/form3api/branch/main/graph/badge.svg?token=4BN07UH560)](https://codecov.io/gh/screwyprof/form3api)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Form3 Take Home Exercise

## Introduction

This is a solution to [Form3 Take Home Exercise](https://github.com/form3tech-oss/interview-accountapi) done by 
[Maksim Shcherbo](https://www.linkedin.com/in/maxim-shcherbo-3204582b/) aka [*screwyprof*](https://github.com/screwyprof).

## Usage

```go
import "github.com/screwyprof/form3api"
```

Create a new Client instance, then call the corresponding methods to get what you want, for example:

```go
c := form3api.NewClient(nil, "http://localhost:8080/v1")
	
accountID := "51646a03-a52e-4e51-b405-cf2b8078c1a8"
acc, err := c.FetchAccount(context.Background(), form3api.FetchAccount{AccountID: accountID})
```
Take a loot at [Client Demo](example/client_demo.go) for more details.

## Running integration tests
To run tests in docker use `make test-docker`, or its equivalent `docker-compose up --build --abort-on-container-exit`

## Development

### Development Log
[This document](docs/Development.md) guides your though the steps I've taken to implement this library. Each step is 
represented by a separate [Github Issue](https://github.com/screwyprof/form3api/issues?q=is%3Aissue+is%3Aclosed).

### Architecture Decision Records
High-level decisions on the project architecture are captured in the following [ADR.md](docs/ADR.md) document.

### Building the project
Simply run `make` or `make all` to install all the required dependencies and tools, to run linters and all the tests.
Use `make help` to get information about additional targets.

### Running tests
To run unit and E2E tests run `make test`. To run unit tests only `make test-unit`. To run E2E tests `make test-e2e`

### Running linters
Install the linters by running `make tools`. This step is required once only. To run linters use `make lint`.

### Formatting code
Use `make fmt` to run go fmt
