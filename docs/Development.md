# Development log

## Understand the task
[Github Issue](https://github.com/screwyprof/form3api/issues/1)

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

Considering that [the contract](https://developer.form3.tech/#9642bfad-d524-49c2-857c-f6becb69bd90) is already defined,
and the server side (the fake account service) is already implemented I'm ready to peruse the docs and run a few queries
to make sure the provided fake service works as expected. After that, I will be able to start working on the client.

## Make some high-level decisions
[ADRs](https://github.com/screwyprof/form3api/blob/create_account/docs/ADR.md)

I spent some time to capture some design decisions in a document.

## Setup local environment with all the necessary tools
[Github Issue](https://github.com/screwyprof/form3api/issues/2)

I've added a `Makefile` to deal with deps, tests and linters. Run `make help` to get a list of all possible targets.

## Setup CI Pipeline
[Github Issue](https://github.com/screwyprof/form3api/issues/4)

One of the first things I usually do when I build a project from scratch is set up the CI pipeline as early as possible. 
The tests must run on every commit. Apart from that I would love to run linters as well. For the sake of this demo, 
I decided to use Github Actions to help me run the builds.

I've decided to commit the whole vendor folder to make CI jobs deterministic. Normally this decision should be based on 
a company-wide convention document, where it is decided whether we need stable and repeatable builds in the CI or would 
give it up to the go mod to deals with the deps to make the repo size smaller.

## Start With the First Feature: Create Account
[Github Issue](https://github.com/screwyprof/form3api/issues/8)

Now that I have a working CI pipeline I can start the actual development. 

I'd like to start with implementing the very first acceptance test for Create Account. At this point I don't expect it 
to turn green, so it will be skipped from running on the CI for a while, however I think it's a good starting point to 
capture the expected input and output for Create Account. 

Basically I'm following the [A-TDD cycle](https://www.agilealliance.org/glossary/atdd/). Write an acceptance test, then 
implement the feature though the TDD cycles until the test passes.

**NB!** As I mentioned it earlier, a real API client should deal with such things as load-balancing, retries, timeouts, 
rate-limits, security concerns, etc... For the simplicity reasons I will neglect these things for now. Another thing to 
keep in mind is validating the response data. Validation is another huge topic on its own which I won't touch. There are 
very many ways to do it. One of the possible options would be to check the response against a pre-defined JSON schema.

## Fetch Account
[Github Issue](https://github.com/screwyprof/form3api/issues/10)

This step is pretty straightforward. I'm following the same [A-TDD cycle](https://www.agilealliance.org/glossary/atdd/)
to drive the development. I updated Readme to include usage section. I also added an example application to demonstrate
some basic client features.

## Delete Account
[Github Issue](https://github.com/screwyprof/form3api/issues/12)

I expected this step to be trivial, however I had to take care of a few things. The major thing is error handling.
The API responses can be tricky. Sometimes they are not consistent with the schema. Now, when an error occurs the client 
tries to unmarshal it into an expected APIError structure and does nothing on failure. In any case the response body is
now preserved for debugging/logging purposes.

Now that the DeleteAccount method is implemented, I can update the previous tests to include a cleanup section, so that
the tests accounts were removed automatically after the tests finish.

## A note on tests structure

I'm following the AAA (Arrange, Act, Assert) pattern to structure the tests. Some tests may
also need a cleanup, in that case I use Annihilate section. Usually I start writing tests from bottom to top: first
I implement the Assert section when I think on what I should check, then goes the Act section, where the SUT is run.
Finally, I set up the SUT in the Arrange section.

In case the test needs to run some clean-up, I add the Annihilate section after Arrange to make sure the cleanup
function runs even when the test asserts fail and stop the test execution flow.