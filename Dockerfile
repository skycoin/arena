#Here we using multistage docker file to reduce image size
# Build Stage
FROM golang:1.14-alpine AS build-env
RUN apk --no-cache add build-base git mercurial gcc make
ADD ./ /src
RUN cd /src && go build -o arena-add

# Final Stage
FROM alpine
COPY --from=build-env /src/arena-add .
ENTRYPOINT ./arena-add