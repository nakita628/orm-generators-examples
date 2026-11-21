package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

object ExperimentVersionTable : IdTable<String>("\"ExperimentVersion\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val experimentId = text("experimentId")
    val projectId = text("projectId")

    /**
     * The number the version is addressed and displayed by. Numbered rows hold
     * a gap-free sequence. The rolling autosave row rides on the counter
     * instead, which is what keeps it clear of that sequence.
     */
    val version = integer("version")

    /**
     * `Experiment.workbenchVersion` at the moment this row was last written.
     * The rolling row is rewritten in place, so this is the only column that
     * says how recent its content is, and the version list sorts and pages by
     * it. It also joins a row to the live state: the row whose `counterVersion`
     * equals the experiment's counter holds what the workbench shows now.
     */
    val counterVersion = integer("counterVersion")
    val autoSaved = bool("autoSaved").default(false)
    val commitMessage = text("commitMessage").nullable()
    val authorId = text("authorId").nullable()

    /** Who wrote this version: "user", "langy" or "api". */
    val authorLabel = text("authorLabel")

    /**
     * The run that wrote this version, when a run wrote it.
     *
     * A run writes its cells into the workbench state, which advances the
     * counter. Without this, the page that started that run reads the bump as
     * somebody else's write, refuses to save, and asks the reader to reload
     * over their own unsaved edits. Named here, the page adopts the version and
     * carries on.
     */
    val runId = text("runId").nullable()
    val state = jsonb<String>("state", { it }, { it })

    /** Schema version of `state`, so a later shape change can be read forward. */
    val schemaVersion = text("schemaVersion").default("1")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"ExperimentVersion_pkey\"")

    init {
        index("\"ExperimentVersion_projectId_experimentId_autoSaved_idx\"", false, projectId, experimentId, autoSaved)
        index("\"ExperimentVersion_projectId_experimentId_version_idx\"", false, projectId, experimentId, version)
        index("\"ExperimentVersion_projectId_experimentId_counterVersion_idx\"", false, projectId, experimentId, counterVersion)
        index("\"ExperimentVersion_createdAt_idx\"", false, createdAt)
        uniqueIndex("\"ExperimentVersion_experimentId_version_key\"", experimentId, version)
    }
}
