"""The targets of one Djangoer example (apps/djangoer/examples/<name>)."""

load("@rules_shell//shell:sh_binary.bzl", "sh_binary")
load("@rules_shell//shell:sh_test.bzl", "sh_test")

# What Djangoer writes into the example; everything else in it is an input.
OUTPUTS = ["er.md", "schema.dbml", "er.svg", "docs/**", "models/**"]

def djangoer_example(name = "generate", tags = []):
    """Declares `<name>.test` and `<name>.write` for the example in this package.

    - `bazel test :generate.test` installs the example's locked environment with uv, runs
      Djangoer in it, and fails when er.md, schema.dbml, er.svg, docs/ (and models/) differ from
      what it wrote.
    - `bazel run :generate.write` writes them into the source tree.

    The environment comes from PyPI, so both run outside the sandbox with the network, and the
    test is tagged `djangoer` (left out of `bazel test //...`; see .bazelrc). See
    //apps/djangoer/tools:run.py.

    Args:
        name: Prefix of the targets.
        tags: More tags for both.
    """
    inputs = native.glob(
        ["**"],
        exclude = OUTPUTS + ["BUILD.bazel", ".venv/**", "src/**", "data/**", "**/__pycache__/**"],
    )
    data = inputs + [
        "//apps/djangoer:.python-version",
        "//apps/djangoer/scripts:collect_models.py",
        "//apps/djangoer/tools:run.py",
        "//tools/golden:check.sh",
        "//tools/golden:write.sh",
        "//tools/uv",
    ]
    args = [
        "$(rootpath //tools/uv)",
        "$(rootpath //apps/djangoer/tools:run.py)",
        native.package_name(),
    ]
    tags = tags + ["djangoer", "no-sandbox", "requires-network"]

    # For //apps/djangoer:lint: the example's own Python (settings.py, a sample's models.py) for
    # ruff, and docs/README.md and models/, which scripts/readme.py links when they are there.
    native.filegroup(
        name = "lint_inputs",
        srcs = native.glob(
            ["**/*.py", "docs/README.md", "models/**"],
            exclude = ["docs/**/*.py", "**/__pycache__/**"],
            allow_empty = True,
        ),
        visibility = ["//apps/djangoer:__pkg__"],
    )
    sh_test(
        name = name + ".test",
        size = "large",
        timeout = "eternal",
        srcs = ["//apps/djangoer/tools:run.sh"],
        args = ["check"] + args,
        data = data + native.glob(OUTPUTS, allow_empty = True),
        tags = tags,
    )
    sh_binary(
        name = name + ".write",
        srcs = ["//apps/djangoer/tools:run.sh"],
        args = ["write"] + args,
        data = data,
        tags = tags,
    )
