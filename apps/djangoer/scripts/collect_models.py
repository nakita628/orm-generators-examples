"""Copy the source of every model of a Django project into a directory: the files that define
them and the abstract models they inherit their fields from, at their path in site-packages.

    python scripts/collect_models.py DEST    # with DJANGO_SETTINGS_MODULE set, in the project's env

Run by generate.py. It prints the files it copied, relative to DEST.
"""

from __future__ import annotations

import inspect
import shutil
import sys
import sysconfig
from pathlib import Path

import django
from django.apps import apps
from django.db import models


def root_of(file: Path) -> Path:
    """The directory a file was installed in: site-packages first, so that a project that puts its
    own package on sys.path (Label Studio's `core`…) keeps its package's name, else the entry of
    sys.path it is in."""
    site = Path(sysconfig.get_paths()["purelib"]).resolve()
    if site in file.parents:
        return site
    roots = {Path(p).resolve() for p in sys.path if p}
    for root in sorted(roots, key=lambda r: -len(r.parts)):
        if root in file.parents:
            return root
    raise ValueError(f"{file} is not under sys.path")


def main(dest: Path) -> None:
    django.setup()
    files: set[Path] = set()
    for model in apps.get_models(include_auto_created=False):
        for cls in model.__mro__:
            if not (isinstance(cls, type) and issubclass(cls, models.Model)) or cls is models.Model:
                continue
            try:
                source = inspect.getsourcefile(cls)
            except TypeError:
                continue
            if source:
                files.add(Path(source).resolve())
    if dest.exists():
        shutil.rmtree(dest)
    for file in sorted(files):
        relative = file.relative_to(root_of(file))
        target = dest / relative
        target.parent.mkdir(parents=True, exist_ok=True)
        shutil.copyfile(file, target)
        print(relative)


if __name__ == "__main__":
    main(Path(sys.argv[1]).resolve())
