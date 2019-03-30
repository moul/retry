FROM golang:1.12-alpine as build
RUN apk add --update make
COPY . /go/src/moul.io/retry
WORKDIR /go/src/moul.io/retry
RUN make install

FROM alpine
COPY --from=build /go/bin/retry /bin/retry
ENTRYPOINT ["retry"]
