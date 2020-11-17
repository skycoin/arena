FROM golang:1.15-alpine as build

RUN apk add --update --no-cache musl-dev gcc linux-headers git bash \
 && rm -rf /var/cache/apk/* \
 && wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.32.2 \
 && go get golang.org/x/tools/cmd/goimports \
 && go get github.com/vektra/mockery/v2/... \
 && go get -u github.com/golang/dep/...

ADD . /go/src/github.com/skycoin/arena
WORKDIR /go/src/github.com/skycoin/arena

RUN dep ensure -vendor-only \
 && go vet ./... \
 && golangci-lint run ./... \
 && go test --timeout=1m -coverprofile=coverage.out  ./... \
 && go install -ldflags "-X github.com/skycoin/arena/pkg/version.Release=$(git rev-parse HEAD) -X github.com/skycoin/arena/pkg/version.BuildTime=$(date -u +'%Y-%m-%dT%H:%M:%SZ')" ./...

# Final stage ( minimized for production)
FROM alpine:3.12
RUN mkdir -p /app

COPY --from=build /go/bin/* /usr/local/bin/

WORKDIR /app