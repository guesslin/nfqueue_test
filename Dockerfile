FROM golang:1.12
LABEL maintainer="guesslin1986@gmail.com"

RUN apt-get update && apt-get install -y iptables pkg-config libnetfilter-queue-dev
RUN mkdir -p /apps
WORKDIR /apps
COPY . /apps/

RUN go build -o nfqueue

CMD /apps/nfqueue
