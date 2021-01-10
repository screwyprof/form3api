# Architecture Decision Records  (ADR)
Please note that the stated decisions bellow should be properly documented and put in docs folder in the project's repo.
See https://github.com/joelparkerhenderson/architecture_decision_record for more information about ADR.

For the simplicity sake, I won't create detailed ADRs with all the reasoning behind, but simply create a list here.

## Programming Language
GoLang 1.14 with go modules. 

## Project Structure
I'm going to be as close as possible to [Standard Go Project Layout](https://github.com/golang-standards/project-layout) 
convention, but for simplicity reasons, to make the top level package API cleaner, I'd like to avoid the excessive 
nesting with sub folders like `pkg`, given that we deal with a library.

## Linters and quality control tools
For now, I've decided to use [golang-ci](https://github.com/golangci/golangci-lint) only. More tools maybe added later.
[CodeBbeat](https://codebeat.co/open-source/go), and [BetterCode](https://bettercodehub.com/) can be good candidates.

## Testing tools and libraries
I'll start with the build-in golang testing framework along with my simple assertion helper.
if the things will ge too complicated I may consider switching to something more sophisticated like 
[Testify](https://github.com/stretchr/testify).

## Code coverage
Considering that the project shouldn't get too complicated or way too big, [Codecov](https://about.codecov.io/) seems
to be a good choice.

## CI Pipeline
One of the first things I usually do when I build a project from scratch is set up the CI pipeline as early as possible.
The tests must be run on every commit. Apart from that I would love to run linters as well. If it was an application
I would also run the actual building process, but given that we have a library, this stage can be omitted. For the sake
of this demo, I decided to use Github Actions to help me run the builds. Makefile was updated to include CI targets.
Apart from that test coverage support added via codecov.io

Also, I've decided to commit the whole `vendor` folder to make CI jobs deterministic. Normally this decision should be
based on a company-wide convention document, where it is decided whether we need stable and repeatable builds in the CI
or would give it up to the go mod to deal with the deps to make the repo size smaller.

## Useful tools
Postman can come in handy to deal with REST API queries.