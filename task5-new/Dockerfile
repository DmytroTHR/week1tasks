#syntax=docker/dockerfile:1

##Build
FROM golang:latest AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY README.md ./
COPY *.txt ./

RUN go build -o /task5-docker

##Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /task5-docker /task5-docker

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/task5-docker"]