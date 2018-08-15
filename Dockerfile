FROM golang:1.10 as build
COPY . /go/src/github.com/moul/retry
WORKDIR /go/src/github.com/moul/retry
RUN make install

FROM alpine
COPY --from=build /go/bin/retry /bin/retry
ENTRYPOINT ["retry"]
