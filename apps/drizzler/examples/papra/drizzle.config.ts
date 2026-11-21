import { defineConfig } from "drizzle-kit";

export default defineConfig({
  schema: ["./src/modules/**/*.table.ts", "./src/modules/**/*.tables.ts"],
  dialect: "turso",
  out: "./drizzle",
});
