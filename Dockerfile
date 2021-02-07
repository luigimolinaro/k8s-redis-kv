FROM ubuntu:20.04

#Install packet
RUN apt update
RUN apt install -y wget

# install Golang
ENV GOPATH=/usr/src/.go
RUN wget -qO- https://golang.org/dl/go1.15.6.linux-amd64.tar.gz | tar -C /usr/local -xzf -
RUN mkdir -p "${GOPATH}"
ENV PATH="${PATH}:/usr/local/go/bin:${GOPATH}/bin"
ENV GO111MODULE=on

# get go dependencies
RUN go get github.com/julienschmidt/httprouter
RUN go get github.com/garyburd/redigo/redis

ADD go-rest /opt/go-rest/
WORKDIR /opt/go-rest/
#RUN go build
CMD /opt/go-rest/go-rest
