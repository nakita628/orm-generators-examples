// Stand-in: only what the tables import from this file.
export type JsonSerializableValue = string | number | boolean | null | JsonSerializableValue[] | { [key: string]: JsonSerializableValue };
