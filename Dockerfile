FROM golang:1-alpine AS builder

ENV GOPATH=/go
RUN mkdir -p "/go/src/github.com/szazeski/wifiled"
WORKDIR /go/src/github.com/szazeski/wifiled

#COPY go.mod .
#COPY go.sum .
#RUN go mod download

COPY . .

RUN go build -o /wifiled

## Deploy
FROM alpine:3
LABEL maintainer="steve@wifiled.org"
LABEL build_date="2024-08-10"
LABEL built_version="0.5.0"

WORKDIR /
COPY --from=builder /wifiled /wifiled

ENTRYPOINT ["/wifiled"]


# docker buildx build --platform linux/arm/v7,linux/arm64,linux/amd64 --tag szazeski/wifiled:0.4.0 --tag szazeski/wifiled:latest -push .
