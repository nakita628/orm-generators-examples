import { defineConfig } from "drizzle-kit";

export default defineConfig({
  dialect: "postgresql",
  schema: "./lib/db/schema",
  tablesFilter: ["lipi_*"],
  out: "./drizzle",
});
