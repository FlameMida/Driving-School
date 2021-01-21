FROM golang:alpine

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

WORKDIR /go/src/gin-vue-admin
COPY . .
RUN go env && go build -o server .

FROM alpine:latest
LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

WORKDIR /go/src/gin-vue-admin
COPY --from=0 /go/src/Driving-school/server ./
COPY --from=0 /go/src/Driving-school/config.yaml ./
COPY --from=0 /go/src/Driving-school/resource ./resource

EXPOSE 8888

ENTRYPOINT ./server
