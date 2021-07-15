FROM golang:1.12-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

ADD . /sdx
RUN cd /sdx && make sdx

FROM alpine:latest

WORKDIR /SmartDex-Chain

COPY --from=builder /SmartDex-Chain/build/bin/sdx /usr/local/bin/sdx

RUN chmod +x /usr/local/bin/sdx

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/usr/local/bin/sdx"]

CMD ["--help"]
