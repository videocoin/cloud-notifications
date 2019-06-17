FROM golang:1.12.4 as builder
WORKDIR /go/src/github.com/videocoin/cloud-notifications
COPY . .
RUN make build

FROM bitnami/minideb:jessie
RUN apt-get update && apt-get -y install ca-certificates
COPY --from=builder /go/src/github.com/videocoin/cloud-notifications/bin/notifications /opt/videocoin/bin/notifications
COPY --from=builder /go/src/github.com/videocoin/cloud-notifications/templates /opt/videocoin/bin/
CMD ["/opt/videocoin/bin/notifications"]
