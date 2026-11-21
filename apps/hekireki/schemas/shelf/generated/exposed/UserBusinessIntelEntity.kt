package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class UserBusinessIntelEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, UserBusinessIntelEntity>(UserBusinessIntelTable)

    var howDidYouHearAboutUs by UserBusinessIntelTable.howDidYouHearAboutUs
    var jobTitle by UserBusinessIntelTable.jobTitle
    var teamSize by UserBusinessIntelTable.teamSize
    var companyName by UserBusinessIntelTable.companyName
    var primaryUseCase by UserBusinessIntelTable.primaryUseCase
    var currentSolution by UserBusinessIntelTable.currentSolution
    var timeline by UserBusinessIntelTable.timeline
    var signupPlan by UserBusinessIntelTable.signupPlan
    var signupTrial by UserBusinessIntelTable.signupTrial
    var utmSource by UserBusinessIntelTable.utmSource
    var utmMedium by UserBusinessIntelTable.utmMedium
    var utmCampaign by UserBusinessIntelTable.utmCampaign
    var utmContent by UserBusinessIntelTable.utmContent
    var user by UserEntity referencedOn UserBusinessIntelTable.userId
    var userId by UserBusinessIntelTable.userId
    var createdAt by UserBusinessIntelTable.createdAt
    var updatedAt by UserBusinessIntelTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(UserBusinessIntelTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
