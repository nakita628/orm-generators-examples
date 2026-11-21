#!/usr/bin/env bash
# Holds a committed copy of generated files to a fresh generation.
#
#   check.sh [--loose-png] <fresh dir> <committed dir> [<path>...]
#
# Both directories are in runfiles: the fresh one is a generation target's output, the committed
# one the source tree. With paths, only those (files or directories, relative to both) are
# compared, for an example whose generated files sit beside its inputs; without, the two whole
# trees are.
#
# With --loose-png a PNG only has to be a PNG of the same width and height: a rasterizer such as
# resvg draws text with the platform's own build and fonts, so the pixels of the same SVG differ
# from one machine to the next. Every other file has to match byte for byte.
set -euo pipefail

loose_png=false
if [[ "${1:-}" == --loose-png ]]; then
  loose_png=true
  shift
fi
fresh=$1
committed=$2
shift 2
paths=("$@")
[[ ${#paths[@]} -gt 0 ]] || paths=(.)

stale() {
  echo >&2
  echo "the committed copy under $committed is stale: run 'bazel run' on the .write target" >&2
  exit 1
}

status=0
for path in "${paths[@]}"; do
  if [[ ! -e "$committed/$path" ]]; then
    echo "missing from the committed copy: $committed/$path" >&2
    status=1
  elif [[ ! -e "$fresh/$path" ]]; then
    echo "committed but no longer generated: $committed/$path" >&2
    status=1
  elif ! $loose_png; then
    diff -r "$committed/$path" "$fresh/$path" || status=1
  else
    diff -r -x '*.png' "$committed/$path" "$fresh/$path" || status=1
    # The same PNGs on both sides, each with the signature and IHDR (width, height, depth,
    # color type) of the other.
    if ! diff <(cd "$committed/$path" 2>/dev/null && find . -name '*.png' | sort) \
              <(cd "$fresh/$path" 2>/dev/null && find . -name '*.png' | sort); then
      status=1
    fi
    while IFS= read -r png; do
      [[ -f "$fresh/$path/$png" ]] || continue
      if ! cmp -s -n 26 "$committed/$path/$png" "$fresh/$path/$png"; then
        echo "PNG header (size or format) differs: $committed/$path/$png" >&2
        status=1
      fi
    done < <(cd "$committed/$path" 2>/dev/null && find . -name '*.png')
  fi
done
[[ $status -eq 0 ]] || stale
