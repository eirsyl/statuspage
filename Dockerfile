FROM golang:1.9
MAINTAINER Eirik Martiniussen Sylliaas <eirik@sylliaas.no>

RUN mkdir -p /go/src/github.com/eirsyl/statuspage
COPY . /go/src/github.com/eirsyl/statuspage/
WORKDIR /go/src/github.com/eirsyl/statuspage

RUN set -e \
    && go get -u github.com/golang/dep/cmd/dep \
    && dep ensure \
    && go build

CMD ["./statuspage"]