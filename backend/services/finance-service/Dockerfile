# syntax=docker/dockerfile:1

# --- Build Stage ---
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

# --- Run Stage ---
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 CMD wget --spider -q http://localhost:8080/health || exit 1
CMD ["./main"]
