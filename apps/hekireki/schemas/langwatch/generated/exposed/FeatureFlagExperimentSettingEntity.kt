package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * One experiment setting per (flag, subject). A USER row is that person's
 * own opt-in; an ORGANIZATION or PROJECT row is an owner's explicit policy
 * for that exact scope. No row means inherit.
 */
class FeatureFlagExperimentSettingEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<FeatureFlagExperimentSettingEntity>(FeatureFlagExperimentSettingTable)

    val flagKey by FeatureFlagExperimentSettingTable.flagKey
    val subjectType by FeatureFlagExperimentSettingTable.subjectType
    val subjectId by FeatureFlagExperimentSettingTable.subjectId
    var enabled by FeatureFlagExperimentSettingTable.enabled
    var changedByUserId by FeatureFlagExperimentSettingTable.changedByUserId
    var createdAt by FeatureFlagExperimentSettingTable.createdAt
    var updatedAt by FeatureFlagExperimentSettingTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(FeatureFlagExperimentSettingTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
