package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
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
class AttachmentExportEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AttachmentExportEntity>(AttachmentExportTable)

    var createdAt by AttachmentExportTable.createdAt
    var updatedAt by AttachmentExportTable.updatedAt
    var surveyId by AttachmentExportTable.surveyId
    var workspace by WorkspaceEntity referencedOn AttachmentExportTable.workspaceId
    var workspaceId by AttachmentExportTable.workspaceId
    var requestedBy by UserEntity optionalReferencedOn AttachmentExportTable.requestedById
    var requestedById by AttachmentExportTable.requestedById
    var status by AttachmentExportTable.status

    /** The response filter at request time, a `TResponseFilterCriteria`; null exports every response. */
    var filterCriteria by AttachmentExportTable.filterCriteria

    /** IANA zone the archive's folder names are stamped in: the org's display zone at request time. */
    var timeZone by AttachmentExportTable.timeZone

    /**
     * Bumped on every enqueue and part of the job id, so a re-dispatch never dedupes against a job an
     * older pod dropped as unknown during a rolling upgrade.
     */
    var dispatchCount by AttachmentExportTable.dispatchCount

    /** The pre-flight's counts, shown while the archive builds. */
    var fileCount by AttachmentExportTable.fileCount
    var responseCount by AttachmentExportTable.responseCount

    /** What the archive actually holds, once built. */
    var includedFileCount by AttachmentExportTable.includedFileCount
    var totalBytes by AttachmentExportTable.totalBytes
    var truncated by AttachmentExportTable.truncated
    var storageKey by AttachmentExportTable.storageKey

    /** A short code (`stale`, `storage_upload_failed`, …), never a message or file name. */
    var errorCode by AttachmentExportTable.errorCode
    var startedAt by AttachmentExportTable.startedAt
    var finishedAt by AttachmentExportTable.finishedAt
    var expiresAt by AttachmentExportTable.expiresAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AttachmentExportTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
