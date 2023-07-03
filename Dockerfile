FROM caddy:latest-builder AS builder

RUN xcaddy build \
	--with github.com/emrancu/caddy-plugin@latest

FROM caddy:latest

COPY --from=builder /usr/bin/caddy /usr/bin/caddy

CMD ["caddy", "run", "--config", "/etc/caddy/Caddyfile", "--adapter", "caddyfile"]
