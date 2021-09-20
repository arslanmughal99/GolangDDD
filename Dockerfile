FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -ldflags="-s -w" -o app .

RUN chmod +x -R /build

FROM alpine

# Init for PID=host
COPY --from=builder /build/app /app

RUN apk add dumb-init
ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["sh", "-c", "/app"]