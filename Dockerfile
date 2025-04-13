FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app
COPY . .

RUN go mod tidy

ENV TARGETS="linux/amd64 windows/amd64 darwin/amd64 darwin/arm64"

RUN mkdir -p /dist && \
    for TARGET in $TARGETS; do \
        OS=$(echo $TARGET | cut -d/ -f1); \
        ARCH=$(echo $TARGET | cut -d/ -f2); \
        EXT=""; \
        [ "$OS" = "windows" ] && EXT=".exe"; \
        echo "Building for $OS/$ARCH"; \
        CGO_ENABLED=0 GOOS=$OS GOARCH=$ARCH go build -ldflags="-s -w" \
        -o /dist/cai-${OS}-${ARCH}${EXT} ./cmd/cai; \
    done
