# Rust with the crates of Cargo.toml compiled: the SeaORM harness builds at run time with the
# generated entities in src/models, into a target directory kept in a Docker volume.
FROM rust:1-slim
RUN apt-get update && apt-get install -y --no-install-recommends pkg-config && rm -rf /var/lib/apt/lists/*
WORKDIR /harness
COPY Cargo.toml Cargo.lock ./
RUN mkdir -p src && echo 'fn main() {}' > src/main.rs && cargo build -j 2 && rm -rf src
