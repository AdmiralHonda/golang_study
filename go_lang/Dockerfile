FROM golang:latest
ADD . /go
WORKDIR /go
RUN go get iok_server/v1
CMD [ "go","run","main.go" ]