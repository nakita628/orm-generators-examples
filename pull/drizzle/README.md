# pull/drizzle

The SQL drizzle-kit writes from the schema of every [drizzler example](../../apps/drizzler/examples),
written up in Prisma. **[Results](results/README.md)**

For `drizzle/<example>`, pull.mjs runs the example's own `drizzle-kit export` (the version of the
drizzler catalog in the root `pnpm-workspace.yaml`), which prints the SQL of the Drizzle schema
in the example's dialect (PostgreSQL, MySQL, or SQLite for SQLite and Turso), runs that SQL in an
empty database, dumps the tables to `<example>/source.sql`, and `prisma db pull` writes
`<example>/schema.prisma`.

The examples are listed in apps/drizzler; one that cannot be pulled goes in `skip` of a
`sources.json` here, with the reason.

Each project's own `README.md` holds the SQL every ORM creates from the write-up (`sql/`) against
`source.sql`; [results/](results/README.md) has the layer's matrices.
