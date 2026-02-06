FROM golang:1.25-trixie AS build

WORKDIR /src

ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    go build -trimpath -o=bootstrap -v main.go

FROM debian:12.13-slim

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
      libimage-exiftool-perl \
      tzdata \
      ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=build /src/bootstrap /bootstrap

CMD ["/bootstrap"]
