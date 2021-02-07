FROM golang:latest

RUN apt install iputils-ping -y

COPY bin/client /opt
WORKDIR /opt

RUN sleep 10
CMD ./client
