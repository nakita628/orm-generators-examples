package teable.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object FieldPeriodicRefreshStateTable : IdTable<String>("field_periodic_refresh_state") {
    override val id =
        text("field_id")
            .references(
                FieldTable.id,
                onDelete = ReferenceOption.CASCADE,
                onUpdate = ReferenceOption.CASCADE,
                fkName = "field_periodic_refresh_state_field_id_fkey",
            ).entityId()
    val tableId = text("table_id")
    val baseId = text("base_id")
    val enabled = bool("enabled").default(true)
    val timeZone = varchar("time_zone", 64).default("UTC")
    val containsNowPrecision = bool("contains_now_precision").default(false)
    val lastRefreshedAt = pgTimestamptz("last_refreshed_at", 6).nullable()
    val lastRefreshedLocalDate = varchar("last_refreshed_local_date", 16).nullable()
    val lastRangeHash = varchar("last_range_hash", 128).nullable()
    val lastStatus = varchar("last_status", 32).default("idle")
    val lastError = text("last_error").nullable()
    val pendingTaskId = varchar("pending_task_id", 64).nullable()
    val pendingLocalDate = varchar("pending_local_date", 16).nullable()
    val pendingRangeHash = varchar("pending_range_hash", 128).nullable()
    val pendingStartedAt = pgTimestamptz("pending_started_at", 6).nullable()
    val createdTime = pgTimestamptz("created_time", 6).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val lastModifiedTime = pgTimestamptz("last_modified_time", 6).nullable().clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "field_periodic_refresh_state_pkey")

    init {
        index("field_periodic_refresh_state_base_id_enabled_idx", false, baseId, enabled)
        index("field_periodic_refresh_state_table_id_idx", false, tableId)
        index("field_periodic_refresh_state_enabled_local_date_idx", false, enabled, lastRefreshedLocalDate)
    }
}
