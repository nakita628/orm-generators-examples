"""The targets of one drizzler example (apps/drizzler/examples/<name>)."""

load("@aspect_rules_js//js:defs.bzl", "js_run_binary", "js_test")
load("@npm//:defs.bzl", "npm_link_all_packages")
load("//tools/golden:defs.bzl", "golden")

def drizzler_example(schema = None, dbml_known_issue = None, typecheck = True):
    """Declares the targets of the example in this package.

    - `:node_modules`: the packages of the example's package.json, linked by rules_js.
    - `:docs`: drizzler's output (README.md, er.svg, er.png, er-dark.png, tables/*.svg,
      schema.dbml, er.md), a directory in bazel-bin.
    - `:docs.write` / `:docs.test`: write it into docs/, or hold docs/ to it.
    - `:typecheck`: tsc on the schema, against the drizzle-orm of the lockfile.
    - `:dbml`: @dbml/core parses the generated schema.dbml.

    Args:
        schema: The schema file to pass drizzler, for an example whose drizzle.config.ts does
            not point at it; by default drizzler reads drizzle.config.ts.
        dbml_known_issue: A drizzler issue (URL) that makes the DBML fail to parse; `:dbml` then
            has to fail, so it notices the fix.
        typecheck: False for an example whose upstream schema does not type-check on its own.
    """
    npm_link_all_packages(name = "node_modules")

    srcs = native.glob(
        ["**"],
        exclude = [
            "BUILD.bazel",
            "README.md",
            "docs/**",
            "node_modules/**",
        ],
    )

    js_run_binary(
        name = "docs",
        srcs = srcs + [":node_modules"],
        args = ["_docs"] + ([schema] if schema else []),
        chdir = native.package_name(),
        out_dirs = ["_docs"],
        tool = "//apps/drizzler/tools:generate",
        mnemonic = "Drizzler",
        progress_message = "Generating the drizzler docs of %s" % native.package_name(),
    )

    golden(
        name = "docs",
        generated = ":docs",
        committed = "docs",
        # er.png and er-dark.png are rasterized by resvg with the platform's fonts.
        loose_png = True,
    )

    if typecheck:
        js_test(
            name = "typecheck",
            chdir = native.package_name(),
            data = srcs + [":node_modules"],
            entry_point = "//apps/drizzler/tools:tsc",
        )

    js_test(
        name = "dbml",
        args = ["$(rootpath :docs)/schema.dbml"] + (["--expect-failure"] if dbml_known_issue else []),
        data = [
            ":docs",
            "//apps/drizzler:node_modules/@dbml/core",
        ],
        entry_point = "//apps/drizzler/tools:check_dbml",
    )
