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
[CodeBeat](https://codebeat.co/open-source/go), and [BetterCode](https://bettercodehub.com/) can be good candidates.

## Testing tools and libraries
I'll start with the build-in golang testing framework along with my simple assertion helper.
if the things will ge too complicated I may consider switching to something more sophisticated like 
[Testify](https://github.com/stretchr/testify).

## Code coverage
Considering that the project shouldn't get too complicated or way too big, [Codecov](https://about.codecov.io/) seems
to be a good choice.

## CI Pipeline
Github Actions

## Useful tools
Postman can come in handy to deal with REST API queries.