# Ruby 3.4 with the gems of Gemfile.lock: the Active Record harness runs in this image, because
# Active Record 8.1 needs a newer Ruby than many machines have.
FROM ruby:3.4-slim
RUN apt-get update \
 && apt-get install -y --no-install-recommends build-essential libpq-dev default-libmysqlclient-dev libyaml-dev pkg-config \
 && rm -rf /var/lib/apt/lists/*
WORKDIR /harness
COPY Gemfile Gemfile.lock* ./
RUN bundle install --jobs 4
