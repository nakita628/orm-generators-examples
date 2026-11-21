package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object AuditImageTable : IdTable<String>("\"AuditImage\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val imageUrl = text("imageUrl")
    val thumbnailUrl = text("thumbnailUrl").nullable()
    val description = text("description").nullable()
    val auditSessionId =
        reference(
            "auditSessionId",
            AuditSessionTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AuditImage_auditSessionId_fkey\"",
        )
    val auditAssetId =
        optReference(
            "auditAssetId",
            AuditAssetTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AuditImage_auditAssetId_fkey\"",
        )
    val uploadedById =
        optReference(
            "uploadedById",
            UserTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AuditImage_uploadedById_fkey\"",
        )
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AuditImage_organizationId_fkey\"",
        )

    /**
     * When the photo was taken, from its EXIF data (or the phone app's copy of
     * it). Null when the photo carried no usable date. EXIF can be edited, so
     * this is the best signal available, not proof.
     */
    val capturedAt = pgTimestamp("capturedAt", 3).nullable()

    /**
     * How the photo reached Shelf. Null only for photos uploaded before this
     * field existed; a phone upload with no or an unknown source is PHONE.
     */
    val captureSource =
        pgEnum("captureSource", "\"AuditImageCaptureSource\"", AuditImageCaptureSource.entries, AuditImageCaptureSource::dbName).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"AuditImage_pkey\"")

    init {
        index("\"AuditImage_auditSessionId_idx\"", false, auditSessionId)
        index("\"AuditImage_auditAssetId_idx\"", false, auditAssetId)
        index("\"AuditImage_organizationId_idx\"", false, organizationId)
        index("\"AuditImage_uploadedById_idx\"", false, uploadedById)
    }
}
