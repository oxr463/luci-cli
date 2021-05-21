FROM openwrtorg/rootfs:x86-64

RUN mkdir /var/lock && \
    touch /var/lock/opkg.lock && \
    opkg update && \
    opkg install curl \
                 jq \
                 luci \
                 luci-ssl

