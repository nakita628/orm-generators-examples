package plunk.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * One transactional notification held for one recipient (or for the whole project) by a
 * send with `delivery.throttle`. While it is held, the same notification is not sent again.
 *
 * The claim is taken by inserting (or taking over an expired) row, so the unique constraint
 * decides the race between concurrent sends, like `IdempotencyKey`. It is held from the
 * moment the send is accepted, re-anchored to the real send time once SES has the message,
 * and released when the email never gets there (cancelled, or failed before sending).
 */
class ThrottleClaimEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ThrottleClaimEntity>(ThrottleClaimTable)

    var project by ProjectEntity referencedOn ThrottleClaimTable.projectId
    var projectId by ThrottleClaimTable.projectId
    var key by ThrottleClaimTable.key
    var recipientKey by ThrottleClaimTable.recipientKey
    var contact by ContactEntity optionalReferencedOn ThrottleClaimTable.contactId
    var contactId by ThrottleClaimTable.contactId
    var emailId by ThrottleClaimTable.emailId
    var workflow by WorkflowEntity optionalReferencedOn ThrottleClaimTable.workflowId
    var workflowId by ThrottleClaimTable.workflowId
    var windowSeconds by ThrottleClaimTable.windowSeconds
    var expiresAt by ThrottleClaimTable.expiresAt
    var createdAt by ThrottleClaimTable.createdAt
    var updatedAt by ThrottleClaimTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(ThrottleClaimTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
