# form3api
![form3api-client](https://github.com/screwyprof/form3api/workflows/form3api-client/badge.svg)
[![codecov](https://codecov.io/gh/screwyprof/form3api/branch/main/graph/badge.svg?token=4BN07UH560)](https://codecov.io/gh/screwyprof/form3api)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Form3 Take Home Exercise

## Introduction

This is a possible solution to [Form3 Take Home Exercise](https://github.com/form3tech-oss/interview-accountapi) done by 
[Maksim Shcherbo](https://www.linkedin.com/in/maxim-shcherbo-3204582b/) aka [*screwyprof*](https://github.com/screwyprof).

## Development Log
I've decided to [document](docs/Development.md) the steps I've taken to implement the task.

## Architecture Decision Records
I'm going to document some of my high-level decisions in the following [ADR.md](docs/ADR.md) document.

### Building the project
Simple run `make` or `make all` to install all the required dependencies and tools, to run linters and all the tests.
Use `make help` to get information about additional targets.

### Running tests
To run unit and E2E tests run `make test`. To run unit tests only `make test-unit`. To run E2E tests `make test-e2e`

### Running linters
To run local linters use `make lint`

### Formatting code
Use `make fmt` to run go fmt

