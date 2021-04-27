FROM alpine:3.13.5

ARG version

ENV VERSION ${version:-v3.0}
ENV ARCHIVE_URL https://github.com/CastawayLabs/cachet-monitor/releases/download/${VERSION}/cachet_monitor_linux_amd64
ENV CACHET_CONFIG_NAME /etc/cachet-monitor.json

RUN apk add wget git curl bash grep
RUN adduser -S -s /bin/bash -u 1001 ci
RUN wget ${ARCHIVE_URL} && \
    mv cachet_monitor_linux_amd64 /usr/bin/cachet_monitor && \
    chmod a+x /usr/bin/cachet_monitor

USER 1001

ENTRYPOINT /usr/bin/cachet_monitor -c $CACHET_CONFIG_NAME