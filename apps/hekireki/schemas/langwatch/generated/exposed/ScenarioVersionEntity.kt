package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

class ScenarioVersionEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ScenarioVersionEntity>(ScenarioVersionTable)

    var scenarioId by ScenarioVersionTable.scenarioId

    /**
     * The project of the parent case, carried here so every read filters by
     * tenant without a join. The service is what keeps the pair honest: no
     * composite foreign key can, under relationMode="prisma".
     */
    var projectId by ScenarioVersionTable.projectId
    var version by ScenarioVersionTable.version
    var authorId by ScenarioVersionTable.authorId

    /** Who wrote this version: "user", "api", "cli" or "langy". */
    var authorLabel by ScenarioVersionTable.authorLabel
    var changeDescription by ScenarioVersionTable.changeDescription

    /**
     * Envelope { schemaVersion, fields, changedFields } holding the editable
     * field set as saved and the names of the fields this save changed.
     */
    var snapshot by ScenarioVersionTable.snapshot

    /** Schema version of `snapshot`, so a later shape change can be read forward. */
    var schemaVersion by ScenarioVersionTable.schemaVersion
    var createdAt by ScenarioVersionTable.createdAt
}
