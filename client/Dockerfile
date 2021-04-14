FROM golang:1.16-alpine as builder

LABEL maintainer="reoxey"

WORKDIR /client

COPY . .

RUN CGO_ENABLED=0 go build -o cli .

FROM alpine

COPY --from=builder /client/cli .
