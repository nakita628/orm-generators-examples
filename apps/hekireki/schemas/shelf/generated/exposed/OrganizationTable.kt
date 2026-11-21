package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object OrganizationTable : IdTable<String>("\"Organization\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val name = text("name").default("Personal")
    val type = pgEnum("type", "\"OrganizationType\"", OrganizationType.entries, OrganizationType::dbName).default(OrganizationType.PERSONAL)
    val userId =
        reference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Organization_userId_fkey\"",
        )
    val currency = pgEnum("currency", "\"Currency\"", Currency.entries, Currency::dbName).default(Currency.USD)
    val imageId =
        optReference(
            "imageId",
            ImageTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Organization_imageId_fkey\"",
        ).uniqueIndex("\"Organization_imageId_key\"")
    val enabledSso = bool("enabledSso").default(false)
    val ssoDetailsId =
        optReference(
            "ssoDetailsId",
            SsoDetailsTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Organization_ssoDetailsId_fkey\"",
        )
    val selfServiceCanSeeCustody = bool("selfServiceCanSeeCustody").default(false)
    val selfServiceCanSeeBookings = bool("selfServiceCanSeeBookings").default(false)
    val baseUserCanSeeCustody = bool("baseUserCanSeeCustody").default(false)
    val baseUserCanSeeBookings = bool("baseUserCanSeeBookings").default(false)
    val managersCanEditAssets = bool("managersCanEditAssets").default(false)
    val barcodesEnabled = bool("barcodesEnabled").default(false)
    val barcodesEnabledAt = pgTimestamp("barcodesEnabledAt", 3).nullable()
    val usedBarcodeTrial = bool("usedBarcodeTrial").default(false)
    val auditsEnabled = bool("auditsEnabled").default(false)
    val auditsEnabledAt = pgTimestamp("auditsEnabledAt", 3).nullable()
    val usedAuditTrial = bool("usedAuditTrial").default(false)
    val workspaceDisabled = bool("workspaceDisabled").default(false)
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val hasSequentialIdsMigrated = bool("hasSequentialIdsMigrated").default(false)
    val qrIdDisplayPreference =
        pgEnum("qrIdDisplayPreference", "\"QrIdDisplayPreference\"", QrIdDisplayPreference.entries, QrIdDisplayPreference::dbName).default(
            QrIdDisplayPreference.QR_ID,
        )
    val showShelfBranding = bool("showShelfBranding").default(true)
    val showQrCodesOnPdfs = bool("showQrCodesOnPdfs").default(true)
    val customEmailFooter = varchar("customEmailFooter", 500).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"Organization_pkey\"")

    init {
        index("\"Organization_userId_idx\"", false, userId)
        index("\"Organization_ssoDetailsId_idx\"", false, ssoDetailsId)
    }
}
