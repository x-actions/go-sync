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

#     wget https://github.com/x-actions/git-mirrors/releases/download/${VERSION}/git-mirrors-linux && \
#     chmod +x /tmp/git-mirrors-linux && \
#     mv /tmp/git-mirrors-linux /usr/local/bin/git-mirrors
RUN apk update && apk add --no-cache git git-lfs bash wget curl openssh-client tree && rm -rf /var/cache/apk/*

RUN mkdir /usr/local/gsync/ && \
    cd /usr/local/gsync/ && \
    curl -s https://api.github.com/repos/x-actions/go-sync/releases/latest | \
    sed -r -n '/browser_download_url/{/linux.tar.gz/{s@[^:]*:[[:space:]]*"([^"]*)".*@\1@g;p;q}}' | xargs wget && \
    tar xzf *linux.tar.gz -C /usr/local/gsync/ && \
    cp /usr/local/gsync/gsync_*_linux/gsync /usr/local/bin/ && \
    rm -rf /usr/local/gsync/

ADD entrypoint.sh /
RUN chmod +x /entrypoint.sh

WORKDIR /github/workspace
ENTRYPOINT ["/entrypoint.sh"]
