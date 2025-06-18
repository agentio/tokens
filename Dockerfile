FROM golang:1.24.4 AS builder
WORKDIR /app
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -v -o tokens .

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/tokens /usr/local/bin/tokens
RUN mkdir /data
WORKDIR /data
ENTRYPOINT ["/usr/local/bin/tokens"]
CMD ["serve"]
