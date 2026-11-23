# A base image of a compat Dockerfile (ruby:3.4-slim, gradle:8.14-jdk21, ...) made to build inside a
# Claude Code cloud session, re-tagged under its own name so that the Dockerfiles stay as they are.
# Only prepare.sh builds it; on a machine of your own the original images are used.
#
# The session reaches the internet through an HTTPS proxy that re-signs TLS with its own CA, and
# only over HTTPS: the CA is added to the system store (and to Java's), Debian's apt sources move
# from http:// to https://, Java gets the proxy as system properties (it ignores HTTPS_PROXY), and
# Gradle fetches Maven Central from Google's mirror of it (maven-mirror.gradle), with more retries.
ARG BASE
FROM ${BASE}
ARG ORIGINAL_USER=root
ARG JAVA_PROXY_OPTS=""
USER root
COPY ccr-ca.crt /usr/local/share/ca-certificates/ccr-agent-proxy.crt
COPY maven-mirror.gradle /tmp/ccr-maven-mirror.gradle
# An image without the ca-certificates package (elixir:*-slim) gets the CA as its whole bundle; its
# Dockerfile's own `apt-get install ca-certificates` then rebuilds the bundle with it included.
RUN set -e; \
    mkdir -p /etc/ssl/certs; \
    if command -v update-ca-certificates >/dev/null 2>&1; then update-ca-certificates >/dev/null 2>&1; \
    else cat /usr/local/share/ca-certificates/ccr-agent-proxy.crt >> /etc/ssl/certs/ca-certificates.crt; fi; \
    if [ -n "${JAVA_HOME:-}" ] && [ -x "$JAVA_HOME/bin/keytool" ]; then \
      "$JAVA_HOME/bin/keytool" -importcert -noprompt -cacerts -storepass changeit -alias ccr-agent-proxy \
        -file /usr/local/share/ca-certificates/ccr-agent-proxy.crt >/dev/null; fi; \
    for f in /etc/apt/sources.list /etc/apt/sources.list.d/*; do \
      if [ -f "$f" ]; then sed -i 's|http://deb.debian.org|https://deb.debian.org|g' "$f"; fi; done; \
    if [ -n "${GRADLE_HOME:-}" ]; then mkdir -p "$GRADLE_HOME/init.d" && cp /tmp/ccr-maven-mirror.gradle "$GRADLE_HOME/init.d/"; fi; \
    rm -f /tmp/ccr-maven-mirror.gradle; \
    mkdir -p /root/.composer && printf '{"config": {"preferred-install": "source"}}\n' > /root/.composer/config.json
# Composer: GitHub's dist zips go through the session's per-repository GitHub access, which a
# package's repository is not part of; a git clone of a public repository goes through.
ENV COMPOSER_HOME=/root/.composer
ENV SSL_CERT_FILE=/etc/ssl/certs/ca-certificates.crt \
    CURL_CA_BUNDLE=/etc/ssl/certs/ca-certificates.crt \
    NODE_EXTRA_CA_CERTS=/usr/local/share/ca-certificates/ccr-agent-proxy.crt \
    HEX_CACERTS_PATH=/etc/ssl/certs/ca-certificates.crt \
    JAVA_TOOL_OPTIONS=${JAVA_PROXY_OPTS}
LABEL ccr-ca-patched=1
USER ${ORIGINAL_USER}
