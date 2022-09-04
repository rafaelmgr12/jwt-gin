FROM golang:alpine AS build-env
ENV GOPATH /go
WORKDIR /go/src/github.com/rafaelmgr12/jwt-gin
COPY . /go/src/github.com/rafaelmgr12/jwt-gin
RUN cd /go/src/github.com/rafaelmgr12/jwt-gin && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/jwt github.com/rafaelmgr12/jwt-gin/

FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*
WORKDIR /app
COPY --from=build-env /go/src/github.com/rafaelmgr12/jwt-gin/build/jwt /app
COPY .env /app

EXPOSE 8080

ENTRYPOINT [ "./jwt" ]