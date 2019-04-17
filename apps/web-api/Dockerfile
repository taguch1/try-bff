# download binary
#===============================================================
FROM golang:1.12-stretch as build-env

## dowload entrykit
ENV ENTRYKIT_VERSION=0.4.0
RUN curl -fsL https://github.com/progrium/entrykit/releases/download/v${ENTRYKIT_VERSION}/entrykit_${ENTRYKIT_VERSION}_Linux_x86_64.tgz | \
  tar xfz - -C /usr/local/bin/

FROM nginx:1.15.12-alpine
EXPOSE 80 443

COPY health /usr/share/nginx/html
COPY error /usr/share/nginx/html
COPY conf/ /etc/nginx
COPY index.html /usr/share/nginx/html

ENV TEMPLATE_DIR=/etc/nginx/
COPY --from=build-env /usr/local/bin/entrykit /usr/local/bin/entrykit
RUN entrykit --symlink \
  && echo -e '#!/bin/sh\ntemplate_dir=${TEMPLATE_DIR:-/app/config/}\ntemplates="$(find ${template_dir} -type f -name '*tmpl' | sed -e 's/.tmpl//g')"\nexec /usr/local/bin/render ${templates} "$@"' > /render.sh \
  && chmod +x /render.sh
CMD ["/render.sh", "--", "/usr/sbin/nginx"]

STOPSIGNAL SIGQUIT


