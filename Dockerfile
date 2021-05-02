#first stage - builder
FROM golang:1.14.0-stretch as builder

COPY . /go-qr-generator
WORKDIR /go-qr-generator

RUN CGO_ENABLED=0 GOOS=linux go build -o go-qr-generator


#second stage 
FROM alpine:latest

RUN apk add --no-cache tzdata

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /root/

COPY --from=builder /go-qr-generator .

CMD ["./go-qr-generator"]