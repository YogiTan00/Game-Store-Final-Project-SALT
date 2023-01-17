FROM golang:1.19-alpine

RUN apk add build-base

ADD ./ /app
WORKDIR /app

ENV SERVICENAME finalproject

ENV APP_HOME /app
ENV GO111MODULE on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go mod vendor
RUN go mod verify

RUN go build -o ${SERVICENAME} /app/cmd/main.go

CMD ["/app/finalproject"]