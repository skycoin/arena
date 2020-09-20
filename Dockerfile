FROM golang:1.14.2-stretch

RUN mkdir -p /go/src/github.com/skycoin/arena
WORKDIR /go/src/github.com/skycoin/arena

COPY . .

RUN go install -v

CMD ["arena"]