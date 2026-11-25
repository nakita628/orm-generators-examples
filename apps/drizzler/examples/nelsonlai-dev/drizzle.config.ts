import { defineConfig } from "drizzle-kit";

export default defineConfig({
  dialect: "postgresql",
  schema: ["./src/schema/index.ts", "./src/relations.ts"],
  out: "./drizzle",
});
