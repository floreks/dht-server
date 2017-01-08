FROM floreks/kubepi-base

COPY build/dht-server-linux-arm-6 /usr/bin/dht-server

ENTRYPOINT ["/usr/bin/dht-server"]

EXPOSE 3000
