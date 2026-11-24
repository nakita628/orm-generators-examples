# tools/upstream

The daily upstream sync: every example follows the project its schema comes from and the newest
release of its generator; what a generator cannot handle becomes an issue on that generator's
repository.

A Claude Code scheduled task runs it once a day in a fresh cloud session, following
[`.claude/skills/upstream-sync/SKILL.md`](../../.claude/skills/upstream-sync/SKILL.md), and pushes
what passes straight to `main`.

| File | |
| --- | --- |
| [`check.mjs`](check.mjs) | reports what is behind: upstream schemas (hekireki's `sources.json`, drizzler's `- Source:` links), djangoer's locked packages, and the generator versions |
| [`follow-drizzler.mjs`](follow-drizzler.mjs) | follows one drizzler example to a newer upstream commit with a 3-way merge per file (the example's changes from upstream are kept) |
| [`blocked.json`](blocked.json) | pins held back, each with the generator issue that holds it (or `null` and a reason, for an upstream problem) |

```bash
node tools/upstream/check.mjs                      # Markdown report
node tools/upstream/check.mjs --app drizzler       # one app: hekireki, drizzler, djangoer, generators
node tools/upstream/check.mjs --name calcom,dub    # some examples
node tools/upstream/check.mjs --json report.json   # also as JSON
node tools/upstream/check.mjs --apply-hekireki     # move sources.json to HEAD and re-import
```

What counts as behind:

- **hekireki / drizzler**: the schema file (or directory) differs between the pinned commit and
  the upstream default branch's HEAD. A new upstream commit that leaves the schema alone is not
  reported. `moved` means the path is gone at HEAD.
- **djangoer**: `uv lock --upgrade` would move a package the example's `pyproject.toml` names (the
  project itself, Django, djangoer…). Transitive-only updates are not counted.
- **generators**: hekireki and drizzler on npm, djangoer on PyPI, against the pinned versions.

`blocked.json` entries:

```json
[
  {
    "app": "hekireki",
    "name": "calcom",
    "target": "<upstream commit, djangoer update list, generator version, or *>",
    "issue": "https://github.com/nakita628/hekireki/issues/123",
    "reason": "one line"
  }
]
```

## Machine setup (cloud sessions)

The cloud environment `orm-generators-examples` allows the hosts Bazel needs (`bcr.bazel.build`,
`releases.bazel.build`, `mirror.bazel.build`, `nodejs.org`, `binaries.prisma.sh`) and its setup
script does the following. Run it by hand if `bazel` is missing:

```bash
npm install -g @bazel/bazelisk
# GitHub archive downloads are limited to the session's repositories; most of the Bazel
# toolchain archives are on mirror.bazel.build.
printf 'rewrite github.com/(.*) mirror.bazel.build/github.com/$1\nrewrite github.com/(.*) github.com/$1\n' > ~/.bazel-downloader.cfg
cat >> ~/.bazelrc <<'EOF'
common --downloader_config=/root/.bazel-downloader.cfg
common --action_env=NODE_EXTRA_CA_CERTS=/root/.ccr/ca-bundle.crt --action_env=SSL_CERT_FILE=/root/.ccr/ca-bundle.crt
common --test_env=NODE_EXTRA_CA_CERTS=/root/.ccr/ca-bundle.crt --test_env=SSL_CERT_FILE=/root/.ccr/ca-bundle.crt
EOF
```

None of this belongs in the repository's `.bazelrc`: it is specific to the cloud sandbox's proxy.
