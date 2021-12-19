FROM nginx
LABEL maintainer="abba123"

USER root
WORKDIR /root

RUN apt-get update \
    && apt-get install -y git \
    && git clone https://github.com/abba123/PokerManager \
    && mv PokerManager/vue/web/dist/index.html /usr/share/nginx/html/ -f \
    && mv PokerManager/vue/web/dist/static /usr/share/nginx/html/ -f \