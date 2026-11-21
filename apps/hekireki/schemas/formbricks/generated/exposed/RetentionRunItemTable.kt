package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * What a run did, one row per target: the History CSV. Ids, counts, names and the recipient's email,
 * never response content. `targetId` is not a foreign key on purpose: the target is often gone, and
 * `targetName` is a snapshot taken when the row is written so the CSV still says what it was.
 */
object RetentionRunItemTable : IdTable<String>("\"RetentionRunItem\"") {
    override val id = text("id").clientDefault { CUID.randomCUID2().toString() }.entityId()
    val runId =
        reference(
            "runId",
            RetentionRunTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"RetentionRunItem_runId_fkey\"",
        )
    val targetType = pgEnum("targetType", "\"RetentionTargetType\"", RetentionTargetType.entries, RetentionTargetType::dbName)
    val targetId = text("targetId")
    val targetName = text("targetName").nullable()
    val action = pgEnum("action", "\"RetentionRunItemAction\"", RetentionRunItemAction.entries, RetentionRunItemAction::dbName)

    /** 1, except for grouped response deletions, where it is the number of responses deleted. */
    val count = integer("count").default(1)

    /** Who was emailed, for `notified` rows. */
    val recipient = text("recipient").nullable()
    val skipReason = pgEnum("skipReason", "\"RetentionSkipReason\"", RetentionSkipReason.entries, RetentionSkipReason::dbName).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"RetentionRunItem_pkey\"")

    init {
        index("\"RetentionRunItem_runId_id_idx\"", false, runId, id)
        index("\"RetentionRunItem_targetId_idx\"", false, targetId)
    }
}
