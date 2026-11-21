"""Runs Djangoer on one example and holds its committed output to the result, or writes it.

    run.py check <uv> <example dir>     # bazel test //apps/djangoer/examples/<name>:generate.test
    run.py write <uv> <example dir>     # bazel run //apps/djangoer/examples/<name>:generate.write

Each example is a standalone uv project (pyproject.toml, uv.lock) with a settings.py that installs
the project's apps. `[tool.example]` in its pyproject.toml says what the docs cover:

- `all = true`: a library installed from PyPI. Djangoer runs with `--all`, and the source of every
  model on the diagram is copied into models/ (scripts/collect_models.py).
- otherwise a sample written here: its apps are directories of the example, and Djangoer runs
  without `--all`, so that they are the project's own and Django's models are left out.
- `settings`, `env`, `package_paths`: the settings module (default `settings`), environment
  variables, and installed packages whose directory goes on PYTHONPATH (Label Studio's `core`).

The example's inputs are copied into a temporary directory and Djangoer writes er.md,
schema.dbml, er.svg and docs/ (and models/) under another; tools/golden then compares or copies
them. The environment is installed with `uv run --locked` (the uv.lock has to be up to date) into
a cache directory per example, so a second run installs nothing.
"""

from __future__ import annotations

import os
import shutil
import subprocess
import sys
import tempfile
import tomllib
from pathlib import Path

OUTPUTS = ["er.md", "schema.dbml", "er.svg", "docs"]
NOT_INPUTS = {*OUTPUTS, "models", "BUILD.bazel", ".venv", "src", "data", "__pycache__"}


def main(mode: str, uv: str, example: Path) -> int:
    here = Path(__file__).resolve().parent
    djangoer_root = example.parent.parent
    golden = here.parent.parent.parent / "tools" / "golden" / f"{mode}.sh"
    collect_models = djangoer_root / "scripts" / "collect_models.py"
    name = example.name
    config = tomllib.loads((example / "pyproject.toml").read_text())["tool"]["example"]
    settings = config.get("settings", "settings")
    paths = OUTPUTS + (["models"] if config.get("all") else [])

    cache = Path(os.environ.get("DJANGOER_CACHE") or Path.home() / ".cache" / "orm-generators-examples" / "djangoer")
    env = {
        **os.environ,
        "UV_LINK_MODE": "copy",
        "UV_PROJECT_ENVIRONMENT": str(cache / "venvs" / name),
        # The Python of apps/djangoer/.python-version, installed by uv, whatever the machine has.
        "UV_PYTHON": (djangoer_root / ".python-version").read_text().strip(),
        "UV_PYTHON_PREFERENCE": "only-managed",
        "UV_NO_PROGRESS": "1",
    }

    with tempfile.TemporaryDirectory(prefix=f"djangoer-{name}-", dir=os.environ.get("TEST_TMPDIR")) as tmp:
        work, out = Path(tmp) / name, Path(tmp) / "out"
        out.mkdir()
        # Runfiles are symlinks into the source tree: copy the files themselves, so Python and
        # uv see the example where they can write beside it.
        shutil.copytree(
            example,
            work,
            symlinks=False,
            ignore=lambda d, names: NOT_INPUTS & set(names) if Path(d) == example else set(),
        )

        def uv_run(*args: str, extra_env: dict[str, str] | None = None) -> subprocess.CompletedProcess[str]:
            return subprocess.run(
                [uv, "run", "--locked", "--quiet", "--project", str(work), *args],
                cwd=work,
                env={**env, **(extra_env or {})},
                text=True,
                capture_output=True,
                check=False,
            )

        pythonpath = [str(work)]
        for package in config.get("package_paths", []):
            found = uv_run(
                "python",
                "-c",
                f"import importlib.util, pathlib; print(pathlib.Path(importlib.util.find_spec({package!r}).origin).parent)",
            )
            if found.returncode != 0:
                print(found.stdout + found.stderr, file=sys.stderr)
                return 1
            pythonpath.append(found.stdout.strip())
        run_env = {"PYTHONPATH": os.pathsep.join(pythonpath), **config.get("env", {})}

        options = ["--settings", settings, *(["--all"] if config.get("all") else [])]
        steps = [
            (
                "djangoer er",
                [
                    "djangoer",
                    "er",
                    "-o",
                    str(out / "er.md"),
                    "-o",
                    str(out / "schema.dbml"),
                    "-o",
                    str(out / "er.svg"),
                    *options,
                ],
            ),
            ("djangoer docs", ["djangoer", "docs", "-o", str(out / "docs"), *options]),
        ]
        if config.get("all"):
            steps.append(("models", ["python", str(collect_models), str(out / "models")]))
        for label, command in steps:
            extra = {**run_env, "DJANGO_SETTINGS_MODULE": settings} if label == "models" else run_env
            result = uv_run(*command, extra_env=extra)
            print(f"{name}: {label}: {'ok' if result.returncode == 0 else f'exit {result.returncode}'}")
            if result.returncode != 0:
                print((result.stdout + result.stderr).strip()[-8000:], file=sys.stderr)
                return 1
            if result.stderr.strip():
                print(result.stderr.strip()[-4000:], file=sys.stderr)

        committed = str(example) if mode == "check" else os.path.relpath(example, djangoer_root.parent.parent)
        return subprocess.run(["bash", str(golden), str(out), committed, *paths], check=False).returncode


if __name__ == "__main__":
    if len(sys.argv) != 4 or sys.argv[1] not in ("check", "write"):
        sys.exit(__doc__)
    sys.exit(main(sys.argv[1], str(Path(sys.argv[2]).resolve()), Path(sys.argv[3]).absolute()))
