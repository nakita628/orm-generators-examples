# pull

The SQL that other ORMs generate, written up in Prisma, and the SQL every ORM hekireki generates
from that write-up, held to it. A well-known project's own ORM creates its tables from the
project's definitions (Django models, a Drizzle schema, a Rails `db/schema.rb`, ...), the tables
are kept as SQL, and `prisma db pull` writes them up as a Prisma schema. Then hekireki generates
every ORM's code from that write-up, and:

- **the SQL**: Prisma and each ORM that creates tables create them on their own from the
  write-up, and their SQL is held to the project's own: does Django's SQL from hekireki's Django
  models come out as Django's SQL from the project's models? Drizzle's? SQLAlchemy's, GORM's,
  SeaORM's, EF Core's, Exposed's?
- **the runtime**: the code of every ORM runs against the project's own tables: whether what
  hekireki writes from a schema nobody wrote for Prisma is valid for each ORM, and works.

| Layer | ORM that writes the SQL | Projects | How the tables are made | Results |
| --- | --- | --: | --- | --- |
| [django](django) | Django | every [djangoer example](../apps/djangoer/examples) | Django's schema editor over every model ([create_tables.py](django/create_tables.py)), on PostgreSQL | [results](django/results/README.md) |
| [drizzle](drizzle) | Drizzle | every [drizzler example](../apps/drizzler/examples) | `drizzle-kit export`, in the example's dialect | [results](drizzle/results/README.md) |
| [activerecord](activerecord) | Active Record | Mastodon, Forem, Chatwoot, Lobsters, Maybe, Loomio, Campfire ([sources.json](activerecord/sources.json)) | Active Record loads the app's `db/schema.rb` at a pinned commit ([load_schema.rb](activerecord/load_schema.rb)) | [results](activerecord/results/README.md) |
| [sqlalchemy](sqlalchemy) | SQLAlchemy | MLflow ([sources.json](sqlalchemy/sources.json)) | the project's own command (`mlflow db upgrade`), or `metadata.create_all()` ([create_all.py](sqlalchemy/create_all.py)) | [results](sqlalchemy/results/README.md) |

## Layout

```text
pull/
  pull.mjs                  the whole run: tables, SQL, write-up, hekireki, the SQL, the runtime
  lib/sql.mjs               holds one set of tables' SQL to another
  <orm>/
    README.md               the layer's projects and how their tables are made
    sources.json            projects that are not examples of this repository, and `skip`
    <project>/
      README.md             every ORM's SQL against source.sql, with what differs
      source.sql            the SQL of the tables the project's own ORM created
      schema.prisma         `prisma db pull` of those tables: the write-up
      sql/
        prisma.sql          the tables `prisma db push` creates from the write-up
        <orm>.sql           the tables each ORM creates from hekireki's code for it
        <orm>.diff          the whole diff against source.sql, beside each one that differs
      source.json           where it came from, the versions, when it was checked, the comparison
      generated/            hekireki's output from the write-up (not committed)
    results/                the runtime checks, and the layer's matrix of both
```

- `source.sql` and every `sql/<orm>.sql` are dumped the same way (pg_dump, mysqldump, or SQLite's
  own schema), without pg_dump's comment lines and what changes from run to run, so that one reads
  against the other line for line.
- `schema.prisma` is what `prisma db pull` wrote. Where Prisma writes a schema it refuses itself
  (an integer foreign key on a bigint id, a default on an `Unsupported` field, a Gin index over
  column types Prisma does not take together, an SQLite expression default without the
  parentheses `prisma db push` needs, an index called `sqlite_autoindex_*`, a name SQLite
  reserves), pull.mjs makes the smallest change that validates and lists it in the header and in
  `source.json`'s `fixes`.

## The SQL

| SQL | From the write-up, by |
| --- | --- |
| `prisma.sql` | `prisma db push` |
| `drizzle.sql` | `drizzle-kit export` of hekireki-drizzle's schema |
| `sqlalchemy.sql` | `metadata.create_all()` of hekireki-sqlalchemy's models |
| `django.sql` | `makemigrations` and `migrate` of hekireki-django's models |
| `gorm.sql` | `AutoMigrate` of hekireki-gorm's models |
| `sea-orm.sql` | `Schema::create_table_from_entity` of hekireki-sea-orm's entities |
| `efcore.sql` | `EnsureCreated()` of hekireki-efcore's context |
| `exposed.sql` | `SchemaUtils.create` of hekireki-exposed's tables (PostgreSQL only) |

Kysely (types only), Ecto (schemas only), Active Record and Eloquent (models only) create no
tables. Each SQL is held to `source.sql` ([lib/sql.mjs](lib/sql.mjs)), once what is not a table is
left out (session settings, comments, extensions):

- `=` **identical**: the same statements in the same order;
- `≈` **the same but for names**: the same statements once constraint and index names are set
  aside and the order of columns and statements does not count;
- `≠` the rest: each project's README shows the tables written differently, column by column, and
  the statements only one side has, and `sql/<orm>.diff` is the whole diff.

Each ORM's SQL is also held to `prisma.sql`, what Prisma itself creates from the same write-up. A
difference from `prisma.sql` is hekireki's or the ORM's (Django makes every foreign key
`DEFERRABLE INITIALLY DEFERRED`, for one); a difference from `source.sql` alone is something the
write-up does not say (a check constraint, a trigger), which is Prisma's.

## The runtime

For each project, every hekireki generator runs on the write-up (one at a time, so that one
failing does not hide the others), and
[apps/hekireki/compat](../apps/hekireki/compat)'s harnesses run the generated code against a
database made from `source.sql`, not from `prisma db push`: for each ORM (Drizzle, Kysely,
SQLAlchemy, Django, GORM, SeaORM, Ecto, Active Record, Eloquent, EF Core, Exposed), every table has
a model and every column an attribute, a row goes in through the ORM, comes back as written, is
updated, joins through every relation and goes out. A generator that fails leaves
`generated/<orm>.error`, shown in the results.

## Usage

Docker, Node, [uv](https://docs.astral.sh/uv/) and Go are needed, as for compat, and
`pnpm install` at the root (drizzle-kit, Prisma and hekireki come from the workspace).

```sh
node pull/pull.mjs --list                                  # every source, and its last check
node pull/pull.mjs django/wagtail activerecord/mastodon    # write up these, and check the SQL and the runtime
node pull/pull.mjs --random 3                              # three: never pulled first, then the oldest
node pull/pull.mjs --pull-only drizzle/northwind           # only the write-up (and prisma.sql)
node pull/pull.mjs --orm django,activerecord django/wagtail  # only these ORMs
node pull/pull.mjs --compare-only django/wagtail           # hold the SQL in sql/ to source.sql again
node pull/pull.mjs --hekireki latest django/wagtail        # with another hekireki release
```

`--hekireki <version>` installs that release in the cache and checks the committed write-ups
with it, leaving pull/ as it is: whether a finding is still there in a newer release than the one
apps/hekireki pins. In a Claude Code cloud session, run [tools/cloud](../tools/cloud) first.

A scheduled task writes up and checks three projects a day, and files what is hekireki's on
[hekireki](https://github.com/nakita628/hekireki/issues): see
[.claude/skills/pull-check](../.claude/skills/pull-check/SKILL.md).

## License

Each project's tables (`source.sql`, `sql/`) and their write-up (`schema.prisma`) come from that
project's definitions, under that project's license (`source.json`, or the example's README in
apps/).
