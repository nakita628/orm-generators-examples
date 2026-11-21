# PHP 8.3 with the PDO drivers and illuminate/database: the Eloquent harness runs in this image.
FROM php:8.3-cli
RUN apt-get update && apt-get install -y --no-install-recommends libpq-dev unzip git \
 && docker-php-ext-install pdo_pgsql pdo_mysql \
 && rm -rf /var/lib/apt/lists/*
COPY --from=composer:2 /usr/bin/composer /usr/bin/composer
WORKDIR /harness
COPY composer.json composer.lock* ./
RUN composer install --no-interaction --no-progress
