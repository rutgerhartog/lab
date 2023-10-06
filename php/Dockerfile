FROM curlimages/curl AS downloader 

RUN curl -sL https://github.com/moodle/moodle/archive/refs/tags/v3.10.3.tar.gz --output /tmp/moodle.tar.gz



# -------------------------------------------------------------------------------------------------

FROM php:7.3-apache AS main 

COPY --from=downloader /tmp/moodle.tar.gz /tmp/moodle.tar.gz
RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y libpq-dev libpng-dev libicu-dev libzip-dev
RUN docker-php-ext-configure pgsql && docker-php-ext-install -j$(nproc) pgsql 
RUN docker-php-ext-configure zip && docker-php-ext-install -j$(nproc) zip
RUN docker-php-ext-configure gd && docker-php-ext-install -j$(nproc) gd
RUN docker-php-ext-configure intl && docker-php-ext-install -j$(nproc) intl
RUN mkdir -p /var/www/html && mkdir -p /var/www/moodledata && tar xvzf /tmp/moodle.tar.gz --directory /tmp  && mv /tmp/moodle-3.10.3/* /var/www/html


# # -------------------------------------------------------------------------------------------------
# FROM php:5.5-apache AS main 

# ENV DEBIAN_FRONTEND=noninteractive

# COPY --from=downloader /tmp/app.tar.gz /tmp/app.tar.gz
# COPY sources.list /etc/apt/sources.list 
# COPY --from=downloader /tmp/composer.phar /usr/local/bin/composer 

# RUN apt-get update; apt-get install -y --allow-unauthenticated debian-archive-keyring && apt-get install -y --force-yes php5-pgsql unzip libpq-dev
# RUN mkdir -p /var/www/html && cd /tmp && tar xvzf /tmp/app.tar.gz -C /var/www/html
# # RUN echo "extension=pgsql.so" > /usr/local/etc/php/conf.d/custom.ini

# RUN docker-php-ext-configure pgsql && docker-php-ext-install -j$(nproc) pgsql 

# # Install composer 
# RUN chmod +x /usr/local/bin/composer

# USER 33
# RUN cd /var/www/html/phpmyadmin-RELEASE_4_8_1 && composer update --no-dev


# USER 0 
# RUN chmod -R 777 /var/www/html

# # -------------------------------------------------------------------------------------------------