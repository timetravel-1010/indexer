# Build the app
FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o api cmd/api/main.go

# Run the app
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/api .

EXPOSE 8080

CMD ["./api"]
