FROM golang:1.16-alpine as builder

RUN echo "Build...."

ENV GO111MODULE=on

WORKDIR /go/src/app/discounts-applier
RUN apk add git
RUN apk add build-base musl-dev

COPY ./ ./
RUN go mod download -x

RUN ls -la
RUN go get -d -v cmd/api
RUN GOOS=linux GOARCH=amd64 go build -a -v -o api cmd/api/main.go
RUN GOOS=linux GOARCH=amd64 go build -a -v -o writer cmd/writer/main.go
COPY ./cmd/writer/data.json ./

ARG MONGO_URI
EXPOSE 8080
ENTRYPOINT /bin/sh