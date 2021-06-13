FROM golang:1.10-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /SDXchain
RUN cd /SDXchain && make SDX

FROM alpine:latest

LABEL maintainer="deva@swapdex.net"

WORKDIR /SDXchain

COPY --from=builder /SDXchain/build/bin/SDX /usr/local/bin/SDX

RUN chmod +x /usr/local/bin/SDX

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/usr/local/bin/SDX"]

CMD ["--help"]
