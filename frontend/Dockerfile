FROM docker:24.0.7-dind

ARG DEBIAN_FRONTEND=noninteractive
RUN apk update && apk add npm bash

COPY docker-entrypoint.sh /export/docker-entrypoint.sh
RUN chmod +x /export/docker-entrypoint.sh
WORKDIR /export

ENTRYPOINT ["/export/docker-entrypoint.sh"]   
