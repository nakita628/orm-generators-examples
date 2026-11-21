package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
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
class GithubBranchPullRequestCheckEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, GithubBranchPullRequestCheckEntity>(GithubBranchPullRequestCheckTable)

    var organizationId by GithubBranchPullRequestCheckTable.organizationId
    var repositoryHost by GithubBranchPullRequestCheckTable.repositoryHost

    /** "owner/repo", stored LOWERCASED, matching GithubPullRequest. */
    var repositoryFullName by GithubBranchPullRequestCheckTable.repositoryFullName
    var headBranch by GithubBranchPullRequestCheckTable.headBranch
    var lastCheckedAt by GithubBranchPullRequestCheckTable.lastCheckedAt
    var prCount by GithubBranchPullRequestCheckTable.prCount

    /**
     * Set while the branch resolves to no pull request; cleared once one is
     * found.
     */
    var notFoundAt by GithubBranchPullRequestCheckTable.notFoundAt

    /** The earliest instant GitHub may be asked about this branch again. */
    var recheckAfter by GithubBranchPullRequestCheckTable.recheckAfter

    /** Consecutive lookups that found nothing, backing off the recheck. */
    var attempts by GithubBranchPullRequestCheckTable.attempts

    /**
     * When a reader last asked about this branch, whether or not GitHub was
     * called.
     */
    var lastRequestedAt by GithubBranchPullRequestCheckTable.lastRequestedAt
    var createdAt by GithubBranchPullRequestCheckTable.createdAt
    var updatedAt by GithubBranchPullRequestCheckTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(GithubBranchPullRequestCheckTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
