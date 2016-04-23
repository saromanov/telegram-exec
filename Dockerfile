FROM golang:1.6-onbuild
RUN go get github.com/saromanov/telegram-exec
WORKDIR $GOPATH/src/github.com/saromanov/telegram-exec
RUN go build -o $GOPATH/bin/telegram .
RUN $GOPATH/bin/telegram
