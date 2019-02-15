# blunderlist-gateway [![Build Status](https://travis-ci.com/tomasbasham/blunderlist-gateway.svg?branch=master)](https://travis-ci.com/tomasbasham/blunderlist-gateway) [![Maintainability](https://api.codeclimate.com/v1/badges/428e6cae5d8321a778ed/maintainability)](https://codeclimate.com/github/tomasbasham/blunderlist-gateway/maintainability)

A fictitious todo application through which to teach how to implement a
microservice architecture. For the full list of services required to run this
application visit
[Blunderlist](https://github.com/tomasbasham?utf8=✓&tab=repositories&q=blunderlist)
on GitHub.

This repository implements an API gateway that aggregates data from multiple
backends, acting as the single entry point for all clients. The API gateway
handles requests in one of two ways. Some requests are simply proxied/routed to
the appropriate service whilst others are handled by fanning out to multiple
services.

The intent of this repository is to provided an isolated layer between client
and services to abstract how the application as a whole is separated into its
component services. In addition it provides the most optimal API surface for
the intended client, as opposed to exposing redundant interfaces that
complicated interactions.

## Prerequisites

You will need the following things properly installed on your computer.

* [Git](https://git-scm.com/)
* [Go](https://golang.org/)
* [Docker](https://www.docker.com/)

## Installation

* `git clone <repository-url>` this repository
* `cd blunderlist-gateway`

## Running / Development

* `docker run --rm -it -p 8080:8080 --env-file .env gateway`
* Visit your app at [http://localhost:8080](http://localhost:8080).

### Running Tests

* `docker run --rm gateway golint ./...`
* `docker run --rm gateway go vet ./...`
* `docker run --rm gateway go test ./...`

### Building

* `docker build -t gateway --target builder .`

## Further Reading / Useful Links

* [Go](https://golang.org/)
* [json:api](https://jsonapi.org/)
