FROM debian:stable-slim AS main 

RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    bind9 \
    bind9utils \
    dnsutils \
    nano \
    net-tools 


COPY config/ /etc/bind/


CMD ["/usr/sbin/named", "-f"]
