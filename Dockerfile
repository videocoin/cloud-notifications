FROM alpine:3.7

RUN apk add --no-cache ca-certificates

COPY bin/notifications /opt/videocoin/bin/notifications
ADD templates /opt/videocoin/bin/

CMD ["/opt/videocoin/bin/notifications"]
