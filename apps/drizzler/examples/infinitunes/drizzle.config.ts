import { defineConfig } from "drizzle-kit";

export default defineConfig({
  dialect: "postgresql",
  schema: "./src/schema.ts",
  tablesFilter: ["infinitunes_*"],
  out: "./drizzle",
});
