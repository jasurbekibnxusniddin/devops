# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /src
COPY go.mod .
COPY main.go .
RUN go build -o /hello

# Run stage
FROM alpine:3.18
RUN adduser -D -u 1000 appuser
COPY --from=builder /hello /hello
USER appuser
EXPOSE 8081
ENTRYPOINT ["/hello"]
