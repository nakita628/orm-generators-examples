// Stand-in for the image URL helpers: the custom types call them only in toDriver/fromDriver.
import type { CoverType } from "@/lib/types/media-common.types";

export const getImageFilename = (value: string): string => value.split("/").pop() ?? value;
export const getImageUrl = (coverType: CoverType, filename: string): string => `/static/${coverType}/${filename}`;
