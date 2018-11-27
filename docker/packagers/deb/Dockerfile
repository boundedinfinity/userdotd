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
RUN go get && go build -o userdotd

RUN apt-get install ruby ruby-dev rubygems build-essential -y
RUN gem install --no-ri --no-rdoc fpm

RUN fpm -s dir -t deb -n userdotd -v 0.0.0 -C $GOPATH/src/github.com/boundedinfinity/userdotd

# for testing vv
RUN dpkg-deb -I userdotd_0.0.0_amd64.deb
RUN apt install ./userdotd_0.0.0_amd64.deb
RUN apt list --installed | grep userdotd
