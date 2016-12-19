FROM golang:1.7.4-alpine

WORKDIR /opt

COPY src ./src
COPY Makefile .

RUN apk update \
  && apk add make \
  && make build

WORKDIR /opt/bin

EXPOSE 5000

ENTRYPOINT ["./wikid-server"]
