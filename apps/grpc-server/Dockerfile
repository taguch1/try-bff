FROM golang:1.12-stretch as build-env
RUN apt-get update && apt-get install -y build-essential --no-install-recommends && rm -rf /var/lib/apt/lists/*
RUN mkdir -p /src
WORKDIR /src
COPY . .
RUN make build
RUN make install PREFIX=/usr/local INSTALL_BIN=${INSTALL_BIN:-app}

FROM alpine:3.9.2
WORKDIR /app
COPY --from=build-env /usr/local/bin/app /app/server
CMD ["/app/server"]
