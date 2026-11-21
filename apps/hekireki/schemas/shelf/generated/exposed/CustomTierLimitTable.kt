package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object CustomTierLimitTable : IdTable<String>("\"CustomTierLimit\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val userId =
        optReference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"CustomTierLimit_userId_fkey\"",
        ).uniqueIndex("\"CustomTierLimit_userId_key\"")
    val canImportAssets = bool("canImportAssets").default(true)
    val canExportAssets = bool("canExportAssets").default(true)
    val canImportNRM = bool("canImportNRM").default(true)
    val canHideShelfBranding = bool("canHideShelfBranding").default(true)
    val canUseReports = bool("canUseReports").default(true)
    val canUseRecurringReminders = bool("canUseRecurringReminders").default(true)
    val canUsePushNotifications = bool("canUsePushNotifications").default(true)
    val maxCustomFields = integer("maxCustomFields").default(1000)
    val maxOrganizations = integer("maxOrganizations").default(1)
    val isEnterprise = bool("isEnterprise").default(false)
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"CustomTierLimit_pkey\"")
}
