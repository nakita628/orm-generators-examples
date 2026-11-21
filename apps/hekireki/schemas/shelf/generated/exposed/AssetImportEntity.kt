package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * One CSV content import. Holds the parsed rows from preview until the import
 * completes, so the job imports exactly the file the user previewed.
 * A partial unique index on (organizationId) WHERE status IN (QUEUED, RUNNING) allows one
 * active import per workspace. Prisma cannot express it; it lives in migration
 * 20261015120100_asset_import_one_active_per_workspace.
 */
class AssetImportEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AssetImportEntity>(AssetImportTable)

    var status by AssetImportTable.status
    var phase by AssetImportTable.phase
    var fileName by AssetImportTable.fileName
    var totalRows by AssetImportTable.totalRows

    /** Rows committed so far. Advanced in the same transaction as each chunk. */
    var processedRows by AssetImportTable.processedRows

    /** Parsed rows (`StoredImportRow[]`). Cleared when the import completes. */
    var rows by AssetImportTable.rows

    /** Name to id maps written by the PREPARE step (`ResolvedImportEntities`). */
    var resolved by AssetImportTable.resolved
    var canUseBarcodes by AssetImportTable.canUseBarcodes

    /** User-facing reason the import failed. Null unless status is FAILED. */
    var errorMessage by AssetImportTable.errorMessage

    /** Rows imported without a conflicting piece (`ImportWarning[]`), capped at 500. */
    var warnings by AssetImportTable.warnings

    /** Exact number of warnings, including any beyond the stored cap. */
    var warningCount by AssetImportTable.warningCount

    /** A step holds the import until this time. Null when no step is running. */
    var lockedUntil by AssetImportTable.lockedUntil
    var confirmedAt by AssetImportTable.confirmedAt
    var startedAt by AssetImportTable.startedAt
    var completedAt by AssetImportTable.completedAt
    var createdAt by AssetImportTable.createdAt
    var updatedAt by AssetImportTable.updatedAt
    var createdBy by UserEntity referencedOn AssetImportTable.createdById
    var createdById by AssetImportTable.createdById
    var organization by OrganizationEntity referencedOn AssetImportTable.organizationId
    var organizationId by AssetImportTable.organizationId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AssetImportTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
