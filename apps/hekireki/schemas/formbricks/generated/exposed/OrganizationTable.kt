package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Represents the top-level organizational hierarchy in Formbricks.
 * Self-hosting instances typically have a single organization, while Formbricks Cloud
 * supports multiple organizations with multi-tenancy.
 *
 * @property id - Unique identifier for the organization
 * @property name - Display name of the organization
 * @property memberships - User memberships within the organization
 * @property workspaces - Collection of workspaces owned by the organization
 * @property billing - JSON field containing billing information
 * @property whitelabel - Whitelabel configuration for the organization
 * @property isAISmartToolsEnabled - Controls access to AI smart tools (e.g. translations) that never touch collected data
 * @property displayTimeZone - Optional IANA time zone used for human-facing response timestamps (exports, integrations) and for the calendar days behind the analytics date filters (survey summary, charts, dashboards); null means UTC
 */
object OrganizationTable : IdTable<String>("\"Organization\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val name = text("name")

    /** [OrganizationWhitelabel] */
    val whitelabel = jsonb<String>("whitelabel", { it }, { it }).default("{}")
    val isAISmartToolsEnabled = bool("isAISmartToolsEnabled").default(false)
    val displayTimeZone = text("displayTimeZone").nullable()
    override val primaryKey = PrimaryKey(id, name = "\"Organization_pkey\"")
}
