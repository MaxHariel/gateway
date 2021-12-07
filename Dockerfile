FROM golang:1.17.3

WORKDIR /go/src

CMD ["tail", "-f", "/dev/null"]