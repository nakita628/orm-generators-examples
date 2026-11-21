// Stand-in for the `ai` package: only the type the schema imports.
declare module "ai" {
  export type ModelMessage = unknown;
}
