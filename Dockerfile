FROM resin/rpi-raspbian

# We need this to run our cross-compiled go binary
# as it looks by default for '/lib/ld-linux.so.3' as interpreter
RUN ln -s /lib/arm-linux-gnueabihf/ld-linux.so.3 /lib/ld-linux.so.3

COPY build/dht-server-linux-arm-5 /usr/bin/dht-server-linux-arm-5

CMD ["/usr/bin/dht-server-linux-arm-5"]

EXPOSE 3000