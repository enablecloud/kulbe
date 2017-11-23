FROM golang:1.9.1
COPY . /go/src/github.com/enablecloud/kulbe
WORKDIR /go/src/github.com/enablecloud/kulbe
RUN ls /go/src;export GOPATH=/go;export PATH=$PATH:/go/bin;mkdir /go/bin;curl https://glide.sh/get | sh;glide up -v --force;make build

FROM alpine:latest

COPY --from=0 /go/src/github.com/enablecloud/kulbe/kulbe /bin/kulbe


MAINTAINER Sebastien DIAZ <sebastien.diaz@gmail.com>

ARG VCS_REF
ARG BUILD_DATE

# Metadata
LABEL org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.docker.dockerfile="/Dockerfile"

ENV HELM_LATEST_VERSION="v2.6.2"

RUN apk add --update ca-certificates \
 && apk add --update -t deps wget \
 && wget http://storage.googleapis.com/kubernetes-helm/helm-${HELM_LATEST_VERSION}-linux-amd64.tar.gz \
 && gunzip helm-${HELM_LATEST_VERSION}-linux-amd64.tar.gz \
 && tar -xvf helm-${HELM_LATEST_VERSION}-linux-amd64.tar \
 && mv linux-amd64/helm /usr/bin \
 && apk del --purge deps \
 && rm /var/cache/apk/* \
 && rm -f /helm-${HELM_LATEST_VERSION}-linux-amd64.tar.gz
CMD ["/bin/kulbe"]
EXPOSE 9300