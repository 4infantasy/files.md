# syntax=docker/dockerfile:1.7

# --- build stage ---------------------------------------------------------
FROM golang:1.24-alpine AS build
WORKDIR /src

# Cache deps first.
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG VERSION=dev
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags "-s -w -X main.Version=${VERSION}" \
    -o /out/server ./cmd/server

# --- runtime stage -------------------------------------------------------
FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata && \
    addgroup -g 1000 app && adduser -D -u 1000 -G app app

WORKDIR /app
COPY --from=build /out/server /app/server
COPY web /app/web

# Storage is a named volume so files persist across container restarts.
RUN mkdir -p /app/storage /app/tokens && chown -R app:app /app
VOLUME ["/app/storage", "/app/tokens"]


USER app

ENTRYPOINT ["/app/server"]
