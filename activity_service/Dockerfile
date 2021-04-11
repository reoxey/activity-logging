FROM golang:1.16-alpine as builder

LABEL maintainer="reoxey"

WORKDIR /useract/activity

COPY . .

RUN CGO_ENABLED=0 go test -v . && go build -o app .

FROM alpine

COPY --from=builder /useract/activity/app .

EXPOSE 9001

CMD ["/app"]
