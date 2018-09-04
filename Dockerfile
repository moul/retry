FROM golang:1.10 as build
COPY . /go/src/moul.io/retry
WORKDIR /go/src/moul.io/retry
RUN make install

FROM alpine
COPY --from=build /go/bin/retry /bin/retry
ENTRYPOINT ["retry"]
