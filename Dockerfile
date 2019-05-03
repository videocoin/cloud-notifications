FROM alpine:3.7

COPY bin/notifications /opt/videocoin/bin/notifications
COPY templates.yaml /opt/videocoin/bin/templates.yaml

CMD ["/opt/videocoin/bin/notifications"]
