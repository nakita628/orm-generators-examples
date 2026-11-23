---
name: pull-check
description: The daily pull of orm-generators-examples. Three well-known projects' own ORMs (Django, Drizzle, Active Record, SQLAlchemy) create their tables, the SQL is written up in Prisma by `prisma db pull` under pull/<orm>/<project>/, hekireki generates every ORM from that write-up, and the generated code runs against those tables; what is hekireki's fault becomes an issue on nakita628/hekireki. Use when the scheduled pull task fires, or when asked to "pull", "ORM の SQL を Prisma に書き起こす", or to check hekireki against another ORM's schema.
---

# Daily pull

Goal: every day, three projects whose schemas were not written for Prisma go through
`pull/pull.mjs` (read its header and `pull/README.md` once): their SQL is written up in Prisma under
`pull/<orm>/<project>/`, and every ORM hekireki generates from the write-up is checked against the
project's own tables. What hekireki gets wrong becomes an issue on nakita628/hekireki; the
write-ups and results go to `main`.

Work from the root of the repository.

## 0. Set up

1. `git checkout main && git pull --ff-only`, then `pnpm install`.
2. `pkg-config --exists mysqlclient || apt-get install -y libmysqlclient-dev pkg-config` (the
   compat Django harness builds mysqlclient).
3. Docker. On a machine of your own, use it as it is. In a cloud session:
   `eval "$(bash tools/cloud/prepare.sh | tail -1)"`. It builds every compat image (a few
   minutes; SeaORM's is the longest). If an image does not build, say which in the report and go
   on with `--orm` set to the others.
4. Attach `nakita628/hekireki` (add_repo, access `push`; no need to clone it), to file issues.

## 1. Run

```bash
node pull/pull.mjs --random 3
```

It picks projects never pulled before (in an order fixed by the date), then the ones checked
longest ago, and for each: the project's ORM creates the tables, they are dumped to `source.sql`,
`prisma db pull` writes `schema.prisma`, every hekireki generator runs on it, and compat's
harnesses run the generated code against a database made from `source.sql`. It takes from a few
minutes to an hour. A project that cannot be pulled is recorded in its `source.json` and shown
under "Not pulled" in its layer's results.

## 2. Read the results

`pull/<orm>/results/README.md` (the layer's matrix) and `pull/<orm>/results/<project>.md` (each
ORM's findings, and the diff of the tables it creates itself). Beside each write-up in
`pull/<orm>/<project>/`: `source.sql`, `schema.prisma` (its header lists the changes pull.mjs made
so that Prisma accepts it), `prisma-diff.sql`, `source.json`, and `generated/` (not committed).

## 3. Is it still there in hekireki's latest release?

apps/hekireki pins a hekireki version that can be behind the latest. Before taking anything as
hekireki's, run the same projects with the latest release:

```bash
node pull/pull.mjs --hekireki latest <orm>/<project> ...
```

It writes to the cache (the path is printed) and leaves pull/ as it is. A finding that is gone
there is already fixed: do not file it (the daily upstream sync brings the new release in).

## 4. Whose is each finding?

Read the generated code and the table (`source.sql`) before deciding.

**hekireki's**: file it (step 5).
- a generator fails (`generated/<orm>.error`, "❌ generator failed") on a schema Prisma accepts;
- the generated code does not compile or load (a "did not run" whose error is in the generated
  code, not in the harness);
- the generated code disagrees with the table it is meant for: a wrong column name or type, a
  nullability, a default the ORM does not fill (or fills where the database does), a relation
  that does not load, a key it cannot use;
- the tables the ORM creates on its own differ from the write-up in something the write-up
  states (a lost index or unique, a wrong type or nullability, a dropped `map:` name, a wrong
  `onDelete`).

**Prisma's**: not filed. Anything in `prisma-diff.sql`, anything listed in the write-up's header
under "Changed from what `prisma db pull` wrote", a column Prisma writes up as `Unsupported(...)`,
and a pull that does not validate.

**The pipeline's**: fix it here, not upstream.
- The project's ORM could not create its tables (a missing extension, a function the app creates
  outside its schema, a gem's DSL): give the source a `before_sql` / `after_sql` in the layer's
  `sources.json`, or teach the loader (`pull/activerecord/load_schema.rb`), and pull it again. If
  it cannot be done, add it to `skip` of the layer's `sources.json` with the reason.
- A compat harness or the plan fails on something valid: fix it in apps/hekireki/compat
  (`lib/plan.mjs`, `<orm>/harness.*`), in the style of the file, so that every ORM is held to the
  table, and run again. Say in the report what changed.

When the same finding is already in compat's forward results (`apps/hekireki/compat/results/`),
it is likely the same root cause: check its issue first.

## 5. File what is hekireki's

1. Search first, open and closed:
   `gh api 'repos/nakita628/hekireki/issues?state=all&per_page=100' --jq '.[] | select(.pull_request == null) | "\(.number) \(.state) \(.title)"'`
   (more pages with `&page=2`…). Same problem and open: comment with the new case. Closed in a
   release older than the latest: it is back; open a new issue that links it.
2. Reduce it to the smallest Prisma schema that shows it (a model or two from the write-up), and
   run the generator on that alone to confirm, with the latest release
   (`node ~/.cache/orm-generators-examples/pull/hekireki-<version>/generate.mjs --schema <file> --out generated:<dir>`).
3. Write it in English, in the style of hekireki's existing issues: the title names the generator
   and the symptom (`hekireki-sea-orm: an Int @db.SmallInt field is i32, which SQLx will not decode from int2`);
   the body gives the hekireki version, the project (`activerecord/mastodon`, with the upstream
   commit from `source.json`), that the schema is `prisma db pull` of the tables the project's
   own ORM created, the reduced schema, what the generator wrote, the error against the table,
   what it should write and why. End with where it was found
   (`https://github.com/nakita628/orm-generators-examples/tree/main/pull`).
   `gh api repos/nakita628/hekireki/issues -f title='…' -f body='…'`
4. One issue per root cause, never one without the reduced schema. List the issues in the report.

## 6. Push

```bash
git add pull apps/hekireki/compat
git commit -m "pull: <orm>/<project>, <orm>/<project>, <orm>/<project>"
git pull --rebase && git push origin main
```

`pull/*/*/generated/` is ignored; never commit it. If a run also rewrote compat's forward results
(`apps/hekireki/compat/results/`), leave those out. Retry a failed push up to 4 times with backoff.

## 7. Report

End with a short report in Japanese: the three projects, their result rows, the issues filed or
commented on (links), what was Prisma's, what was fixed in the pipeline, and what could not run.
