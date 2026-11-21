import { defineConfig } from "drizzle-kit";

export default defineConfig({
  dialect: "postgresql",
  schema: "./server/src/api/db/schema.ts",
  out: "./drizzle",
});
