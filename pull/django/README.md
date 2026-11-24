# pull/django

The tables Django creates from the models of every [djangoer example](../../apps/djangoer/examples),
written up in Prisma. **[Results](results/README.md)**

For `django/<example>`, pull.mjs copies the example (its `pyproject.toml`, `uv.lock` and
`settings.py`), runs it in the example's own locked environment (the one djangoer's Bazel targets
use) with PostgreSQL as the database, and [create_tables.py](create_tables.py) has Django's
schema editor create the table of every model, as `migrate --run-syncdb` does for an app without
migrations: tables, many-to-many tables, indexes, constraints and foreign keys, from the models
as they are, whatever the project's migrations say. The tables are dumped to
`<example>/source.sql` and `prisma db pull` writes `<example>/schema.prisma`.

The examples are listed in apps/djangoer; an example that cannot be pulled (a GIS project needs
PostGIS, for instance) goes in `skip` of a `sources.json` here, with the reason. An example whose
models need what the project's migrations create (a sequence that a `db_default` draws on) gets
a `before_sql` / `after_sql` under `sources` in the same file: Weblate's `Workspace.metric_id`.

Each project's own `README.md` holds the SQL every ORM creates from the write-up (`sql/`) against
`source.sql`; [results/](results/README.md) has the layer's matrices.
