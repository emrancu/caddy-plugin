# Stage 1: Build custom Caddy binary
FROM golang:1.20 AS builder

RUN go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest

WORKDIR /build

# Enable Go modules
ENV GO111MODULE=on


# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download module dependencies
RUN go mod download

RUN xcaddy build \
    --with github.com/caddyserver/replace-response

# Stage 2: Create final Docker image
FROM caddy

COPY --from=builder /build/caddy /usr/bin/caddy

EXPOSE 80

CMD ["caddy", "run", "--config", "/etc/caddy/Caddyfile"]