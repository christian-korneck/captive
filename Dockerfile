FROM golang:1.13 AS builder
WORKDIR /go/src/github.com/christian-korneck/captive
COPY . .
ENV GO111MODULE=on CGO_ENABLED=0
RUN go build ./...

FROM scratch
COPY --from=builder /go/src/github.com/christian-korneck/captive/captive /
#nobody:nogroup
USER 65534:65534
ENTRYPOINT ["/captive"]

