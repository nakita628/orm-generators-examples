// Stand-in: only the types the schema imports from this file.
import type { MediaType } from "@/lib/utils/enums";

export type CoverType = `${MediaType}-covers` | "profile-covers" | "profile-back-covers";
