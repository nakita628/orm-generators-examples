#!/usr/bin/env bash
# Runs run.py with the Python of apps/djangoer/.python-version, which the uv of MODULE.bazel
# installs.
#
#   run.sh <check|write> <uv> <run.py> <example dir>
set -euo pipefail
mode=$1
uv=$(realpath "$2")
run_py=$3
example=$4
python=$(cat "$(dirname "$run_py")/../.python-version")
exec "$uv" run --no-project --quiet --python "$python" --python-preference only-managed \
  python "$run_py" "$mode" "$uv" "$example"
