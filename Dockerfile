FROM golang:1.22-bookworm AS base
ARG username
ARG exec_user_id
ENV GOCACHE="/home/logger/go/cache/build"
ENV GOMODCACHE="/home/logger/go/pkg/mod"
ENV GOPATH="/home/logger/go"
RUN groupadd -g $exec_user_id -o $username
RUN useradd -r -u $exec_user_id -g $username $username -m
RUN mkdir -p /srv/logger
RUN mkdir -p /home/logger/go
RUN chown $username:$username /home/logger/go -R
USER $username:$username
WORKDIR /home/logger/src

