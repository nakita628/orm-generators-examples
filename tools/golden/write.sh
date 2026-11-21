#!/usr/bin/env bash
# Copies a fresh generation over the committed copy in the source tree.
#
#   write.sh [--loose-png] <fresh dir> <committed dir> [<path>...]
#
# The fresh directory is in runfiles; the committed one is relative to the workspace root (run
# through `bazel run`, which sets BUILD_WORKSPACE_DIRECTORY). With paths, only those are
# replaced (each is removed first, so a file the generator no longer writes goes too); without,
# the whole committed directory is.
#
# Bazel leaves its outputs read-only, and `cp -R` would carry that over (and fails outright on a
# file system that refuses to set a directory's mode, as a container's bind mount can); the tree
# is streamed through tar with the modes rewritten to the usual 755 and 644.
#
# With --loose-png a committed PNG whose format and size the fresh one shares is kept: its pixels
# only differ by the machine's fonts (see check.sh), and rewriting it would be churn.
set -euo pipefail

loose_png=false
if [[ "${1:-}" == --loose-png ]]; then
  loose_png=true
  shift
fi
fresh=$1
relative=$2
committed=${BUILD_WORKSPACE_DIRECTORY:?run this through bazel run}/$relative
shift 2
paths=("$@")
[[ ${#paths[@]} -gt 0 ]] || paths=(.)

if [[ ! -d "$fresh" ]]; then
  echo "no generated output at $fresh" >&2
  exit 1
fi

mkdir -p "$committed"
kept=$(mktemp -d)
trap 'rm -rf "$kept"' EXIT
if $loose_png; then
  for path in "${paths[@]}"; do
    [[ -d "$committed/$path" ]] || continue
    while IFS= read -r png; do
      if [[ -f "$fresh/$path/$png" ]] && cmp -s -n 26 "$committed/$path/$png" "$fresh/$path/$png"; then
        mkdir -p "$(dirname "$kept/$path/$png")"
        cp -p "$committed/$path/$png" "$kept/$path/$png"
      fi
    done < <(cd "$committed/$path" && find . -name '*.png')
  done
fi
for path in "${paths[@]}"; do
  if [[ "$path" == . ]]; then
    chmod -R u+w "$committed" 2>/dev/null || true
    find "$committed" -mindepth 1 -delete
  else
    chmod -R u+w "$committed/$path" 2>/dev/null || true
    rm -rf "${committed:?}/$path"
  fi
  if [[ -e "$fresh/$path" ]]; then
    (cd "$fresh" && tar -cf - --mode='a-x,u+rwX,go+rX' "$path") | tar -C "$committed" -xf - --no-same-permissions
  fi
done
if $loose_png; then
  (cd "$kept" && find . -name '*.png' -print0 | xargs -0 -r -I{} cp -p {} "$committed/{}")
fi
echo "wrote $relative"
