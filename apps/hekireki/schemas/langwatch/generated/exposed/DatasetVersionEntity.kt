package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/** One immutable snapshot of a dataset's columns and rows; the rows are JSONL chunks in object storage. */
class DatasetVersionEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, DatasetVersionEntity>(DatasetVersionTable)

    var projectId by DatasetVersionTable.projectId
    var datasetId by DatasetVersionTable.datasetId
    var version by DatasetVersionTable.version

    /** A name given by hand; a named version is never thinned by retention. */
    var name by DatasetVersionTable.name

    /** "autosave", "manual", "experiment_run", "restore" or "initial". */
    var trigger by DatasetVersionTable.trigger

    /** SHA-256 over the columns and the rows in order; equal content saves no new version. */
    var contentHash by DatasetVersionTable.contentHash
    var columnTypes by DatasetVersionTable.columnTypes
    var rowCount by DatasetVersionTable.rowCount
    var sizeBytes by DatasetVersionTable.sizeBytes
    var chunkCount by DatasetVersionTable.chunkCount
    var restoredFromVersion by DatasetVersionTable.restoredFromVersion
    var authorId by DatasetVersionTable.authorId

    /** "user", "api_key" or "system": who saved the version when authorId is null. */
    var authorKind by DatasetVersionTable.authorKind

    /** Set when an experiment pins this version; a pinned version is never thinned. */
    var pinnedAt by DatasetVersionTable.pinnedAt
    var createdAt by DatasetVersionTable.createdAt
}
