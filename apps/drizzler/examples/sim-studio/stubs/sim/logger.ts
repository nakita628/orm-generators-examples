// Stand-in for @sim/logger: only what sealed-text.ts imports.
export function createLogger(_name: string, _options?: Record<string, unknown>) {
  return { info: (..._args: unknown[]) => {}, warn: (..._args: unknown[]) => {} };
}
