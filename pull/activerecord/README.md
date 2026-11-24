# pull/activerecord

The tables Active Record creates from the `db/schema.rb` of well-known Rails apps, written up in
Prisma. **[Results](results/README.md)**

| App | Database | License |
| --- | --- | --- |
| [Mastodon](https://github.com/mastodon/mastodon) | PostgreSQL | AGPL-3.0 |
| [Forem](https://github.com/forem/forem) | PostgreSQL | AGPL-3.0 |
| [Chatwoot](https://github.com/chatwoot/chatwoot) | PostgreSQL | MIT (core) |
| [Maybe](https://github.com/maybe-finance/maybe) | PostgreSQL | AGPL-3.0 |
| [Loomio](https://github.com/loomio/loomio) | PostgreSQL | AGPL-3.0 |
| [Lobsters](https://github.com/lobsters/lobsters) | SQLite | BSD-3-Clause |
| [Campfire](https://github.com/basecamp/once-campfire) | SQLite | MIT |

Each is pinned to a commit in [sources.json](sources.json). For `activerecord/<app>`, pull.mjs
downloads the app's `db/schema.rb` at that commit and [load_schema.rb](load_schema.rb) loads it
with Active Record alone, the way `rails db:schema:load` does, in apps/hekireki/compat's Active
Record image. What an app's schema.rb takes from a gem rather than from Active Record is defined
there the way the gem does it: scenic's `create_view`, fx's `create_function` and
`create_trigger`, hairtrigger's chained `create_trigger(...).on(...)`, neighbor's `t.vector`. A
schema.rb of a newer Rails than the image's loads under the newest version it knows.

`before_sql` and `after_sql` in sources.json run around the load, for what the app creates
outside its schema.rb: Mastodon's `timestamp_id()` id default and the sequences it draws on.

`skip`: Discourse, whose db/structure.sql `prisma db pull` writes up with relations between integer
and bigint columns that `prisma validate` refuses in more places than pull.mjs changes.

Each project's own `README.md` holds the SQL every ORM creates from the write-up (`sql/`) against
`source.sql`; [results/](results/README.md) has the layer's matrices.
