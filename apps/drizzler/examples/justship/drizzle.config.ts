import { defineConfig } from "drizzle-kit";

export default defineConfig({
  dialect: "turso",
  schema: "./src/lib/server/database/schema.ts",
  out: "./drizzle",
});
