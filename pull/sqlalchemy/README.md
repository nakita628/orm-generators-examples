# pull/sqlalchemy

The tables SQLAlchemy projects create, written up in Prisma. **[Results](results/README.md)**

| Project | How the tables are made | License |
| --- | --- | --- |
| [MLflow](https://github.com/mlflow/mlflow) | `mlflow db upgrade` (its SQLAlchemy models with their Alembic migrations) | Apache-2.0 |

Each is a package from PyPI, pinned in [sources.json](sources.json). pull.mjs runs it with uv,
with psycopg and PyMySQL beside it, either by the project's own command (`command`, with `{url}`
for the database) or by [create_all.py](create_all.py), which imports the modules that define the
models (`modules`) and runs their declarative base's `metadata.create_all()` (`base`).

MLflow's models alone do not make a schema PostgreSQL accepts (a foreign key to a column without
a unique constraint), so its own command is used, as its users do.

Each project's own `README.md` holds the SQL every ORM creates from the write-up (`sql/`) against
`source.sql`; [results/](results/README.md) has the layer's matrices.
