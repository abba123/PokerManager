FROM golang
LABEL maintainer="abba123"

USER root
WORKDIR /root

RUN git clone https://github.com/abba123/PokerManager \
    && cd PokerManager \
    && go mod download \
    && go build -o backend