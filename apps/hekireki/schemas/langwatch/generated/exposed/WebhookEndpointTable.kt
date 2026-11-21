package langwatch.models

import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Org-anchored outbound webhook endpoint (the webhook platform, enterprise
 * only). The signing secret is AES-256-GCM encrypted at rest (shared
 * encrypt/decrypt, CREDENTIALS_SECRET) and returned in plaintext exactly
 * once, at create or roll time. `enabledEvents` holds registry event types
 * or family wildcards ("gateway.*", "*"). Failure-streak columns drive the
 * 72h auto-disable. Spec: specs/webhooks/webhook-endpoints.feature.
 */
object WebhookEndpointTable : IdTable<String>("\"WebhookEndpoint\"") {
    override val id = text("id").entityId()
    val organizationId = text("organizationId")

    /**
     * Where this endpoint delivers. Everything above the last hop is shared:
     * same batching, same retry ladder, same delivery log, same signature over
     * the same bytes.
     */
    val destinationKind =
        pgEnum("destinationKind", "\"WebhookDestinationKind\"", WebhookDestinationKind.entries, WebhookDestinationKind::dbName).default(
            WebhookDestinationKind.HTTP,
        )

    /**
     * The receiver URL. Present on `http` endpoints and null on the others;
     * the per-kind shape is pinned by a CHECK constraint Prisma cannot model.
     */
    val url = text("url").nullable()

    /**
     * The Amazon SQS queue this endpoint delivers to, on `sqs` endpoints. The
     * region and the owning account are read off it rather than configured
     * separately, so they can never disagree with the queue.
     */
    val sqsQueueUrl = text("sqsQueueUrl").nullable()

    /**
     * The role we assume to write to the queue, and the ExternalId that proves
     * the request came from us rather than from anyone who learned the role's
     * name. The recommended credential mode.
     */
    val sqsRoleArn = text("sqsRoleArn").nullable()
    val sqsExternalId = text("sqsExternalId").nullable()

    /**
     * A static key pair for the queue, encrypted at rest exactly like the
     * signing secret. Both halves are present or both are absent: a half-filled
     * pair is what the AWS SDK reads as "we have credentials" and stops looking
     * with, so the CHECK constraint refuses it.
     */
    val sqsAccessKeyId = text("sqsAccessKeyId").nullable()
    val sqsSecretAccessKeyEncrypted = text("sqsSecretAccessKeyEncrypted").nullable()
    val secretEncrypted = text("secretEncrypted")

    /**
     * How deliveries are formatted and signed. Null: the batch envelope signed
     * `t=,v1=`. `legacy_sha256`: one raw message per POST signed `sha256=`, kept
     * only for endpoints migrated from governance anomaly destinations.
     */
    val signatureScheme = text("signatureScheme").nullable()

    /**
     * Set only by governance's anomaly destination migration: a create that
     * repeats it in the same organization answers the endpoint it made. An
     * archived endpoint gives its key up, so a retry after cleanup creates anew.
     */
    val idempotencyKey = text("idempotencyKey").nullable()

    /**
     * The secret this endpoint signed with before the last roll. Deliveries
     * carry a second `v1=` computed from it until it expires, so a receiver
     * swaps secrets on its own schedule instead of dropping deliveries in the
     * gap. Null once the window closes or if the secret was never rolled.
     */
    val previousSecretEncrypted = text("previousSecretEncrypted").nullable()

    /** When the previous secret stops being signed with and accepted. */
    val previousSecretExpiresAt = pgTimestamp("previousSecretExpiresAt", 3).nullable()
    val enabledEvents = array("enabledEvents", TextColumnType()).nullable().databaseDefault("'{}'") { emptyList() }
    val status =
        pgEnum("status", "\"WebhookEndpointStatus\"", WebhookEndpointStatus.entries, WebhookEndpointStatus::dbName).default(
            WebhookEndpointStatus.ACTIVE,
        )

    /** Why the endpoint is disabled: "auto_failures_72h" or "manual". */
    val disabledReason = text("disabledReason").nullable()
    val disabledAt = pgTimestamp("disabledAt", 3).nullable()

    /**
     * Start of the current consecutive-failure streak; cleared on any 2xx.
     * 72 hours of unbroken failures from this instant auto-disables.
     */
    val failingSince = pgTimestamp("failingSince", 3).nullable()

    /**
     * Delivery controls, each server-bounded (see webhookDeliveryControls
     * constants): envelopes per POST, coalescing window before a partial
     * batch ships, and concurrent POSTs to this receiver.
     */
    val maxBatchSize = integer("maxBatchSize").default(100)
    val maxBatchDelayMs = integer("maxBatchDelayMs").default(250)
    val maxInFlight = integer("maxInFlight").default(4)

    /**
     * Self-hosted only: deliveries accept the receiver's self-signed TLS
     * certificate. Off by default; the hosted product ignores it.
     */
    val allowSelfSignedCertificate = bool("allowSelfSignedCertificate").default(false)
    val lastSuccessAt = pgTimestamp("lastSuccessAt", 3).nullable()
    val lastFailureAt = pgTimestamp("lastFailureAt", 3).nullable()
    val archivedAt = pgTimestamp("archivedAt", 3).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"WebhookEndpoint_pkey\"")

    init {
        index("\"WebhookEndpoint_organizationId_status_idx\"", false, organizationId, status)
        uniqueIndex("\"WebhookEndpoint_organizationId_idempotencyKey_key\"", organizationId, idempotencyKey)
    }
}
