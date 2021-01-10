# form3api
![form3api-client](https://github.com/screwyprof/form3api/workflows/form3api-client/badge.svg)
[![codecov](https://codecov.io/gh/screwyprof/form3api/branch/main/graph/badge.svg?token=4BN07UH560)](https://codecov.io/gh/screwyprof/form3api)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Form3 Take Home Exercise

## Introduction

This is a possible solution to [Form3 Take Home Exercise](https://github.com/form3tech-oss/interview-accountapi) done by 
[Maksim Shcherbo](https://www.linkedin.com/in/maxim-shcherbo-3204582b/) aka [*screwyprof*](https://github.com/screwyprof).

## Foreword

Having read the description of the task I feel like I need to elaborate on a few things first. Given that the task 
doesn't have a clear set of acceptance criteria, and I don't have a chance to ask the questions to gather the 
requirements, I think it would be reasonable for me to rely on my experience to make a few assumptions.

Normally, I would prefer an API client to be automatically generated from the server-side code, given that the API 
contracts are well-defined. I've worked on a few micro-service based projects with a large set of services (60+), 
so I was able to witness that supporting a few versions of clients for a particular micro-service manually is a very 
tedious and error-prone process.

It is stated in the task that no external HTTP-client libraries can be used. Building such a library is a huge topic on 
its own with a lot of nitty-gritty to take into account such as throttling, rate-limiting, authentication, 
foul-tolerance, error-handling, load-balancing, etc… In real life, I wouldn't invent the wheel if I didn't have a good 
reason to do so.

Usually, when I work on an API service I would start with an API scheme like OpenAPI to create an API contract. 
After that, I would generate a mock-server pre-populated with sample request-response objects. Then I would write 
an acceptance test to check that when the given request is sent, the expected response is received. All these stages 
can be done in Postman for example. Apart from that, Newman can be used to run those tests as part of the CI process.

Once I have a mock-server endpoint, I'm ready to work on the client part, which in the latter end can be used in 
integration tests for the server-side - the best way to make sure that the client and server work properly. Of course, 
the details may vary from project to project, some more advanced tools like Pact maybe used, but the whole idea stays 
the same - server and client are tested together.

Providing that [the contract](https://developer.form3.tech/#9642bfad-d524-49c2-857c-f6becb69bd90) is already defined, 
and the server side (the fake account service) is already implemented I'm ready to peruse the docs and run a few queries
to make sure the provided fake service works as expected. After that, I will be able to start working on the client.

## Local environment and tools
I use GoLang 1.14 with go modules. [golang-ci](https://github.com/golangci/golangci-lint) helps me run linters. 
Postman can come in handy to run queries to the API service. In order to make my life easier I've created a Makefile.

### Building the project
Simple run `make` or `make all` to install all the required dependencies and tools, to run linters and all the tests.
Use `make help` to get information about additional targets.

### Running tests
To run unit and E2E tests run `make test`. To run unit tests only `make test-unit`. To run E2E tests `make test-e2e`

### Running linters
To run local linters use `make lint`

### Formatting code
Use `make fmt` to run go fmt

## CI with Github Actions
One of the first things I usually do when I build a project from scratch is set up the CI pipeline as early as possible.
The tests must be run on every commit. Apart from that I would love to run linters as well. If it was an application
I would also run the actual building process, but given that we have a library, this stage can be omitted. For the sake
of this demo, I decided to use Github Actions to help me run the builds. Makefile was updated to include CI targets.

Also, I've decided to commit the whole vendor folder to make CI jobs deterministic. Normally this decision should be
based on a company-wide convention document, where it is decided whether we need stable and repeatable builds in the CI
or would give it up to the go mod to deal with the deps to make the repo size smaller.
