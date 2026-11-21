"""Holds the generated files an example commits to a fresh generation.

Every example of every app (hekireki, drizzler, djangoer) commits what its generator wrote, so
the output can be read on GitHub and reviewed in a diff. `golden` adds the two targets that keep
that copy honest.
"""

load("@rules_shell//shell:sh_binary.bzl", "sh_binary")
load("@rules_shell//shell:sh_test.bzl", "sh_test")

def golden(name, generated, committed, paths = [], loose_png = False, **kwargs):
    """Declares `<name>.write` and `<name>.test` for a directory a generation target writes.

    - `bazel run :<name>.write` copies the generation into the source tree.
    - `bazel test :<name>.test` fails when the committed copy differs from a fresh generation.

    Args:
        name: The generation target's name; the targets are `<name>.write` and `<name>.test`.
        generated: The label of the generation, a directory (a tree artifact).
        committed: The committed directory, relative to this package.
        paths: Only these files and directories, relative to both, are compared and written; by
            default the whole directory is. For an example whose outputs sit beside its inputs.
        loose_png: A PNG only has to match in format and size, not in pixels: for a rasterizer
            whose output depends on the platform's fonts (see check.sh).
        **kwargs: Passed to both targets (`tags`, `visibility`, ...).
    """
    prefix = "" if committed in ("", ".") else committed + "/"
    args = ["$(rootpath %s)" % generated, native.package_name() + "/" + prefix.rstrip("/")] + paths
    if paths:
        committed_files = native.glob([prefix + p for p in paths] + [prefix + p + "/**" for p in paths], allow_empty = True)
    else:
        committed_files = native.glob([prefix + "**"], allow_empty = True)
    sh_binary(
        name = name + ".write",
        srcs = ["//tools/golden:write.sh"],
        args = (["--loose-png"] if loose_png else []) + args,
        data = [generated],
        **kwargs
    )
    sh_test(
        name = name + ".test",
        srcs = ["//tools/golden:check.sh"],
        args = (["--loose-png"] if loose_png else []) + args,
        data = [generated] + committed_files,
        **kwargs
    )
