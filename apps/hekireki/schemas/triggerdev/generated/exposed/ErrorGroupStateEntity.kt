package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * *
 * * Error group state is used to track when a user has interacted with an error (ignored/resolved)
 * * The actual error data is in ClickHouse.
 */
class ErrorGroupStateEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ErrorGroupStateEntity>(ErrorGroupStateTable)

    var organization by OrganizationEntity referencedOn ErrorGroupStateTable.organizationId
    var organizationId by ErrorGroupStateTable.organizationId
    var project by ProjectEntity referencedOn ErrorGroupStateTable.projectId
    var projectId by ErrorGroupStateTable.projectId

    /**
     * *
     * * You can ignore/resolve an error across all environments, or specific ones
     */
    var environment by RuntimeEnvironmentEntity optionalReferencedOn ErrorGroupStateTable.environmentId
    var environmentId by ErrorGroupStateTable.environmentId
    var taskIdentifier by ErrorGroupStateTable.taskIdentifier
    var errorFingerprint by ErrorGroupStateTable.errorFingerprint
    var status by ErrorGroupStateTable.status

    /**
     * *
     * * Error is ignored until this date
     */
    var ignoredUntil by ErrorGroupStateTable.ignoredUntil

    /**
     * *
     * * Error is ignored until this occurrence rate
     */
    var ignoredUntilOccurrenceRate by ErrorGroupStateTable.ignoredUntilOccurrenceRate

    /**
     * *
     * * Error is ignored until this total occurrences
     */
    var ignoredUntilTotalOccurrences by ErrorGroupStateTable.ignoredUntilTotalOccurrences

    /**
     * Total occurrence count at the time the error was ignored (from ClickHouse).
     * Used with ignoredUntilTotalOccurrences to compute occurrences since ignoring.
     */
    var ignoredAtOccurrenceCount by ErrorGroupStateTable.ignoredAtOccurrenceCount

    /**
     * *
     * * Error was ignored at this date
     */
    var ignoredAt by ErrorGroupStateTable.ignoredAt

    /**
     * *
     * * Reason for ignoring the error
     */
    var ignoredReason by ErrorGroupStateTable.ignoredReason

    /**
     * *
     * * User who ignored the error
     */
    var ignoredByUserId by ErrorGroupStateTable.ignoredByUserId

    /**
     * *
     * * Error was resolved at this date
     */
    var resolvedAt by ErrorGroupStateTable.resolvedAt

    /**
     * *
     * * Error was resolved in this version
     */
    var resolvedInVersion by ErrorGroupStateTable.resolvedInVersion

    /**
     * *
     * * User who resolved the error
     */
    var resolvedBy by ErrorGroupStateTable.resolvedBy
    var createdAt by ErrorGroupStateTable.createdAt
    var updatedAt by ErrorGroupStateTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(ErrorGroupStateTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
