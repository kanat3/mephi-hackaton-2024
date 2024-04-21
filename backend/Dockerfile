FROM golang:1.22.2-bookworm

RUN apt-get update -y && apt-get install -y bash python3 python3-pip
RUN rm /usr/lib/python3.*/EXTERNALLY-MANAGED
RUN pip3 install ultralytics && pip3 install lapx

ENV GOPATH=/export/hackaton-2024
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

COPY docker-entrypoint.sh /export/docker-entrypoint.sh
RUN chmod +x /export/docker-entrypoint.sh
WORKDIR /export

ENTRYPOINT ["/export/docker-entrypoint.sh"]  
