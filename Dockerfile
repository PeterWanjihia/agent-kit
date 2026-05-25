
FROM golang:1.25-alpine AS builder

WORKDIR /app


RUN apk add --no-cache git


COPY go.mod go.sum ./
RUN go mod download


COPY . .


ENV CGO_ENABLED=0


RUN mkdir -p /build


RUN go build -ldflags="-s -w" -o /build/agent-kit ./cmd/agent-cli

# Optional additional binaries 
#RUN go build -ldflags="-s -w" -o /build/agent-node ./cmd/agent-node
#RUN go build -ldflags="-s -w" -o /build/agent-gateway ./cmd/agent-gateway



FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates


COPY --from=builder /build/agent-kit /usr/local/bin/agent-kit

# COPY --from=builder /build/agent-node /usr/local/bin/agent-node
# COPY --from=builder /build/agent-gateway /usr/local/bin/agent-gateway

RUN chmod +x /usr/local/bin/agent-kit

ENTRYPOINT ["agent-kit"]


CMD ["--help"]