FROM golang:alpine

ARG BINARY_NAME="gnocchi-token-renewal"

ENV TZ="Europe/Rome"

RUN set -xe \
    && apk add --update --no-cache tzdata \
    curl

WORKDIR "/go"

COPY "build/${BINARY_NAME}" "bin/app"

ENTRYPOINT [ "/go/bin/app" ]
