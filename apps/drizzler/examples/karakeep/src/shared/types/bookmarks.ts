// Stand-in for @karakeep/shared/types/bookmarks: only what the schema imports.
export const enum BookmarkTypes {
  LINK = "link",
  TEXT = "text",
  ASSET = "asset",
  UNKNOWN = "unknown",
}

export type ZReaderViewReason = string;
