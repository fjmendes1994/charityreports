FROM debian:9.7

RUN apt-get update && \
    apt-get -y upgrade

RUN apt-get install curl git build-essential -y

RUN curl https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz | tar xz && \
    mv go /usr/local




ENV GOPATH=$HOME/go
ENV GOBIN=/usr/local/go/bin
ENV PATH="${PATH}:${GOBIN}"

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

ADD ./target /

ENTRYPOINT /charityreports/bin