package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * The per-branch bookkeeping behind pull-request mapping: when a branch was
 * last looked up on GitHub, how many pull requests it resolved to, and when it
 * may be asked about again. A branch with no pull request is the common case,
 * so `notFoundAt` plus `recheckAfter` is what keeps an unmapped branch from
 * costing a GitHub call on every read.
 * Spec: specs/integrations/github-connection.feature.
 */
object GithubBranchPullRequestCheckTable : IdTable<String>("\"GithubBranchPullRequestCheck\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val organizationId = text("organizationId")
    val repositoryHost = text("repositoryHost")

    /** "owner/repo", stored LOWERCASED, matching GithubPullRequest. */
    val repositoryFullName = text("repositoryFullName")
    val headBranch = text("headBranch")
    val lastCheckedAt = pgTimestamp("lastCheckedAt", 3)
    val prCount = integer("prCount").default(0)

    /**
     * Set while the branch resolves to no pull request; cleared once one is
     * found.
     */
    val notFoundAt = pgTimestamp("notFoundAt", 3).nullable()

    /** The earliest instant GitHub may be asked about this branch again. */
    val recheckAfter = pgTimestamp("recheckAfter", 3).nullable()

    /** Consecutive lookups that found nothing, backing off the recheck. */
    val attempts = integer("attempts").default(0)

    /**
     * When a reader last asked about this branch, whether or not GitHub was
     * called.
     */
    val lastRequestedAt = pgTimestamp("lastRequestedAt", 3)
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"GithubBranchPullRequestCheck_pkey\"")

    init {
        index("GithubBranchPullRequestCheck_recheckAfter_lastRequestedAt_n_idx", false, recheckAfter, lastRequestedAt, notFoundAt)
        uniqueIndex(
            "GithubBranchPullRequestCheck_organizationId_repositoryHost__key",
            organizationId,
            repositoryHost,
            repositoryFullName,
            headBranch,
        )
    }
}
