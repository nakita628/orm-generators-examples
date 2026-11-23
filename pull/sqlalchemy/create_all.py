"""Creates the tables of a SQLAlchemy project from its model definitions: the modules that define
the models are imported, and their declarative base's metadata.create_all() writes the SQL.

    python create_all.py <database url> <module:Base> [module ...]

Run by pull.mjs with uv, in an environment holding the project's package from PyPI and the
database driver. It prints the tables it created, one per line.
"""

from __future__ import annotations

import importlib
import sys

from sqlalchemy import create_engine


def main(url: str, base: str, modules: list[str]) -> int:
    for module in modules:
        importlib.import_module(module)
    module, _, attr = base.partition(":")
    metadata = getattr(importlib.import_module(module), attr).metadata
    if url.startswith("postgresql://"):
        url = "postgresql+psycopg://" + url[len("postgresql://") :]
    elif url.startswith("mysql://"):
        url = "mysql+pymysql://" + url[len("mysql://") :]
    elif url.startswith("file:"):
        url = "sqlite:///" + url[len("file:") :]
    engine = create_engine(url)
    metadata.create_all(engine)
    print("\n".join(sorted(metadata.tables)))
    return 0


if __name__ == "__main__":
    if len(sys.argv) < 3:
        sys.exit(__doc__)
    sys.exit(main(sys.argv[1], sys.argv[2], sys.argv[3:]))
