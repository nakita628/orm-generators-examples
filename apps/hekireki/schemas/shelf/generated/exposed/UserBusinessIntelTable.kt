package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object UserBusinessIntelTable : IdTable<String>("\"UserBusinessIntel\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val howDidYouHearAboutUs = text("howDidYouHearAboutUs").nullable()
    val jobTitle = text("jobTitle").nullable()
    val teamSize = text("teamSize").nullable()
    val companyName = text("companyName").nullable()
    val primaryUseCase = text("primaryUseCase").nullable()
    val currentSolution = text("currentSolution").nullable()
    val timeline = text("timeline").nullable()
    val signupPlan = text("signupPlan").nullable()
    val signupTrial = bool("signupTrial").nullable()
    val utmSource = text("utmSource").nullable()
    val utmMedium = text("utmMedium").nullable()
    val utmCampaign = text("utmCampaign").nullable()
    val utmContent = text("utmContent").nullable()
    val userId =
        reference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"UserBusinessIntel_userId_fkey\"",
        ).uniqueIndex("\"UserBusinessIntel_userId_key\"")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"UserBusinessIntel_pkey\"")

    init {
        index("\"UserBusinessIntel_userId_idx\"", false, userId)
        index("\"UserBusinessIntel_companyName_idx\"", false, companyName)
        index("\"UserBusinessIntel_jobTitle_idx\"", false, jobTitle)
        index("\"UserBusinessIntel_teamSize_idx\"", false, teamSize)
        index("\"UserBusinessIntel_utmSource_idx\"", false, utmSource)
        index("\"UserBusinessIntel_utmCampaign_idx\"", false, utmCampaign)
    }
}
