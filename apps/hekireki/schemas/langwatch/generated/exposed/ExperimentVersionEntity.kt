package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class ExperimentVersionEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ExperimentVersionEntity>(ExperimentVersionTable)

    var experimentId by ExperimentVersionTable.experimentId
    var projectId by ExperimentVersionTable.projectId

    /**
     * The number the version is addressed and displayed by. Numbered rows hold
     * a gap-free sequence. The rolling autosave row rides on the counter
     * instead, which is what keeps it clear of that sequence.
     */
    var version by ExperimentVersionTable.version

    /**
     * `Experiment.workbenchVersion` at the moment this row was last written.
     * The rolling row is rewritten in place, so this is the only column that
     * says how recent its content is, and the version list sorts and pages by
     * it. It also joins a row to the live state: the row whose `counterVersion`
     * equals the experiment's counter holds what the workbench shows now.
     */
    var counterVersion by ExperimentVersionTable.counterVersion
    var autoSaved by ExperimentVersionTable.autoSaved
    var commitMessage by ExperimentVersionTable.commitMessage
    var authorId by ExperimentVersionTable.authorId

    /** Who wrote this version: "user", "langy" or "api". */
    var authorLabel by ExperimentVersionTable.authorLabel

    /**
     * The run that wrote this version, when a run wrote it.
     *
     * A run writes its cells into the workbench state, which advances the
     * counter. Without this, the page that started that run reads the bump as
     * somebody else's write, refuses to save, and asks the reader to reload
     * over their own unsaved edits. Named here, the page adopts the version and
     * carries on.
     */
    var runId by ExperimentVersionTable.runId
    var state by ExperimentVersionTable.state

    /** Schema version of `state`, so a later shape change can be read forward. */
    var schemaVersion by ExperimentVersionTable.schemaVersion
    var createdAt by ExperimentVersionTable.createdAt
    var updatedAt by ExperimentVersionTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(ExperimentVersionTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
