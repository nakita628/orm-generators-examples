---
name: hekireki-pull-check
description: The daily pull check of orm-generators-examples. Three well-known projects' own ORMs (Django, Drizzle, Active Record, SQLAlchemy) create their tables, `prisma db pull` reads them back, hekireki generates every ORM from that schema, and the generated code runs against those tables; what is hekireki's fault becomes an issue on nakita628/hekireki. Use when the scheduled pull-check task fires, or when asked to "pull check", "逆方向の互換性チェック", or to check hekireki against another ORM's schema.
---

# Daily pull check

Goal: every day, three projects whose schemas were not written for Prisma go through
`apps/hekireki/compat/pull.mjs` (read its header and the "Pulled schemas" section of
`apps/hekireki/compat/README.md` once). What hekireki gets wrong becomes an issue on
nakita628/hekireki; the results go to `main`.

Work from the root of the repository.

## 0. Set up

1. `git checkout main && git pull --ff-only`, then `pnpm install`.
2. `pkg-config --exists mysqlclient || apt-get install -y libmysqlclient-dev pkg-config` (the
   Django harness builds mysqlclient).
3. Docker. On a machine of your own, use it as it is. In a cloud session:
   `eval "$(bash apps/hekireki/compat/cloud/prepare.sh | tail -1)"`. It builds every compat
   image (a few minutes; the SeaORM one is the longest). If an image does not build, say which
   in the report and go on with `--orm` set to the others.
4. Attach `nakita628/hekireki` (add_repo, access `push`; no need to clone it), to file issues.

## 1. Run

```bash
cd apps/hekireki/compat
node pull.mjs --random 3
```

It picks sources never checked before (in an order fixed by the date), then the ones checked
longest ago, and for each: the project's ORM creates the tables, `prisma db pull`, every hekireki
generator, and run.mjs over every ORM. It takes from a few minutes to an hour (SeaORM builds the
generated entities with Rust; a source with more than ~80 models may not build on SeaORM, as
in the forward results). If one source cannot be pulled at all, pull.mjs says why and goes on.

## 2. Read the results

`apps/hekireki/compat/pull/results/README.md` (the matrix) and `pull/results/<id>.md` (each
ORM's findings, and the diff of the tables it creates itself). Beside each schema in
`pull/schemas/<id>/`: `source.sql`, `schema.prisma` (its header lists the changes pull.mjs made
so that Prisma accepts it), `prisma-diff.sql`, `source.json`, and `generated/` (not committed).

## 3. Whose is each finding?

Read the generated code and the table (`source.sql`) before deciding.

**hekireki's**: file it (step 4).
- a generator fails (`generated/<orm>.error`, "❌ generator failed") on a schema Prisma accepts;
- the generated code does not compile or load (a "did not run" whose error is in the generated
  code, not in the harness);
- the generated code disagrees with the table it is meant for: a wrong column name or type, a
  nullability, a default the ORM does not fill (or fills where the database does), a relation
  that does not load, a key it cannot use;
- the tables the ORM creates on its own differ from the pulled schema in a way the schema states
  (a lost index or unique, a wrong type or nullability, a dropped `map:` name, a wrong `onDelete`).
  A difference in something the Prisma schema does not state is not hekireki's.

**Prisma's**: not filed. Anything in `prisma-diff.sql`, anything listed under the schema
header's "Changed from what `prisma db pull` wrote", and a pull that does not validate.

**The source's or the pipeline's**: fix it here, not upstream.
- The project's ORM could not create its tables (a missing extension, a function the app creates
  outside its schema, a gem's DSL): give the source a `before_sql` / `after_sql` in
  `pull/sources.json`, or teach the loader the DSL (`pull/activerecord/load_schema.rb`), and run
  that source again. If it cannot be done, add it to `skip` with the reason.
- The harness or plan fails on something valid (an `Unsupported(...)` column, an `@@ignore`d
  model): fix `lib/plan.mjs` or the harness, so that every ORM is held to the table, and run again.
  Keep these changes small and in the style of the file; say what changed in the report.

When the same kind of finding is already in the forward results (`compat/results/`), it is likely
the same root cause: check its issue first.

## 4. File what is hekireki's

1. Search first, open and closed:
   `gh api 'repos/nakita628/hekireki/issues?state=all&per_page=100' --jq '.[] | "\(.number) \(.state) \(.title)"'`
   (more pages with `&page=2`…). Same problem and open: comment with the new case. Closed in a
   release older than `source.json`'s `hekireki` version: it is back; open a new issue that links it.
2. Reduce it to the smallest Prisma schema that shows it (a model or two from the pulled schema),
   and run the generator on that alone to confirm
   (`node apps/hekireki/tools/hekireki/generate.mjs --schema <file> --out generated:<dir>`).
3. Write it in English, in the style of hekireki's existing issues: the title names the generator
   and the symptom (`hekireki-django: a varchar(255) primary key pulled from Django loses max_length`);
   the body gives the hekireki version, the source (`rails/mastodon`, with the upstream commit
   link from `source.json`), that the schema came from `prisma db pull` of the tables the
   project's own ORM created, the reduced schema, what the generator wrote, what it should write,
   and why (the ORM's docs, the table in `source.sql`).
   `gh api repos/nakita628/hekireki/issues -f title='…' -f body='…'`
4. One issue per root cause, never one without the reduced schema. List the issue in the report.

## 5. Push

```bash
git add apps/hekireki/compat/pull apps/hekireki/compat/lib apps/hekireki/compat/*.mjs
git commit -m "Pull check: <source>, <source>, <source>"
git pull --rebase && git push origin main
```

`pull/schemas/*/generated/` is ignored; never commit it. If the run also rewrote
`compat/results/` (the forward results), leave those out unless that was the point. Retry a
failed push up to 4 times with backoff.

## 6. Report

End with a short report in Japanese: the three sources (and their projects), the matrix rows,
the issues filed or commented on (links), what was Prisma's, what was fixed in the pipeline, and
what could not run.
