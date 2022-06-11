# create by xiexianbin, Github Action for sync code to cdn
FROM alpine:latest

# Dockerfile build cache 
ENV REFRESHED_AT 2020-01-11

LABEL "com.github.actions.name"="Github Action for sync code to cdn"
LABEL "com.github.actions.description"="Github Action for sync code to cdn."
LABEL "com.github.actions.icon"="home"
LABEL "com.github.actions.color"="green"
LABEL "repository"="http://github.com/x-actions/gsync"
LABEL "homepage"="http://github.com/x-actions/gsync"
LABEL "maintainer"="xiexianbin<me@xiexianbin.cn>"

LABEL "Name"="Github Action for sync code to cdn"

ENV LC_ALL C.UTF-8
ENV LANG en_US.UTF-8
ENV LANGUAGE en_US.UTF-8
# change VERSION when make a release, v1.0.0
ENV VERSION "v1.2.0"

RUN apk update && \
    apk add --no-cache git git-lfs bash wget curl openssh-client tree && \
    rm -rf /var/cache/apk/* && \
    cd /tmp && \
    curl -s https://api.github.com/repos/x-actions/go-sync/releases/latest | \
    sed -r -n '/browser_download_url/{/linux.tar.gz/{s@[^:]*:[[:space:]]*"([^"]*)".*@\1@g;p;q}}' | xargs wget && \
    chmod +x /tmp/gsync-linux && \
    mv /tmp/gsync-linux /usr/local/bin/gsync
#    wget https://github.com/x-actions/go-sync/releases/download/${VERSION}/gsync-linux && \

ADD entrypoint.sh /
RUN chmod +x /entrypoint.sh

WORKDIR /github/workspace
ENTRYPOINT ["/entrypoint.sh"]
