FROM golang:alpine
RUN apk update && apk upgrade && \
  apk add --no-cache bash git openssh curl
#ENV GLIDE_DOWNLOAD_URL https://github.com/Masterminds/glide/releases/download/$GLIDE_VERSION/glide-$GLIDE_VERSION-linux-amd64.zip
#ENV GLIDE_DOWNLOAD_URL https://github.com/Masterminds/glide/releases/download/v0.12.3/glide-v0.12.3-darwin-386.zip
RUN curl https://glide.sh/get | sh


RUN mkdir -p /go/src/salvation-army-api
ADD .  /go/src/salvation-army-api

WORKDIR /go/src/salvation-army-api
RUN glide install
RUN go build -v -o private
EXPOSE 8000
#ENTRYPOINT go run *.go
ENTRYPOINT ./private

#WORKDIR "/opt"
#
#ADD .docker_build/salvation-army-api /opt/bin/salvation-army-api
#
#CMD ["/opt/bin/salvation-army-api"]