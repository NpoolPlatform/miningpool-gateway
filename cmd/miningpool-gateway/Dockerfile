FROM uhub.service.ucloud.cn/entropypool/alpine:3.16

WORKDIR /app

COPY miningpool-gateway /app

ENTRYPOINT ["/app/miningpool-gateway"]

EXPOSE 50610 50611 50612 50613

CMD ["run"]
