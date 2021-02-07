FROM golang:latest

RUN apt install iputils-ping -y

COPY ./bin/server /opt
WORKDIR /opt

RUN sleep 10
CMD ./server
#EXPOSE 9999/tcp
