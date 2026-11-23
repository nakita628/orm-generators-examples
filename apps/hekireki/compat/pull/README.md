# compat/pull

Schemas that other ORMs wrote, read back into Prisma, for [pull.mjs](../pull.mjs): what it does and
how to run it is in the "Pulled schemas" section of the [compat README](../README.md#pulled-schemas-other-orms-tables-read-back-into-prisma).

**[Results](results/README.md)**

| Path | |
| --- | --- |
| [sources.json](sources.json) | the sources that are not examples of this repository (Rails apps, MLflow), and `skip`: the ones that cannot be pulled, with the reason |
| `schemas/<kind>-<name>/` | per source: `source.sql`, `schema.prisma`, `prisma-diff.sql`, `source.json` (committed) and `generated/` (not) |
| `results/` | what run.mjs found, per source, and the matrix |
| [django/create_tables.py](django/create_tables.py) | Django's schema editor over every model of a djangoer example |
| [activerecord/load_schema.rb](activerecord/load_schema.rb) | a Rails app's db/schema.rb, loaded by Active Record alone |
| [sqlalchemy/create_all.py](sqlalchemy/create_all.py) | `metadata.create_all()` over a SQLAlchemy project's models |

Every finding in `results/` is about what hekireki generated, held to the tables in `source.sql`.
What `prisma db pull` itself cannot express is in each `prisma-diff.sql` and in the header of each
`schema.prisma`; it is Prisma's, not hekireki's.
