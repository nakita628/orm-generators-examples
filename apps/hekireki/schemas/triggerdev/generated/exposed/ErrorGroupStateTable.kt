package triggerdev.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * *
 * * Error group state is used to track when a user has interacted with an error (ignored/resolved)
 * * The actual error data is in ClickHouse.
 */
object ErrorGroupStateTable : IdTable<String>("\"ErrorGroupState\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"ErrorGroupState_organizationId_fkey\"",
        )
    val projectId =
        reference(
            "projectId",
            ProjectTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"ErrorGroupState_projectId_fkey\"",
        )
    val environmentId =
        optReference(
            "environmentId",
            RuntimeEnvironmentTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"ErrorGroupState_environmentId_fkey\"",
        )
    val taskIdentifier = text("taskIdentifier")
    val errorFingerprint = text("errorFingerprint")
    val status =
        pgEnum("status", "\"ErrorGroupStatus\"", ErrorGroupStatus.entries, ErrorGroupStatus::dbName).default(ErrorGroupStatus.UNRESOLVED)

    /**
     * *
     * * Error is ignored until this date
     */
    val ignoredUntil = pgTimestamp("ignoredUntil", 3).nullable()

    /**
     * *
     * * Error is ignored until this occurrence rate
     */
    val ignoredUntilOccurrenceRate = integer("ignoredUntilOccurrenceRate").nullable()

    /**
     * *
     * * Error is ignored until this total occurrences
     */
    val ignoredUntilTotalOccurrences = integer("ignoredUntilTotalOccurrences").nullable()

    /**
     * Total occurrence count at the time the error was ignored (from ClickHouse).
     * Used with ignoredUntilTotalOccurrences to compute occurrences since ignoring.
     */
    val ignoredAtOccurrenceCount = long("ignoredAtOccurrenceCount").nullable()

    /**
     * *
     * * Error was ignored at this date
     */
    val ignoredAt = pgTimestamp("ignoredAt", 3).nullable()

    /**
     * *
     * * Reason for ignoring the error
     */
    val ignoredReason = text("ignoredReason").nullable()

    /**
     * *
     * * User who ignored the error
     */
    val ignoredByUserId = text("ignoredByUserId").nullable()

    /**
     * *
     * * Error was resolved at this date
     */
    val resolvedAt = pgTimestamp("resolvedAt", 3).nullable()

    /**
     * *
     * * Error was resolved in this version
     */
    val resolvedInVersion = text("resolvedInVersion").nullable()

    /**
     * *
     * * User who resolved the error
     */
    val resolvedBy = text("resolvedBy").nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"ErrorGroupState_pkey\"")

    init {
        index("\"ErrorGroupState_environmentId_status_idx\"", false, environmentId, status)
        uniqueIndex("ErrorGroupState_environmentId_taskIdentifier_errorFingerpri_key", environmentId, taskIdentifier, errorFingerprint)
    }
}
