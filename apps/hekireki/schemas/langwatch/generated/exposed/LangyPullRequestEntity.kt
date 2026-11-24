package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * One pull request Langy opened, from the answer whose `gh pr create` printed it. Finished
 * checks look it up by repository and number to wake the conversation; the daily pull request
 * cap counts a person's rows by day. Rebuilt by a projection replay.
 */
class LangyPullRequestEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, LangyPullRequestEntity>(LangyPullRequestTable)

    var projectId by LangyPullRequestTable.projectId

    /** `owner/repo`, lowercased: GitHub matches repository names without case. */
    var repositoryFullName by LangyPullRequestTable.repositoryFullName
    var prNumber by LangyPullRequestTable.prNumber
    var conversationId by LangyPullRequestTable.conversationId

    /** The conversation's owner, who opened it through Langy. */
    var userId by LangyPullRequestTable.userId

    /** When the answer that opened it was recorded (epoch ms). */
    var openedAt by LangyPullRequestTable.openedAt

    /** The agent_responded event it came from. */
    var sourceEventId by LangyPullRequestTable.sourceEventId
}
