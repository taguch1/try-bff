# download binary
#===============================================================
FROM node:10.15.3-stretch AS  build-env
# dowload entrykit
ENV ENTRYKIT_VERSION=0.4.0
RUN curl -fsL https://github.com/progrium/entrykit/releases/download/v${ENTRYKIT_VERSION}/entrykit_${ENTRYKIT_VERSION}_Linux_x86_64.tgz | \
  tar xfz - -C /usr/local/bin/

WORKDIR /app

# save cache
ADD package.json .
ADD package-lock.json .

VOLUME /app/node_modules

ADD . .
RUN make setup && make release

FROM nginx:1.15.12-alpine
EXPOSE 80 443
COPY --from=build-env /app/dist /usr/share/nginx/html
# TODO ignore from multi stage copy
RUN rm /usr/share/nginx/html/env
COPY nginx/env /usr/share/nginx/html/env
COPY nginx/health /usr/share/nginx/html
COPY nginx/error /usr/share/nginx/html
COPY nginx/conf /etc/nginx

ENV TEMPLATE_DIR=/etc/nginx/
COPY --from=build-env /usr/local/bin/entrykit /usr/local/bin/entrykit
RUN entrykit --symlink \
  && echo -e '#!/bin/sh\ntemplate_dir=${TEMPLATE_DIR:-/app/config/}\ntemplates="$(find ${template_dir} -type f -name "*tmpl" | sed -e 's/.tmpl//g')"\nexec /usr/local/bin/render ${templates} "$@"' > /render.sh \
  && chmod +x /render.sh

CMD ["/render.sh", "--", "/usr/sbin/nginx"]

STOPSIGNAL SIGQUIT
