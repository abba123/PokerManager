FROM centos
LABEL maintainer="abba123"

USER root
WORKDIR /root

RUN yum install git golang -y \
    && git clone https://github.com/abba123/PokerManager \
    && cd PokerManager \
    && go mod download \
    && go build -o backend

CMD [ "./PokerManager/backend" ]