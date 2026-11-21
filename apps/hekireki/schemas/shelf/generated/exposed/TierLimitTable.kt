package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object TierLimitTable : IdTable<TierId>("\"TierLimit\"") {
    override val id = pgEnum("id", "\"TierId\"", TierId.entries, TierId::dbName).entityId()
    val canImportAssets = bool("canImportAssets").default(false)
    val canExportAssets = bool("canExportAssets").default(false)
    val canImportNRM = bool("canImportNRM").default(false)
    val canHideShelfBranding = bool("canHideShelfBranding").default(false)
    val canUseReports = bool("canUseReports").default(false)
    val canUseRecurringReminders = bool("canUseRecurringReminders").default(false)
    val canUsePushNotifications = bool("canUsePushNotifications").default(false)
    val maxCustomFields = integer("maxCustomFields").default(0)
    val maxOrganizations = integer("maxOrganizations").default(1)
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"TierLimit_pkey\"")
}
