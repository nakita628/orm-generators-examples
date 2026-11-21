package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Tracks database schema migrations.
 * Used to manage and track the state of data structure changes.
 *
 * @property id - Unique identifier for the migration
 * @property name - Name of the migration
 * @property status - Current state of the migration (pending, applied, failed)
 * @property startedAt - When the migration began
 * @property finishedAt - When the migration completed
 */
class DataMigrationEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, DataMigrationEntity>(DataMigrationTable)

    var startedAt by DataMigrationTable.startedAt
    var finishedAt by DataMigrationTable.finishedAt
    var name by DataMigrationTable.name
    var status by DataMigrationTable.status
}
