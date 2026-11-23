# pull

The SQL that other ORMs generate, written up in Prisma. A well-known project's own ORM creates its
tables from the project's definitions (Django models, a Drizzle schema, a Rails `db/schema.rb`,
...), the tables are kept as SQL, and `prisma db pull` writes them up as a Prisma schema. Then
hekireki generates every ORM's code from that write-up, and the code runs against the tables the
project's own ORM created: whether what hekireki writes from a schema nobody wrote for Prisma is
valid for each ORM, and works.

| Layer | ORM that writes the SQL | Projects | How the tables are made | Checks |
| --- | --- | --: | --- | --- |
| [django](django) | Django | every [djangoer example](../apps/djangoer/examples) | Django's schema editor over every model ([create_tables.py](django/create_tables.py)), on PostgreSQL | [results](django/results/README.md) |
| [drizzle](drizzle) | Drizzle | every [drizzler example](../apps/drizzler/examples) | `drizzle-kit export`, in the example's dialect | [results](drizzle/results/README.md) |
| [activerecord](activerecord) | Active Record | Mastodon, Forem, Chatwoot, Lobsters, Maybe, Loomio, Campfire ([sources.json](activerecord/sources.json)) | Active Record loads the app's `db/schema.rb` at a pinned commit ([load_schema.rb](activerecord/load_schema.rb)) | [results](activerecord/results/README.md) |
| [sqlalchemy](sqlalchemy) | SQLAlchemy | MLflow ([sources.json](sqlalchemy/sources.json)) | the project's own command (`mlflow db upgrade`), or `metadata.create_all()` ([create_all.py](sqlalchemy/create_all.py)) | [results](sqlalchemy/results/README.md) |

## Layout

```text
pull/
  pull.mjs                  the whole run: tables, SQL, write-up, hekireki, checks
  <orm>/
    README.md               the layer's projects and how their tables are made
    sources.json            projects that are not examples of this repository, and `skip`
    <project>/
      source.sql            the tables as the project's ORM created them
      schema.prisma         `prisma db pull` of those tables: the write-up
      prisma-diff.sql       what the write-up cannot say about the tables (only when there is any)
      source.json           where it came from, the versions, when it was pulled and checked
      generated/            hekireki's output from the write-up (not committed)
    results/                what every ORM's generated code did against the tables
```

- `source.sql` comes from pg_dump, mysqldump or SQLite's own schema, without the lines that change
  from run to run.
- `schema.prisma` is what `prisma db pull` wrote. Where Prisma writes a schema it refuses itself
  (an integer foreign key on a bigint id, a default on an `Unsupported` field), pull.mjs makes the
  smallest change that validates and lists it in the header and in `source.json`'s `fixes`.
- `prisma-diff.sql` is what `prisma migrate diff` would run to turn the tables into the write-up's:
  what a Prisma schema cannot say about them (check constraints, triggers, ...). It is Prisma's,
  not hekireki's.

## The checks

For each project, every hekireki generator runs on the write-up (one at a time, so that one
failing does not hide the others), and
[apps/hekireki/compat](../apps/hekireki/compat)'s harnesses run the generated code against a
database made from `source.sql`, not from `prisma db push`: for each ORM (Drizzle, Kysely,
SQLAlchemy, Django, GORM, SeaORM, Ecto, Active Record, Eloquent, EF Core, Exposed), every table has
a model and every column an attribute, a row goes in through the ORM, comes back as written, is
updated, joins through every relation and goes out; and the tables the ORM creates on its own are
compared with the write-up. A generator that fails leaves `generated/<orm>.error`, shown in the
results.

## Usage

Docker, Node, [uv](https://docs.astral.sh/uv/) and Go are needed, as for compat, and
`pnpm install` at the root (drizzle-kit, Prisma and hekireki come from the workspace).

```sh
node pull/pull.mjs --list                                  # every source, and its last check
node pull/pull.mjs django/wagtail activerecord/mastodon    # write up these and check them
node pull/pull.mjs --random 3                              # three: never pulled first, then the oldest
node pull/pull.mjs --pull-only drizzle/northwind           # only the write-up
node pull/pull.mjs --orm django,activerecord django/wagtail  # check only these ORMs
node pull/pull.mjs --hekireki latest django/wagtail        # check with another hekireki release
```

`--hekireki <version>` installs that release in the cache and checks the committed write-ups
with it, leaving pull/ as it is: whether a finding is still there in a newer release than the one
apps/hekireki pins. In a Claude Code cloud session, run [tools/cloud](../tools/cloud) first.

A scheduled task writes up and checks three projects a day, and files what is hekireki's on
[hekireki](https://github.com/nakita628/hekireki/issues): see
[.claude/skills/pull-check](../.claude/skills/pull-check/SKILL.md).

## License

Each project's tables (`source.sql`) and their write-up (`schema.prisma`) come from that project's
definitions, under that project's license (`source.json`, or the example's README in apps/).
