FROM golang:1.12.4 as builder
WORKDIR /go/src/github.com/videocoin/cloud-notifications
COPY . .
RUN make build

FROM bitnami/minideb:jessie
RUN apt update && apt -y install ca-certificates
COPY --from=builder /go/src/github.com/videocoin/cloud-notifications/bin/notifications /opt/videocoin/bin/notifications
CMD ["/opt/videocoin/bin/notifications"]
