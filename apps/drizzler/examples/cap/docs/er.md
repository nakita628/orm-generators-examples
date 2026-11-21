```mermaid
erDiagram
    storage_integrations |o--o{ videos : "(id) - (storageIntegrationId)"
    videos ||--o| video_edits : "(id) - (videoId)"
    videos ||--o{ video_viewer_grants : "(id) - (videoId)"
    messenger_conversations ||--o{ messenger_messages : "(id) - (conversationId)"
    messenger_conversations ||--o{ messenger_support_emails : "(id) - (conversationId)"
    storage_integrations ||--o{ storage_objects : "(id) - (integrationId)"
    developer_apps ||--o{ developer_app_domains : "(id) - (appId)"
    developer_apps ||--o{ developer_api_keys : "(id) - (appId)"
    developer_apps ||--o{ developer_videos : "(id) - (appId)"
    developer_apps ||--o{ developer_transcripts : "(id) - (appId)"
    developer_apps ||--o| developer_credit_accounts : "(id) - (appId)"
    developer_credit_accounts ||--o{ developer_credit_transactions : "(id) - (accountId)"
    developer_apps ||--o{ developer_daily_storage_snapshots : "(id) - (appId)"
    developer_apps ||--o{ developer_delivery_daily : "(id) - (appId)"
    developer_apps ||--o{ developer_webhook_endpoints : "(id) - (appId)"
    developer_apps ||--o{ developer_video_imports : "(id) - (appId)"
    organizations ||..o{ notifications : "(id) - (orgId)"
    users ||..o{ notifications : "(id) - (recipientId)"
    users ||..o{ comments : "(id) - (authorId)"
    videos ||..o{ comments : "(id) - (videoId)"
    comments |o..o{ comments : "(id) - (parentCommentId)"
    users |o..o{ messenger_conversations : "(id) - (userId)"
    users |o..o{ messenger_messages : "(id) - (userId)"
    users ||..o{ messenger_support_emails : "(id) - (userId)"
    users ||..o{ accounts : "(id) - (userId)"
    users ||..o{ s3_buckets : "(id) - (ownerId)"
    organizations |o..o{ s3_buckets : "(id) - (organizationId)"
    users ||..o{ storage_integrations : "(id) - (ownerId)"
    organizations |o..o{ storage_integrations : "(id) - (organizationId)"
    videos |o..o{ storage_objects : "(id) - (videoId)"
    users ||..o{ organizations : "(id) - (ownerId)"
    users ||..o{ sessions : "(id) - (userId)"
    users ||..o{ organization_members : "(id) - (userId)"
    organizations ||..o{ organization_members : "(id) - (organizationId)"
    organizations ||..o{ integration_installations : "(id) - (organizationId)"
    users ||..o{ integration_installations : "(id) - (installedByUserId)"
    organizations ||..o{ organization_invites : "(id) - (organizationId)"
    users ||..o{ organization_invites : "(id) - (invitedByUserId)"
    users ||..o{ videos : "(id) - (ownerId)"
    video_edits ||..o| videos : "(videoId) - (id)"
    folders |o..o{ videos : "(id) - (folderId)"
    videos ||..o{ shared_videos : "(id) - (videoId)"
    organizations ||..o{ shared_videos : "(id) - (organizationId)"
    users ||..o{ shared_videos : "(id) - (sharedByUserId)"
    organizations ||..o{ spaces : "(id) - (organizationId)"
    users ||..o{ spaces : "(id) - (createdById)"
    spaces ||..o{ space_members : "(id) - (spaceId)"
    users ||..o{ space_members : "(id) - (userId)"
    spaces ||..o{ space_videos : "(id) - (spaceId)"
    videos ||..o{ space_videos : "(id) - (videoId)"
    users ||..o{ space_videos : "(id) - (addedById)"
    organizations ||..o{ folders : "(id) - (organizationId)"
    users ||..o{ folders : "(id) - (createdById)"
    folders |o..o{ folders : "(id) - (parentId)"
    users ||..o{ developer_apps : "(id) - (ownerId)"
    developer_credit_accounts ||..o| developer_apps : "(appId) - (id)"
    users ||..o{ developer_credit_accounts : "(id) - (ownerId)"
    users {
        varchar id PK
        varchar name
        varchar lastName
        varchar email
        timestamp emailVerified
        varchar image
        varchar stripeCustomerId
        varchar stripeSubscriptionId
        varchar thirdPartyStripeSubscriptionId
        varchar stripeSubscriptionStatus
        varchar stripeSubscriptionPriceId
        json preferences
        varchar activeOrganizationId
        timestamp created_at
        timestamp updated_at
        json onboardingSteps
        timestamp onboarding_completed_at
        varchar customBucket
        integer inviteQuota
        varchar defaultOrgId
        integer authSessionVersion
        varchar marketingOrigin
    }
    loops_sync_jobs {
        varchar userId PK
        integer revision
        datetime nextAttemptAt
        varchar leaseToken
        datetime leaseUntil
        integer failures
        varchar lastError
        varchar profileHash
        datetime teammateJoinedAt
        varchar syncedEmail
        datetime lastSyncedAt
    }
    accounts {
        varchar id PK
        varchar userId FK
        varchar type
        varchar provider
        varchar providerAccountId
        text access_token
        integer expires_in
        text id_token
        text refresh_token
        integer refresh_token_expires_in
        varchar scope
        varchar token_type
        timestamp createdAt
        timestamp updatedAt
        text tempColumn
    }
    sessions {
        varchar id PK
        varchar sessionToken
        varchar userId FK
        datetime expires
        timestamp created_at
        timestamp updated_at
    }
    verification_tokens {
        varchar identifier PK
        varchar token
        datetime expires
        timestamp created_at
        timestamp updated_at
    }
    organizations {
        varchar id PK
        varchar name
        varchar ownerId FK
        json metadata
        timestamp tombstoneAt
        varchar allowedEmailDomain
        varchar defaultVideoVisibility
        varchar customDomain
        timestamp domainVerified
        json settings
        varchar iconUrl
        varchar shareableLinkIconUrl
        timestamp createdAt
        timestamp updatedAt
        varchar workosOrganizationId
        varchar workosConnectionId
    }
    organization_sso {
        varchar organizationId PK
        varchar purchasedByUserId
        varchar stripeCustomerId
        varchar stripeSubscriptionId
        varchar stripePriceId
        varchar status
        datetime paidThrough
        datetime currentPeriodEnd
        boolean cancelAtPeriodEnd
        varchar checkoutAttemptId
        varchar checkoutCurrency
        varchar checkoutPriceId
        varchar checkoutSessionId
        datetime checkoutStartedAt
        timestamp createdAt
        timestamp updatedAt
    }
    organization_directory_sync {
        varchar organizationId PK
        datetime requestedAt
        datetime syncedAt
        datetime leaseExpiresAt
        varchar lastError
        integer heldRemovals
        datetime removalsApprovedAt
        timestamp createdAt
        timestamp updatedAt
    }
    directory_users {
        varchar id PK
        varchar organizationId
        varchar directoryId
        varchar userId
        varchar email
        boolean active
        timestamp createdAt
        timestamp updatedAt
    }
    organization_members {
        varchar id PK
        varchar userId FK
        varchar organizationId FK
        varchar role
        boolean hasProSeat
        timestamp createdAt
        timestamp updatedAt
    }
    integration_installations {
        varchar id PK
        varchar provider
        varchar externalId
        varchar displayName
        varchar organizationId FK
        varchar installedByUserId FK
        text encryptedCredentials
        json metadata
        timestamp createdAt
        timestamp updatedAt
    }
    organization_invites {
        varchar id PK
        varchar organizationId FK
        varchar invitedEmail
        varchar invitedByUserId FK
        varchar role
        varchar status
        timestamp createdAt
        timestamp updatedAt
        timestamp expiresAt
    }
    folders {
        varchar id PK
        varchar name
        varchar color
        boolean public "Internet-facing public collection link (/c/[id])."
        json settings
        varchar organizationId FK
        varchar createdById FK
        varchar parentId FK
        varchar spaceId
        timestamp createdAt
        timestamp updatedAt
    }
    videos {
        varchar id PK, FK
        varchar ownerId FK
        varchar orgId
        varchar name
        varchar bucket
        varchar storageIntegrationId FK
        float duration "in seconds"
        integer width
        integer height
        integer fps
        json metadata
        boolean public
        json settings
        varchar transcriptionStatus
        json source
        varchar folderId FK
        timestamp createdAt
        datetime effectiveCreatedAt
        timestamp updatedAt
        text password "PRIVATE"
        text xStreamInfo "LEGACY"
        timestamp firstViewEmailSentAt
        boolean isScreenshot
        varchar awsRegion "DEPRECATED"
        varchar awsBucket
        varchar videoStartTime
        varchar audioStartTime
        varchar jobId
        varchar jobStatus
        boolean skipProcessing
    }
    video_edits {
        varchar videoId PK, FK
        varchar sourceKey
        json editSpec
        timestamp createdAt
        timestamp updatedAt
    }
    shared_videos {
        varchar id PK
        varchar videoId FK
        varchar folderId
        varchar organizationId FK
        varchar sharedByUserId FK
        timestamp sharedAt
    }
    video_viewer_grants {
        varchar id PK
        varchar videoId FK
        varchar email
        varchar invitedByUserId
        timestamp createdAt
        timestamp revokedAt
    }
    comments {
        varchar id PK
        varchar type
        text content
        float timestamp
        varchar authorId FK
        varchar videoId FK
        timestamp createdAt
        timestamp updatedAt
        varchar parentCommentId FK
        varchar mediaKey "Media comments ('video'/'audio' type): object key under the parent video's prefix (`${ownerId}/${videoId}/comments/${commentId}/...`), served via /api/storage/object which asserts that prefix."
        float mediaDuration
        json mediaMeta
    }
    notifications {
        varchar id PK
        varchar orgId FK
        varchar recipientId FK
        varchar type
        json data
        varchar videoId
        varchar dedupKey
        timestamp readAt
        timestamp createdAt
    }
    messenger_conversations {
        varchar id PK
        varchar agent
        varchar mode
        varchar userId FK
        varchar anonymousId
        varchar takeoverByUserId
        timestamp takeoverAt
        timestamp createdAt
        timestamp updatedAt
        timestamp lastMessageAt
    }
    messenger_messages {
        varchar id PK
        varchar conversationId FK
        varchar role
        text content
        varchar userId FK
        varchar anonymousId
        timestamp createdAt
    }
    messenger_support_emails {
        varchar id PK
        varchar conversationId FK
        varchar userId FK
        varchar userEmail
        varchar subject
        text message
        timestamp createdAt
    }
    signed_baas {
        varchar id PK
        varchar organizationId
        varchar userId
        varchar status
        varchar stripeSubscriptionId
        varchar entityName
        varchar entityType
        varchar entityAddress
        varchar signerName
        varchar signerTitle
        varchar noticesEmail
        longtext signatureData
        timestamp signedAt
        timestamp emailSentAt
        timestamp createdAt
        timestamp updatedAt
    }
    s3_buckets {
        varchar id PK
        varchar ownerId FK
        varchar organizationId FK
        text region
        text endpoint
        text bucketName
        text accessKeyId
        text secretAccessKey
        text provider
        boolean active
        timestamp createdAt
        timestamp updatedAt
    }
    storage_integrations {
        varchar id PK
        varchar ownerId FK
        varchar organizationId FK
        varchar provider
        varchar displayName
        varchar status
        boolean active
        text encryptedConfig
        text googleDriveAccessToken
        timestamp googleDriveAccessTokenExpiresAt
        varchar googleDriveTokenRefreshLeaseId
        timestamp googleDriveTokenRefreshLeaseExpiresAt
        json googleDriveStorageQuotaCache
        timestamp createdAt
        timestamp updatedAt
    }
    storage_objects {
        varchar id PK
        varchar integrationId FK
        varchar ownerId
        varchar videoId FK
        text objectKey
        varchar objectKeyHash
        varchar providerObjectId
        text uploadSessionUrl
        varchar uploadStatus
        varchar contentType
        bigint contentLength
        json metadata
        timestamp createdAt
        timestamp updatedAt
    }
    auth_api_keys {
        varchar id PK
        varchar userId
        varchar source
        timestamp createdAt
    }
    mobile_push_tokens {
        varchar id PK
        varchar userId
        varchar token
        varchar sessionHash "SHA-256 of the auth_api_keys id that registered it, so the token dies with that session."
        varchar environment
        timestamp createdAt
        timestamp updatedAt
    }
    mobile_push_claims {
        varchar id PK
        timestamp createdAt
    }
    agent_api_keys {
        varchar id PK
        varchar userId
        varchar tokenHash
        varchar name
        json scopes
        timestamp expiresAt
        timestamp revokedAt
        timestamp createdAt
        timestamp lastUsedAt
    }
    agent_api_idempotency {
        varchar id PK
        varchar userId
        varchar operation
        varchar keyHash
        varchar requestHash
        varchar state
        integer statusCode
        json response
        timestamp expiresAt
        timestamp createdAt
        timestamp updatedAt
    }
    agent_api_authorization_codes {
        varchar id PK
        varchar userId
        varchar codeHash
        varchar codeChallenge
        varchar redirectUri
        json scopes
        timestamp expiresAt
        timestamp consumedAt
        timestamp createdAt
    }
    mcp_oauth_clients {
        varchar id PK
        varchar clientId
        varchar clientName
        json redirectUris
        timestamp activatedAt
        timestamp createdAt
    }
    mcp_oauth_registration_quotas {
        varchar windowId PK
        integer registrations
        timestamp expiresAt
    }
    mcp_oauth_codes {
        varchar id PK
        varchar userId
        varchar clientId
        varchar codeHash
        varchar codeChallenge
        varchar redirectUri
        varchar resource
        timestamp expiresAt
        timestamp consumedAt
        timestamp createdAt
    }
    mcp_oauth_tokens {
        varchar id PK
        varchar userId
        varchar clientId
        varchar familyId
        varchar resource
        varchar accessHash
        varchar refreshHash
        timestamp accessExpiresAt
        timestamp refreshExpiresAt
        timestamp revokedAt
        timestamp createdAt
    }
    agent_api_operations {
        varchar id PK
        varchar userId
        varchar kind
        varchar resourceId
        varchar resultResourceId
        varchar state
        json payload
        json result
        varchar errorCode
        text errorMessage
        timestamp createdAt
        timestamp updatedAt
        timestamp completedAt
    }
    spaces {
        varchar id PK
        boolean primary
        varchar name
        varchar organizationId FK
        varchar createdById FK
        varchar iconUrl
        varchar description
        json settings
        text password
        timestamp createdAt
        timestamp updatedAt
        varchar privacy "Org-internal browsability: 'Public' spaces are visible to all org members. Unrelated to the internet-facing `public` flag below."
        boolean public "Internet-facing public collection link (/c/[id])."
    }
    space_members {
        varchar id PK
        varchar spaceId FK
        varchar userId FK
        varchar role
        timestamp createdAt
        timestamp updatedAt
    }
    space_videos {
        varchar id PK
        varchar spaceId FK
        varchar folderId
        varchar videoId FK
        varchar addedById FK
        timestamp addedAt
    }
    video_uploads {
        varchar video_id PK
        bigint uploaded
        bigint total
        timestamp started_at
        timestamp updated_at
        varchar mode
        varchar phase
        integer processing_progress
        varchar processing_message
        text processing_error
        varchar raw_file_key
    }
    video_processing_jobs {
        varchar video_id PK
        varchar owner_id
        varchar generation
        varchar manifest_sha256
        varchar state
        varchar attempt_id
        integer attempt_count
        datetime lease_expires_at
        datetime next_retry_at
        varchar workflow_run_id
        varchar remote_job_id
        json source
        json verification
        json output
        varchar error_code
        text error_message
        datetime created_at
        datetime updated_at
    }
    media_processing_budgets {
        varchar id PK
        bigint reserved_bytes
        bigint limit_bytes
        datetime expires_at
    }
    imported_videos {
        varchar id
        varchar orgId PK
        varchar source PK
        varchar source_id PK
    }
    loom_migration_requests {
        varchar id PK
        varchar organizationId
        varchar activeOrganizationId
        varchar requestedByUserId
        varchar workspaceName
        text customerNote
        text customerReply
        datetime invitedAt
        varchar status
        text capMessage
        integer expectedVideoCount
        integer importedVideoCount
        integer queuedVideoCount
        integer activeImportCount
        varchar activeImportLeaseToken
        datetime activeImportLeaseUntil
        varchar lastOperatorUserId
        datetime lastOperatorAt
        datetime completedAt
        datetime createdAt
        datetime updatedAt
    }
    loom_migration_imports {
        varchar videoId PK
        varchar requestId
        datetime createdAt
    }
    developer_apps {
        varchar id PK, FK
        varchar ownerId FK
        varchar name
        varchar environment
        varchar logoUrl
        timestamp deletedAt
        timestamp createdAt
        timestamp updatedAt
    }
    developer_app_domains {
        varchar id PK
        varchar appId FK
        varchar domain
        timestamp createdAt
    }
    developer_api_keys {
        varchar id PK
        varchar appId FK
        varchar keyType
        varchar keyPrefix
        varchar keyHash
        text encryptedKey
        timestamp lastUsedAt
        timestamp revokedAt
        timestamp createdAt
    }
    developer_videos {
        varchar id PK
        varchar appId FK
        varchar externalUserId
        varchar name
        float duration
        integer width
        integer height
        integer fps
        varchar s3Key
        varchar transcriptionStatus
        json metadata
        varchar status
        varchar playbackId
        varchar maxResolutionTier
        json renditions
        bigint storedBytes
        varchar errorReason
        json processing
        timestamp storagePaidUntil
        timestamp deletedAt
        timestamp filesDeletedAt
        timestamp createdAt
        timestamp updatedAt
    }
    developer_transcripts {
        varchar id PK
        varchar appId FK
        varchar externalUserId
        varchar name
        varchar status
        varchar sourceType
        varchar sourceUrl
        varchar videoId
        timestamp requestedAt
        varchar s3Key
        varchar resultKey
        bigint sourceBytes
        float durationSeconds
        varchar requestedLanguage
        varchar language
        mediumtext text
        integer wordCount
        bigint reservedMicroCredits
        bigint billedMicroCredits
        varchar errorCode
        varchar errorMessage
        json metadata
        varchar runId
        timestamp queuedAt
        timestamp startedAt
        timestamp completedAt
        timestamp deletedAt
        timestamp createdAt
        timestamp updatedAt
    }
    developer_credit_accounts {
        varchar id PK
        varchar appId FK
        varchar ownerId FK
        bigint balanceMicroCredits
        varchar stripeCustomerId
        varchar stripePaymentMethodId
        boolean autoTopUpEnabled
        bigint autoTopUpThresholdMicroCredits
        integer autoTopUpAmountCents
        timestamp createdAt
        timestamp updatedAt
    }
    developer_credit_transactions {
        varchar id PK
        varchar accountId FK
        varchar type
        bigint amountMicroCredits
        bigint balanceAfterMicroCredits
        varchar referenceId
        varchar referenceType
        json metadata
        timestamp createdAt
    }
    developer_daily_storage_snapshots {
        varchar id PK
        varchar appId FK
        varchar snapshotDate
        float totalDurationMinutes
        integer videoCount
        bigint microCreditsCharged
        timestamp processedAt
        timestamp createdAt
    }
    developer_delivery_daily {
        varchar appId PK, FK
        date day PK
        varchar tier PK
        bigint seconds
        bigint bytes
        integer requests
        bigint abandonedBytes
        boolean final
        timestamp finalizedAt
        timestamp updatedAt
    }
    developer_delivery_log_files {
        varchar logFile PK
        varchar etag
        datetime logHour
        json parts
        timestamp processedAt
    }
    developer_webhook_endpoints {
        varchar id PK
        varchar appId FK
        varchar url
        varchar secretPrefix
        text encryptedSecret
        json events
        boolean enabled
        varchar disabledReason
        timestamp lastDeliveryAt
        varchar lastDeliveryStatus
        integer lastResponseStatus
        varchar lastErrorCode
        integer consecutiveFailures
        timestamp failingSince
        timestamp deletedAt
        timestamp createdAt
        timestamp updatedAt
    }
    developer_video_imports {
        varchar videoId PK
        varchar appId FK
        varchar keyId
        varchar sourceUrl
        varchar state
        bigint bytes
        varchar errorCode
        timestamp startedAt
        timestamp finishedAt
        timestamp createdAt
    }
```
