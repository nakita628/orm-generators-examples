package triggerdev.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Defines org-scoped data store overrides (e.g. dedicated ClickHouse for HIPAA orgs).
 * Multiple organizations can share a single data store row via organizationIds.
 */
object OrganizationDataStoreTable : IdTable<String>("\"OrganizationDataStore\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()

    /** Human-readable unique key (e.g. "hipaa-clickhouse-us-east") */
    val key = text("key").uniqueIndex("\"OrganizationDataStore_key_key\"")

    /** Organization IDs that use this data store */
    val organizationIds = array("organizationIds", TextColumnType()).nullable()
    val kind = pgEnum("kind", "\"DataStoreKind\"", DataStoreKind.entries, DataStoreKind::dbName)

    /** Versioned config JSON. Structure is discriminated by the top-level `version` field. */
    val config = jsonb<String>("config", { it }, { it })
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"OrganizationDataStore_pkey\"")

    init {
        index("\"OrganizationDataStore_kind_idx\"", false, kind)
    }
}
