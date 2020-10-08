FROM golang:1.14
ARG APP=addnumbers
WORKDIR /arena
COPY . /usr/local/go/src/arena
RUN  go build   -o './bin/'${APP}  /usr/local/go/src/arena/cmd/${APP}/${APP}.go
CMD ["bash", "-c", "./bin/addnumbers"]
