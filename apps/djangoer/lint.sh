#!/usr/bin/env bash
# ruff check, ruff format --check on the scripts, and readme.py --check. Run by //apps/djangoer:lint.
set -euo pipefail
uv=$(realpath "$1")
cd apps/djangoer
work=$(mktemp -d "${TEST_TMPDIR:-/tmp}/djangoer-lint.XXXXXX")
cp -RL . "$work"
cd "$work"
export UV_PROJECT_ENVIRONMENT=${DJANGOER_CACHE:-$HOME/.cache/orm-generators-examples/djangoer}/venvs/_tooling
export UV_PYTHON_PREFERENCE=only-managed UV_NO_PROGRESS=1
"$uv" run --locked --quiet ruff check .
"$uv" run --locked --quiet ruff format --check scripts tools
"$uv" run --locked --quiet python scripts/readme.py --check
