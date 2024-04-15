FROM debian:bookworm

ARG DEBIAN_FRONTEND=noninteractive
RUN apt-get update -y && apt-get install -y --no-install-recommends python3 npm && \
            apt-get install -y --no-install-recommends build-essential checkinstall libssl-dev
RUN wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash && \
            source ~/.bashrc && nvm install v19.0.0

COPY docker-entrypoint.sh /export/docker-entrypoint.sh
RUN chmod +x /export/docker-entrypoint.sh
WORKDIR /export

ENTRYPOINT ["/export/docker-entrypoint.sh"]     
