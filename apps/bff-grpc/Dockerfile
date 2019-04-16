# build binary and download binary
#===============================================================
FROM golang:1.12-stretch as build-env
RUN apt-get update && apt-get install -y build-essential --no-install-recommends && rm -rf /var/lib/apt/lists/*
RUN mkdir -p /src
WORKDIR /src

## dowload entrykit
ENV ENTRYKIT_VERSION=0.4.0
RUN curl -fsL https://github.com/progrium/entrykit/releases/download/v${ENTRYKIT_VERSION}/entrykit_${ENTRYKIT_VERSION}_Linux_x86_64.tgz | \
  tar xfz - -C /usr/local/bin/

# build app
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
COPY --from=build-env /usr/local/bin//entrykit /usr/local/bin/entrykit
RUN entrykit --symlink
COPY config config
RUN echo -e '#!/bin/sh\nexec /usr/local/bin/render $(find /app/config/ -name \*.json) "$@"' > /render.sh  && chmod +x /render.sh

CMD ["/render.sh", "--", "/app/server"]
