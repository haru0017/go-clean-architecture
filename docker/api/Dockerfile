FROM golang:1.14.2

ENV GO111MODULE=on

WORKDIR /go/src/github.com/haru0017/go-clean-architecture

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build .

RUN go get github.com/pilu/fresh

CMD [ "fresh" ]