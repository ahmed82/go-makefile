FROM alpine:latest

LABEL maintainer="Ahmed Alsalih <a.alsalih2@gmail.com>" \ 
      description="GO-makefile"

RUN echo "http://dl-4/alpinelinux.org.org/alpine/edge/community" >> /etc/apk/repositories;
RUN echo "http://dl-4/alpinelinux.org.org/alpine/edge/testing" >> /etc/apk/repositories;

RUN apk update && apk add make g++ bash

# mkdir /code
VOLUME /code

# dc /codeDefault && make all

ENTREYPOINT ["bash"]
