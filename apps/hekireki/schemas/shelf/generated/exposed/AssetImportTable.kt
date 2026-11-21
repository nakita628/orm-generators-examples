package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * One CSV content import. Holds the parsed rows from preview until the import
 * completes, so the job imports exactly the file the user previewed.
 * A partial unique index on (organizationId) WHERE status IN (QUEUED, RUNNING) allows one
 * active import per workspace. Prisma cannot express it; it lives in migration
 * 20261015120100_asset_import_one_active_per_workspace.
 */
object AssetImportTable : IdTable<String>("\"AssetImport\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val status =
        pgEnum("status", "\"AssetImportStatus\"", AssetImportStatus.entries, AssetImportStatus::dbName).default(AssetImportStatus.DRAFT)
    val phase =
        pgEnum("phase", "\"AssetImportPhase\"", AssetImportPhase.entries, AssetImportPhase::dbName).default(AssetImportPhase.PREPARE)
    val fileName = text("fileName").nullable()
    val totalRows = integer("totalRows")

    /** Rows committed so far. Advanced in the same transaction as each chunk. */
    val processedRows = integer("processedRows").default(0)

    /** Parsed rows (`StoredImportRow[]`). Cleared when the import completes. */
    val rows = jsonb<String>("rows", { it }, { it }).nullable()

    /** Name to id maps written by the PREPARE step (`ResolvedImportEntities`). */
    val resolved = jsonb<String>("resolved", { it }, { it }).nullable()
    val canUseBarcodes = bool("canUseBarcodes").default(false)

    /** User-facing reason the import failed. Null unless status is FAILED. */
    val errorMessage = text("errorMessage").nullable()

    /** Rows imported without a conflicting piece (`ImportWarning[]`), capped at 500. */
    val warnings = jsonb<String>("warnings", { it }, { it }).nullable()

    /** Exact number of warnings, including any beyond the stored cap. */
    val warningCount = integer("warningCount").default(0)

    /** A step holds the import until this time. Null when no step is running. */
    val lockedUntil = pgTimestamp("lockedUntil", 3).nullable()
    val confirmedAt = pgTimestamp("confirmedAt", 3).nullable()
    val startedAt = pgTimestamp("startedAt", 3).nullable()
    val completedAt = pgTimestamp("completedAt", 3).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val createdById =
        reference(
            "createdById",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AssetImport_createdById_fkey\"",
        )
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AssetImport_organizationId_fkey\"",
        )
    override val primaryKey = PrimaryKey(id, name = "\"AssetImport_pkey\"")

    init {
        index("\"AssetImport_organizationId_status_idx\"", false, organizationId, status)
        index("\"AssetImport_createdById_organizationId_status_idx\"", false, createdById, organizationId, status)
    }
}
