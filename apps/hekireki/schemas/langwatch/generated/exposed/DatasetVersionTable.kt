package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/** One immutable snapshot of a dataset's columns and rows; the rows are JSONL chunks in object storage. */
object DatasetVersionTable : IdTable<String>("\"DatasetVersion\"") {
    override val id = text("id").entityId()
    val projectId = text("projectId")
    val datasetId = text("datasetId")
    val version = integer("version")

    /** A name given by hand; a named version is never thinned by retention. */
    val name = text("name").nullable()

    /** "autosave", "manual", "experiment_run", "restore" or "initial". */
    val trigger = text("trigger")

    /** SHA-256 over the columns and the rows in order; equal content saves no new version. */
    val contentHash = text("contentHash")
    val columnTypes = jsonb<String>("columnTypes", { it }, { it })
    val rowCount = integer("rowCount")
    val sizeBytes = long("sizeBytes")
    val chunkCount = integer("chunkCount")
    val restoredFromVersion = integer("restoredFromVersion").nullable()
    val authorId = text("authorId").nullable()

    /** "user", "api_key" or "system": who saved the version when authorId is null. */
    val authorKind = text("authorKind").nullable()

    /** Set when an experiment pins this version; a pinned version is never thinned. */
    val pinnedAt = pgTimestamp("pinnedAt", 3).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"DatasetVersion_pkey\"")

    init {
        index("\"DatasetVersion_projectId_datasetId_idx\"", false, projectId, datasetId)
        uniqueIndex("\"DatasetVersion_datasetId_version_key\"", datasetId, version)
    }
}
