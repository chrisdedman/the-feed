# Use the official Golang image
FROM golang:1.22

RUN apt update \
    && apt install -y build-essential inotify-tools postgresql-client git make gcc \
    && apt clean

ADD . /app

WORKDIR /app

EXPOSE 3000

CMD make run