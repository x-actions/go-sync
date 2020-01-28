# create by xiexianbin, Github Action for sync code to cdn
FROM alpine:latest

# Dockerfile build cache 
ENV REFRESHED_AT 2020-01-11

LABEL "com.github.actions.name"="Github Action for sync code to cdn"
LABEL "com.github.actions.description"="Github Action for sync code to cdn."
LABEL "com.github.actions.icon"="home"
LABEL "com.github.actions.color"="green"
LABEL "repository"="http://github.com/xiexianbin/cdn-sync"
LABEL "homepage"="http://github.com/xiexianbin/cdn-sync"
LABEL "maintainer"="xiexianbin<me@xiexianbin.cn>"

LABEL "Name"="Github Action for sync code to cdn"
LABEL "Version"="1.0.0"

ENV LC_ALL C.UTF-8
ENV LANG en_US.UTF-8
ENV LANGUAGE en_US.UTF-8

RUN apk update && apk add --no-cache git bash go && rm -rf /var/cache/apk/*

RUN export GOPATH=`pwd` && \
    mkdir src && \
    cd src && \
    git clone https://github.com/x-actions/cdn-sync.git && \
    mv cdn-sync cdn-sync.bak && \
    cp -rp cdn-sync.bak/src/cdn_sync . && \
    rm -rf cdn-sync.bak && \
    cd cdn_sync && \
    # GOOS=linux GOARCH=amd64 go build -tags netgo && \
    go build && \
    cp cdn_sync /usr/local/bin

ADD entrypoint.sh /
RUN chmod +x /entrypoint.sh

WORKDIR /github/workspace
ENTRYPOINT ["/entrypoint.sh"]
