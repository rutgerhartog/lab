# Use the official Nginx image as the base image
FROM nginx

# Install required packages and dependencies
RUN apt-get update && apt-get install -y \
    wget \
    git \
    build-essential \
    libpcre3-dev \
    libxml2-dev \
    libcurl4-openssl-dev \
    libgeoip-dev \
    libyajl-dev \
    zlib1g-dev \
    liblmdb-dev

# Download and build ModSecurity
WORKDIR /tmp
RUN git clone --depth 1 -b v3/master --single-branch https://github.com/SpiderLabs/ModSecurity
WORKDIR /tmp/ModSecurity
RUN ./build.sh
RUN ./configure
RUN make
RUN make install

# Download and build the ModSecurity Nginx connector
WORKDIR /tmp
RUN git clone --depth 1 https://github.com/SpiderLabs/ModSecurity-nginx.git
WORKDIR /tmp/ModSecurity-nginx
RUN ./configure --with-compat --add-dynamic-module=../ModSecurity-nginx
RUN make modules

# Copy the ModSecurity Nginx connector .so file to the Nginx modules directory
RUN cp objs/ngx_http_modsecurity_module.so /etc/nginx/modules

# Copy the ModSecurity configuration files
RUN mkdir /etc/nginx/modsecurity
COPY modsecurity.conf-recommended /etc/nginx/modsecurity/modsecurity.conf
COPY unicode.mapping /etc/nginx/modsecurity/unicode.mapping

# Add ModSecurity rules (customize as needed)
COPY custom-modsecurity-rules.conf /etc/nginx/modsecurity/

# Modify the Nginx configuration to include ModSecurity
RUN sed -i '/http {/a\    modsecurity on;\n    modsecurity_rules_file /etc/nginx/modsecurity/custom-modsecurity-rules.conf;' /etc/nginx/nginx.conf

# Expose the HTTP and HTTPS ports
EXPOSE 80
EXPOSE 443

# Start Nginx with ModSecurity
CMD ["nginx", "-g", "daemon off;"]
