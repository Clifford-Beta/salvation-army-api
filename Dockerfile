#FROM golang
#
#ENV GLIDE_VERSION 0.12.3
#
#RUN apt-get update \
# 	&& apt-get install -y unzip --no-install-recommends \
#	&& rm -rf /var/lib/apt/lists/*
FROM golang:alpine
RUN apk update && apk upgrade && \
  apk add --no-cache bash git openssh curl
#ENV GLIDE_DOWNLOAD_URL https://github.com/Masterminds/glide/releases/download/$GLIDE_VERSION/glide-$GLIDE_VERSION-linux-amd64.zip
#ENV GLIDE_DOWNLOAD_URL https://github.com/Masterminds/glide/releases/download/v0.12.3/glide-v0.12.3-darwin-386.zip
RUN curl https://glide.sh/get | sh


RUN mkdir -p /go/src/sendyit/private_api
ADD .  /go/src/sendyit/private_api

WORKDIR /go/src/sendyit/private_api
RUN glide install
RUN go build -v -o private
EXPOSE 8000
#ENTRYPOINT go run *.go
ENTRYPOINT ./private