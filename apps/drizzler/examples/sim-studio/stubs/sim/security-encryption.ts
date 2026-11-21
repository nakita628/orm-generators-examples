// Stand-in for @sim/security/encryption: only what sealed-text.ts imports. No real sealing.
export const isSealed = (value: string): boolean => value.startsWith("sealed:");
export const seal = (plaintext: string, _key: Buffer, _context: string): string => `sealed:${plaintext}`;
export const unseal = (value: string, _key: Buffer, _context: string): string => value.slice(7);
export const sealingKeyId = (_key: Buffer): string => "stub";
