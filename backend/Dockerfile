FROM golang:1.21.4-alpine3.18

RUN apk update && apk add bash

ENV GOPATH=/export/hackaton-2024
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

COPY docker-entrypoint.sh /export/docker-entrypoint.sh
RUN chmod +x /export/docker-entrypoint.sh
WORKDIR /export

ENTRYPOINT ["/export/docker-entrypoint.sh"]     
