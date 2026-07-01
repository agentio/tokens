FROM golang:1.26.4 AS builder
WORKDIR /app
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -v -o tokens .

FROM gcr.io/distroless/static-debian13
COPY --from=builder /app/tokens /usr/local/bin/tokens
WORKDIR /data
ENTRYPOINT ["/usr/local/bin/tokens"]
CMD ["serve"]
