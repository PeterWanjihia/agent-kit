# Build stage
FROM golang:1.22-alpine AS builder

# Install build dependencies for CGO 
RUN apk add --no-cache git gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Building  core infrastructure components
RUN go build -o /app/bin/agent-node ./cmd/agent-node/main.go
RUN go build -o /app/bin/agent-cli ./cmd/agent-cli/main.go
RUN go build -o /app/bin/agent-gateway ./cmd/agent-gateway/main.go

# Final stage
FROM alpine:latest
RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Copying the compiled binaries
COPY --from=builder /app/bin/agent-node .
COPY --from=builder /app/bin/agent-cli .
COPY --from=builder /app/bin/agent-gateway .


EXPOSE 8080 9090

CMD ["./agent-node"]