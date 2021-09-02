FROM golang:latest
ENV SRC_DIR=/go/src/admiralhonda/
ENV GOBIN=/go/bin/

WORKDIR ${SRC_DIR}
ADD ./go_lang ${SRC_DIR}


#RUN go mod init admiralhonda
#RUN go get github.com/gorilla/websocket
RUN go build -o go_lang