FROM docker:24.0.7-dind

RUN apk update && apk add bash
RUN rm -rf /usr/local/go
RUN wget https://go.dev/dl/go1.21.4.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.21.4.linux-amd64.tar.gz

ENV GOPATH=/export/hackaton-2024
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

COPY docker-entrypoint.sh /export/docker-entrypoint.sh
RUN chmod +x /export/docker-entrypoint.sh
WORKDIR /export

ENTRYPOINT ["/export/docker-entrypoint.sh"]     
