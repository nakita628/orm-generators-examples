// Stand-in for @onlook/models: the enums the schema imports are copied as they are; the types it
// imports are loose.

// From packages/models/src/chat/conversation/index.ts
export enum AgentType {
    ROOT = "root",
    USER = "user",
}

// From packages/models/src/project/create.ts
export enum ProjectCreateRequestStatus {
    PENDING = 'pending',
    COMPLETED = 'completed',
    FAILED = 'failed',
}

// From packages/models/src/domain/index.ts
export enum VerificationRequestStatus {
    PENDING = 'pending',
    VERIFIED = 'verified',
    CANCELLED = 'cancelled',
}

// From packages/models/src/hosting/index.ts
export enum DeploymentStatus {
    PENDING = 'pending',
    IN_PROGRESS = 'in_progress',
    COMPLETED = 'completed',
    FAILED = 'failed',
    CANCELLED = 'cancelled',
}

// From packages/models/src/hosting/index.ts
export enum DeploymentType {
    PREVIEW = 'preview',
    CUSTOM = 'custom',
    UNPUBLISH_PREVIEW = 'unpublish_preview',
    UNPUBLISH_CUSTOM = 'unpublish_custom',
}

// From packages/models/src/usage/index.ts
export enum UsageType {
    MESSAGE = 'message',
    DEPLOYMENT = 'deployment',
}

// From packages/models/src/project/role.ts
export enum ProjectRole {
    OWNER = 'owner',
    ADMIN = 'admin',
}

export type ChatSuggestion = Record<string, unknown>;
export type CreateRequestContext = Record<string, unknown>;
export type ChatMessage = Record<string, unknown>;
export type MessageCheckpoints = Record<string, unknown>[];
export type MessageContext = Record<string, unknown>[];
export type AVerificationRecord = Record<string, unknown>;
export type TxtVerificationRecord = Record<string, unknown>;
