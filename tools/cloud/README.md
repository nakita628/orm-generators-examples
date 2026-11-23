# tools/cloud

Docker inside a Claude Code cloud session, for
[apps/hekireki/compat](../../apps/hekireki/compat) and [pull/](../../pull), which runs it. On a
machine of your own, none of this is needed: use your Docker as it is.

```sh
eval "$(bash tools/cloud/prepare.sh | tail -1)"
```

[prepare.sh](prepare.sh):

1. starts dockerd, with mirror.gcr.io in front of Docker Hub (which rate-limits the session's
   shared egress);
2. re-tags each compat Dockerfile's base image with the session proxy's CA in it
   ([base.Dockerfile](base.Dockerfile)): the CA in the system store and Java's, Debian's apt
   sources over https, Java's proxy settings, Maven Central from Google's mirror of it
   ([maven-mirror.gradle](maven-mirror.gradle)), and Composer installing from git;
3. builds every compat image with the host's network, where the proxy is;
4. prints `COMPAT_DOCKER_BUILD_ARGS`, which compat's run.mjs and pull/pull.mjs add to their own
   `docker build`, so that it finds every layer in the cache.

The compat Dockerfiles stay as they are. The environment has to allow `deb.debian.org`,
`repo.hex.pm` and `builds.hex.pm`.
