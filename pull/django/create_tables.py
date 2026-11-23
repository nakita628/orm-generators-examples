"""Creates the tables of every model of a Django project from its model definitions, the way
`migrate --run-syncdb` does for an app without migrations: Django's own schema editor writes the
SQL (tables, many-to-many tables, indexes, constraints, foreign keys) from the models as they are,
whatever the project's migrations say.

    python create_tables.py     # with DJANGO_SETTINGS_MODULE set, in the project's environment

Run by pull.mjs, with a settings module that points DATABASES at the database to fill. It prints
the tables it created, one per line.
"""

from __future__ import annotations

import sys

import django
from django.apps import apps
from django.db import connection


def main() -> int:
    django.setup()
    models = [
        m
        for m in apps.get_models(include_auto_created=False)
        if m._meta.managed and not m._meta.proxy and not m._meta.swapped
    ]
    created: list[str] = []
    with connection.schema_editor() as editor:
        for model in models:
            editor.create_model(model)
            created.append(model._meta.db_table)
            created.extend(
                f.remote_field.through._meta.db_table
                for f in model._meta.local_many_to_many
                if f.remote_field.through._meta.auto_created
            )
    print("\n".join(created))
    return 0


if __name__ == "__main__":
    sys.exit(main())
