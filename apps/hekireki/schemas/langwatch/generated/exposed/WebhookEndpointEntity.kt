package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
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
class WebhookEndpointEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, WebhookEndpointEntity>(WebhookEndpointTable)

    var organizationId by WebhookEndpointTable.organizationId

    /**
     * Where this endpoint delivers. Everything above the last hop is shared:
     * same batching, same retry ladder, same delivery log, same signature over
     * the same bytes.
     */
    var destinationKind by WebhookEndpointTable.destinationKind

    /**
     * The receiver URL. Present on `http` endpoints and null on the others;
     * the per-kind shape is pinned by a CHECK constraint Prisma cannot model.
     */
    var url by WebhookEndpointTable.url

    /**
     * The Amazon SQS queue this endpoint delivers to, on `sqs` endpoints. The
     * region and the owning account are read off it rather than configured
     * separately, so they can never disagree with the queue.
     */
    var sqsQueueUrl by WebhookEndpointTable.sqsQueueUrl

    /**
     * The role we assume to write to the queue, and the ExternalId that proves
     * the request came from us rather than from anyone who learned the role's
     * name. The recommended credential mode.
     */
    var sqsRoleArn by WebhookEndpointTable.sqsRoleArn
    var sqsExternalId by WebhookEndpointTable.sqsExternalId

    /**
     * A static key pair for the queue, encrypted at rest exactly like the
     * signing secret. Both halves are present or both are absent: a half-filled
     * pair is what the AWS SDK reads as "we have credentials" and stops looking
     * with, so the CHECK constraint refuses it.
     */
    var sqsAccessKeyId by WebhookEndpointTable.sqsAccessKeyId
    var sqsSecretAccessKeyEncrypted by WebhookEndpointTable.sqsSecretAccessKeyEncrypted
    var secretEncrypted by WebhookEndpointTable.secretEncrypted

    /**
     * How deliveries are formatted and signed. Null: the batch envelope signed
     * `t=,v1=`. `legacy_sha256`: one raw message per POST signed `sha256=`, kept
     * only for endpoints migrated from governance anomaly destinations.
     */
    var signatureScheme by WebhookEndpointTable.signatureScheme

    /**
     * Set only by governance's anomaly destination migration: a create that
     * repeats it in the same organization answers the endpoint it made. An
     * archived endpoint gives its key up, so a retry after cleanup creates anew.
     */
    var idempotencyKey by WebhookEndpointTable.idempotencyKey

    /**
     * The secret this endpoint signed with before the last roll. Deliveries
     * carry a second `v1=` computed from it until it expires, so a receiver
     * swaps secrets on its own schedule instead of dropping deliveries in the
     * gap. Null once the window closes or if the secret was never rolled.
     */
    var previousSecretEncrypted by WebhookEndpointTable.previousSecretEncrypted

    /** When the previous secret stops being signed with and accepted. */
    var previousSecretExpiresAt by WebhookEndpointTable.previousSecretExpiresAt
    var enabledEvents by WebhookEndpointTable.enabledEvents
    var status by WebhookEndpointTable.status

    /** Why the endpoint is disabled: "auto_failures_72h" or "manual". */
    var disabledReason by WebhookEndpointTable.disabledReason
    var disabledAt by WebhookEndpointTable.disabledAt

    /**
     * Start of the current consecutive-failure streak; cleared on any 2xx.
     * 72 hours of unbroken failures from this instant auto-disables.
     */
    var failingSince by WebhookEndpointTable.failingSince

    /**
     * Delivery controls, each server-bounded (see webhookDeliveryControls
     * constants): envelopes per POST, coalescing window before a partial
     * batch ships, and concurrent POSTs to this receiver.
     */
    var maxBatchSize by WebhookEndpointTable.maxBatchSize
    var maxBatchDelayMs by WebhookEndpointTable.maxBatchDelayMs
    var maxInFlight by WebhookEndpointTable.maxInFlight

    /**
     * Self-hosted only: deliveries accept the receiver's self-signed TLS
     * certificate. Off by default; the hosted product ignores it.
     */
    var allowSelfSignedCertificate by WebhookEndpointTable.allowSelfSignedCertificate
    var lastSuccessAt by WebhookEndpointTable.lastSuccessAt
    var lastFailureAt by WebhookEndpointTable.lastFailureAt
    var archivedAt by WebhookEndpointTable.archivedAt
    var createdAt by WebhookEndpointTable.createdAt
    var updatedAt by WebhookEndpointTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(WebhookEndpointTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
