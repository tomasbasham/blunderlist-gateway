# blunderlist-gateway

This README outlines the details of collaborating on this Go application. A
short introduction of this app could easily go here.

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
