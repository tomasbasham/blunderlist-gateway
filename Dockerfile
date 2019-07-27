FROM golang:1.13-alpine3.10 as builder

WORKDIR /usr/src/app

RUN apk add --no-cache ca-certificates gcc git libc-dev \
  && go get -u golang.org/x/lint/golint \
  && go get -u github.com/golang/mock/mockgen

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go generate ./...

COPY . .

# Build and install a static binary, stripping DWARF debugging information and
# preventing the generation of the Go symbol table.
RUN GOOS=linux GOARCH=amd64 go install -a -ldflags '-w -s -linkmode external -extldflags "-static"' ./cmd/gateway

FROM scratch

COPY --from=builder /go/bin/gateway /gateway
COPY --from=builder /etc/ssl/certs /etc/ssl/certs

ENTRYPOINT ["/gateway"]
