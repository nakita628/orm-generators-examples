package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A survey's response attachments built into one ZIP in object storage by a background job and
 * downloaded through a short-lived presigned URL, so an interrupted download can resume (ENG-6862).
 * Only the requester reads it. The reconcile job deletes the archive at `expiresAt`; deleting the
 * survey deletes it with the survey's storage prefix.
 *
 * A partial unique index in the migration allows one queued or running export per requester and
 * survey, so a double click returns the export already in flight.
 */
object AttachmentExportTable : IdTable<String>("\"AttachmentExport\"") {
    override val id = text("id").clientDefault { CUID.randomCUID2().toString() }.entityId()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val surveyId = text("surveyId")
    val workspaceId =
        reference(
            "workspaceId",
            WorkspaceTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AttachmentExport_workspaceId_fkey\"",
        )
    val requestedById =
        optReference(
            "requestedById",
            UserTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AttachmentExport_requestedById_fkey\"",
        )
    val status =
        pgEnum("status", "\"AttachmentExportStatus\"", AttachmentExportStatus.entries, AttachmentExportStatus::dbName).default(
            AttachmentExportStatus.QUEUED,
        )

    /** The response filter at request time, a `TResponseFilterCriteria`; null exports every response. */
    val filterCriteria = jsonb<String>("filterCriteria", { it }, { it }).nullable()

    /** IANA zone the archive's folder names are stamped in: the org's display zone at request time. */
    val timeZone = text("timeZone")

    /**
     * Bumped on every enqueue and part of the job id, so a re-dispatch never dedupes against a job an
     * older pod dropped as unknown during a rolling upgrade.
     */
    val dispatchCount = integer("dispatchCount").default(0)

    /** The pre-flight's counts, shown while the archive builds. */
    val fileCount = integer("fileCount")
    val responseCount = integer("responseCount")

    /** What the archive actually holds, once built. */
    val includedFileCount = integer("includedFileCount").nullable()
    val totalBytes = long("totalBytes").nullable()
    val truncated = bool("truncated").default(false)
    val storageKey = text("storageKey").nullable()

    /** A short code (`stale`, `storage_upload_failed`, …), never a message or file name. */
    val errorCode = text("errorCode").nullable()
    val startedAt = pgTimestamp("startedAt", 3).nullable()
    val finishedAt = pgTimestamp("finishedAt", 3).nullable()
    val expiresAt = pgTimestamp("expiresAt", 3).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"AttachmentExport_pkey\"")

    init {
        index("\"AttachmentExport_surveyId_requestedById_created_at_idx\"", false, surveyId, requestedById, createdAt)
        index("\"AttachmentExport_status_created_at_idx\"", false, status, createdAt)
        index("\"AttachmentExport_expiresAt_idx\"", false, expiresAt)
        index("\"AttachmentExport_workspaceId_idx\"", false, workspaceId)
        foreignKey(
            surveyId to SurveyTable.id,
            workspaceId to SurveyTable.workspaceId,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            name = "\"AttachmentExport_surveyId_workspaceId_fkey\"",
        )
    }
}
