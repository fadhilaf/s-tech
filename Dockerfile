# Stage 1: Build
FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main ./cmd/main.go

# Stage 2: Run
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/.env . 
RUN mkdir -p /app/static/img

EXPOSE 8000
CMD ["./main"]
