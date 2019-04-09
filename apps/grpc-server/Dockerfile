# build binary and download binary
#===============================================================
FROM golang:1.12-stretch as build-env
RUN GRPC_HEALTH_PROBE_VERSION=v0.2.0 && \
  wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
  chmod +x /bin/grpc_health_probe
RUN apt-get update && apt-get install -y build-essential --no-install-recommends && rm -rf /var/lib/apt/lists/*
RUN mkdir -p /src
WORKDIR /src

# cache
ADD go.mod .
VOLUME /go/pkg/mod
COPY . .
RUN make build
RUN make install PREFIX=/usr/local INSTALL_BIN=${INSTALL_BIN:-app}

# copy binary
#===============================================================
FROM alpine:3.9.2
WORKDIR /app
COPY --from=build-env /usr/local/bin/app /app/server
COPY --from=build-env /bin/grpc_health_probe /bin/grpc_health_probe

# TODO render config
CMD ["/app/server"]
