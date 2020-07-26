FROM golang:1.14-alpine as builder
WORKDIR /go/src/

COPY cmd ./cmd
COPY go.mod go.sum ./
RUN go build -o bot cmd/*


FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/bot .
CMD ["./bot"]