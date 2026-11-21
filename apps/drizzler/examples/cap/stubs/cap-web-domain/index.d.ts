// Stand-in for @cap/web-domain: the schema imports only types from it, as `$type<>` arguments.
export type AiGenerationLanguage = string;
export namespace Agent { type AgentScope = string; }
export namespace Comment { type CommentId = string; type MediaMeta = Record<string, unknown>; }
export namespace Folder { type FolderId = string; }
export namespace ImageUpload { type ImageUrlOrKey = string; }
export namespace Organisation { type OrganisationId = string; }
export namespace PublicCollection { type PublicPageSettings = Record<string, unknown>; }
export namespace S3Bucket { type S3BucketId = string; }
export namespace Space { type SpaceIdOrOrganisationId = string; }
export namespace Storage {
  type StorageIntegrationId = string;
  type StorageIntegrationStatus = string;
  type StorageObjectId = string;
  type StorageObjectMetadata = Record<string, unknown>;
  type StorageProvider = string;
}
export namespace User { type UserId = string; }
export namespace Video { type VideoId = string; }
