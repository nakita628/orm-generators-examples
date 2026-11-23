#!/usr/bin/env bash
# Makes Docker usable for compat inside a Claude Code cloud session; on a machine of your own,
# skip this and use your Docker as it is.
#
#   bash apps/hekireki/compat/cloud/prepare.sh [orm ...]   # default: every ORM with a Dockerfile
#
# 1. Starts dockerd, with mirror.gcr.io in front of Docker Hub (which rate-limits the session's
#    shared egress).
# 2. Re-tags each Dockerfile's base image with the session proxy's CA in it (base.Dockerfile).
# 3. Builds each compat image with the host's network, where the proxy is.
#
# It ends by printing the COMPAT_DOCKER_BUILD_ARGS to export: run.mjs and pull.mjs pass them to
# their own `docker build`, which then finds every layer in the cache.
set -euo pipefail
here=$(cd "$(dirname "$0")" && pwd)
compat=$(dirname "$here")
ca=${CCR_CA_BUNDLE:-/root/.ccr/ca-bundle.crt}
proxy=${HTTPS_PROXY:-${https_proxy:-}}
[ -f "$ca" ] && [ -n "$proxy" ] || { echo "not a cloud session (no $ca or HTTPS_PROXY): use your Docker as it is" >&2; exit 1; }

if ! docker info >/dev/null 2>&1; then
  if [ ! -f /etc/docker/daemon.json ]; then
    mkdir -p /etc/docker
    echo '{"registry-mirrors": ["https://mirror.gcr.io"]}' > /etc/docker/daemon.json
  fi
  (dockerd > /tmp/dockerd.log 2>&1 &)
  for _ in $(seq 1 60); do docker info >/dev/null 2>&1 && break; sleep 1; done
  docker info >/dev/null 2>&1 || { tail -20 /tmp/dockerd.log >&2; exit 1; }
fi

orms=("$@")
[ ${#orms[@]} -gt 0 ] || mapfile -t orms < <(cd "$compat" && ls -d */Dockerfile | cut -d/ -f1)
port=${proxy##*:}
stage=$(mktemp -d)
trap 'rm -rf "$stage"' EXIT
cp "$ca" "$stage/ccr-ca.crt"
cp "$here/base.Dockerfile" "$here/maven-mirror.gradle" "$stage/"
build_args=(--network host --build-arg "HTTPS_PROXY=$proxy" --build-arg "https_proxy=$proxy"
  --build-arg "NO_PROXY=${NO_PROXY:-localhost,127.0.0.1}" --build-arg "no_proxy=${NO_PROXY:-localhost,127.0.0.1}")

for orm in "${orms[@]}"; do
  dockerfile="$compat/$orm/Dockerfile"
  for base in $(grep -E '^FROM ' "$dockerfile" | awk '{print $2}') $(grep -oE 'COPY --from=[^ ]+' "$dockerfile" | cut -d= -f2); do
    if [ "$(docker image inspect -f '{{index .Config.Labels "ccr-ca-patched"}}' "$base" 2>/dev/null)" = 1 ]; then continue; fi
    docker pull -q "$base" >/dev/null
    user=$(docker image inspect -f '{{.Config.User}}' "$base")
    echo "patching $base" >&2
    docker build -q -f "$stage/base.Dockerfile" --build-arg "BASE=$base" --build-arg "ORIGINAL_USER=${user:-root}" \
      --build-arg "JAVA_PROXY_OPTS=-Dhttps.proxyHost=127.0.0.1 -Dhttps.proxyPort=$port -Dhttp.nonProxyHosts=localhost|127.0.0.1 -Dorg.gradle.internal.repository.max.retries=8 -Dorg.gradle.internal.repository.initial.backoff=2000" \
      -t "$base" "$stage" >/dev/null
  done
  echo "building hekireki-compat-$orm" >&2
  log="$stage/$orm.log"
  if ! docker build --progress=plain "${build_args[@]}" -t "hekireki-compat-$orm" "$compat/$orm" > "$log" 2>&1; then
    grep -v '^#[0-9]* \(sha256\|extracting\|resolve\|DONE\|CACHED\)' "$log" | tail -30 >&2
    echo "hekireki-compat-$orm did not build" >&2
    exit 1
  fi
done
echo "ready: ${orms[*]}" >&2
echo "export COMPAT_DOCKER_BUILD_ARGS='${build_args[*]}'"
