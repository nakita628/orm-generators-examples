# Elixir 1.18 with the deps of mix.lock compiled: the Ecto harness runs in this image (postgrex
# needs Elixir 1.15+). At run time only the generated schemas, in lib/generated, are compiled.
FROM elixir:1.18-slim
RUN apt-get update && apt-get install -y --no-install-recommends build-essential git ca-certificates && rm -rf /var/lib/apt/lists/*
WORKDIR /harness
ENV MIX_ENV=dev
RUN mix local.hex --force && mix local.rebar --force
COPY mix.exs mix.lock ./
RUN mix deps.get && mix deps.compile
COPY lib lib
COPY harness.exs .
RUN mix compile
