FROM ubuntu:latest

RUN apt-get update
RUN apt-get install -y wget git gcc

RUN wget -P /tmp https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz

RUN tar -C /usr/local -xzf /tmp/go1.9.2.linux-amd64.tar.gz
RUN rm /tmp/go1.9.2.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" "$GOPATH/src/github.com" "$GOPATH/src/github.com/boundedinfinity" && chmod -R 777 "$GOPATH"

WORKDIR $GOPATH/src/github.com/boundedinfinity
RUN git clone https://github.com/boundedinfinity/userdotd.git

WORKDIR $GOPATH/src/github.com/boundedinfinity/userdotd
RUN git fetch && git checkout init-shell-detection
RUN go get && go install

RUN userdotd
