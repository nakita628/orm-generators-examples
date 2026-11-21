package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

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
object DataMigrationTable : IdTable<String>("\"DataMigration\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val startedAt = pgTimestamp("started_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val finishedAt = pgTimestamp("finished_at", 3).nullable()
    val name = text("name").uniqueIndex("\"DataMigration_name_key\"")
    val status = pgEnum("status", "\"DataMigrationStatus\"", DataMigrationStatus.entries, DataMigrationStatus::dbName)
    override val primaryKey = PrimaryKey(id, name = "\"DataMigration_pkey\"")
}
