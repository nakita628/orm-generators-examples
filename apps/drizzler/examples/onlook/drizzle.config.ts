import { defineConfig } from "drizzle-kit";

export default defineConfig({
  schema: "./src/schema",
  out: "./drizzle",
  dialect: "postgresql",
  schemaFilter: ["public"],
  entities: {
    roles: {
      provider: "supabase",
    },
  },
});
