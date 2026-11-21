package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * What a run did, one row per target: the History CSV. Ids, counts, names and the recipient's email,
 * never response content. `targetId` is not a foreign key on purpose: the target is often gone, and
 * `targetName` is a snapshot taken when the row is written so the CSV still says what it was.
 */
class RetentionRunItemEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, RetentionRunItemEntity>(RetentionRunItemTable)

    var run by RetentionRunEntity referencedOn RetentionRunItemTable.runId
    var runId by RetentionRunItemTable.runId
    var targetType by RetentionRunItemTable.targetType
    var targetId by RetentionRunItemTable.targetId
    var targetName by RetentionRunItemTable.targetName
    var action by RetentionRunItemTable.action

    /** 1, except for grouped response deletions, where it is the number of responses deleted. */
    var count by RetentionRunItemTable.count

    /** Who was emailed, for `notified` rows. */
    var recipient by RetentionRunItemTable.recipient
    var skipReason by RetentionRunItemTable.skipReason
}
