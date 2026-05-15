# Build stage
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Explicitly target linux/amd64
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -o /app/bin/agent-node    ./cmd/agent-node/main.go
RUN go build -o /app/bin/agent-cli     ./cmd/agent-cli/main.go
RUN go build -o /app/bin/agent-gateway ./cmd/agent-gateway/main.go

# Final stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/bin/agent-node    .
COPY --from=builder /app/bin/agent-cli     .
COPY --from=builder /app/bin/agent-gateway .
RUN chmod +x ./agent-node ./agent-cli ./agent-gateway
EXPOSE 8080 9090
CMD ["./agent-node"]