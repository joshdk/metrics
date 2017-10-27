FROM alpine:3.5

RUN apk --update add tar \
    && mkdir /metrics-server

COPY metrics-server.sls.tgz .

RUN tar -xvf metrics-server.sls.tgz --strip-components=1 -C /metrics-server

ENTRYPOINT ["/metrics-server/service/bin/linux-amd64/metrics-server"]

CMD ["/server.yml"]
