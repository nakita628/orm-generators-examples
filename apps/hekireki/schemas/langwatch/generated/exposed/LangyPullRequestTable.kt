package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * One pull request Langy opened, from the answer whose `gh pr create` printed it. Finished
 * checks look it up by repository and number to wake the conversation; the daily pull request
 * cap counts a person's rows by day. Rebuilt by a projection replay.
 */
object LangyPullRequestTable : IdTable<String>("\"LangyPullRequest\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val projectId = text("projectId")

    /** `owner/repo`, lowercased: GitHub matches repository names without case. */
    val repositoryFullName = text("repositoryFullName")
    val prNumber = integer("prNumber")
    val conversationId = text("conversationId")

    /** The conversation's owner, who opened it through Langy. */
    val userId = text("userId")

    /** When the answer that opened it was recorded (epoch ms). */
    val openedAt = double("openedAt")

    /** The agent_responded event it came from. */
    val sourceEventId = text("sourceEventId")
    override val primaryKey = PrimaryKey(id, name = "\"LangyPullRequest_pkey\"")

    init {
        index("\"LangyPullRequest_userId_openedAt_idx\"", false, userId, openedAt)
        uniqueIndex("\"LangyPullRequest_repositoryFullName_prNumber_key\"", repositoryFullName, prNumber)
    }
}
