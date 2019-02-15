# blunderlist-gateway [![Build Status](https://travis-ci.com/tomasbasham/blunderlist-gateway.svg?branch=master)](https://travis-ci.com/tomasbasham/blunderlist-gateway) [![Maintainability](https://api.codeclimate.com/v1/badges/428e6cae5d8321a778ed/maintainability)](https://codeclimate.com/github/tomasbasham/blunderlist-gateway/maintainability) [![Pact](https://blunderlist.pact.dius.com.au/pacts/provider/blunderlist-gateway/consumer/blunderlist/latest/badge.svg)](https://blunderlist.pact.dius.com.au)

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

Optionally to run deployments manually the following tools must be present:

* [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
* [kustomize](https://github.com/kubernetes-sigs/kustomize)

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

## Design

The gateway makes several calls to two other services before aggregating the
results into a single JSON payload. This can be seen in the following:

![list-tasks][list-tasks]\
[diagram][list-tasks-diagram]

## Debugging Running Pods

When there is a problematic pod running within a cluster it may not be
desirable to destroy it without first understanding what went wrong. Instead
the pod should be removed from the load balancer and inspected whilst no longer
serving traffic.

This has been accomplished through the use of labels and selectors within the
Kubernetes manifest files. In particular there is a serving: true label that
indicates to Kubernetes that a pod should be placed within the load balancer
and should be serving traffic. For a pod to be removed from the load balancer
it's labels must be edited in place.

    kubectl --namespace=neptune label pods/<POD NAME> --overwrite serving=false

The replication controller will spin up a new pod to replace the one taken from
the load balancer whilst the problematic pod will remain active for inspection
but not available to serve traffic.

## Further Reading / Useful Links

* [Go](https://golang.org/)
* [json:api](https://jsonapi.org/)

[list-tasks]: /diagrams/list-tasks.svg?raw=true&sanitize=true "List Tasks"
[list-tasks-diagram]: https://sequencediagram.org/index.html#initialData=C4S2BsFMAIBkQM7GgFQIYINYIFA7QK7AD2AdgQLYBGkATtPkcWgMagBuaoZeADmrVAsQ-UsgBEAQQAKASWgBxLpADuaAJ7joGaAHNla9Tn6CQw0RJTEAJsWgBlOuzOQtOkreMChItGOjiAMLEFBSQ-o60ziyu2gjQLCFhYnj6wKoaALQAfB7EAFwAPIngxLT5AMS0kNbZAPR18EjoWAgNhXUlZdkAFACUOHmZOWkZ6vnAGNh4pcS80ABmZdCQrAAW0JNYOKOGOYmh4cBFXeVVNfWNiMDBh2JtdR2nvVuYAwfJwMPZuxr5H0ccOFrNAgA
