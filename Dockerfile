FROM golang:1.7
COPY . /go/src/github.com/moul/retry
WORKDIR /go/src/github.com/moul/retry
RUN make install
ENTRYPOINT ["retry"]
