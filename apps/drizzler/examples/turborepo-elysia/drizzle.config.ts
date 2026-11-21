import { defineConfig } from "drizzle-kit";

export default defineConfig({
  out: "./drizzle",
  schema: ["./src/drizzle/table.schema.ts", "./src/drizzle/table.relation.ts"],
  dialect: "postgresql",
});
