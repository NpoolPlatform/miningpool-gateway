FROM uhub.service.ucloud.cn/entropypool/alpine:3.16

WORKDIR /app

COPY miningpool-gateway /app

ENTRYPOINT ["/app/miningpool-gateway"]

EXPOSE 50860 50861 50862 50863

CMD ["run"]
