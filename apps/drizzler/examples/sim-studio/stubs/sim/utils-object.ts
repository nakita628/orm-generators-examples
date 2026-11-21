// Stand-in for @sim/utils/object: only what sealed-text.ts imports.
export function toRecord(value: unknown): Record<string, unknown> {
  return typeof value === "object" && value !== null ? (value as Record<string, unknown>) : {};
}
