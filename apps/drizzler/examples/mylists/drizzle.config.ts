import { defineConfig } from "drizzle-kit";

export default defineConfig({
  schema: "./src/lib/server/database/schema/index.ts",
  dialect: "sqlite",
  out: "./drizzle",
});
