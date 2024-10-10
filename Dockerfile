# Build Stage
FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/server

# Run Stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/server .

# Copy migration files
COPY migrations/ /migrations/

# Set environment variables (override in docker-compose.yml if needed)
ENV DB_HOST=db
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_PASSWORD=postgres
ENV DB_NAME=filestorage
ENV PORT=8080

EXPOSE 8080

CMD ["./server"]
