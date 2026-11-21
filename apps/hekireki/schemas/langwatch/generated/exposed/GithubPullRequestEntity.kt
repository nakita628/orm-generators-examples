package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A pull request discovered through the organization's GitHub connection and
 * mapped to the branch it was opened from, so a coding-agent session that
 * knows only its branch can name the pull request it produced.
 * Spec: specs/integrations/github-connection.feature.
 *
 * The stored title, state, draft and timestamp columns are SNAPSHOTS taken
 * when the mapping was last refreshed. They are fallback labels only: the
 * frontend reads live status from GitHub, so a stale row never presents
 * itself as current truth.
 */
class GithubPullRequestEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, GithubPullRequestEntity>(GithubPullRequestTable)

    var organizationId by GithubPullRequestTable.organizationId

    /** The GitHub host the repository lives on, e.g. "github.com". */
    var repositoryHost by GithubPullRequestTable.repositoryHost

    /**
     * "owner/repo", stored LOWERCASED so a lookup matches whatever casing the
     * session reported.
     */
    var repositoryFullName by GithubPullRequestTable.repositoryFullName

    /** The branch the pull request was opened from. */
    var headBranch by GithubPullRequestTable.headBranch
    var prNumber by GithubPullRequestTable.prNumber
    var htmlUrl by GithubPullRequestTable.htmlUrl
    var title by GithubPullRequestTable.title

    /** GitHub's own state string: "open" or "closed". */
    var state by GithubPullRequestTable.state
    var isDraft by GithubPullRequestTable.isDraft
    var authorLogin by GithubPullRequestTable.authorLogin
    var prCreatedAt by GithubPullRequestTable.prCreatedAt
    var prClosedAt by GithubPullRequestTable.prClosedAt
    var prMergedAt by GithubPullRequestTable.prMergedAt

    /**
     * GitHub's own `updated_at` for the snapshot this row holds, and the
     * freshness guard on every write to it: a snapshot older than the stored one
     * is skipped. GitHub may deliver webhooks out of order, and without the
     * guard a late `opened` lands after `closed` and reopens a merged pull
     * request. NULL means the row predates the column, so any write is accepted.
     */
    var prUpdatedAt by GithubPullRequestTable.prUpdatedAt

    /** When this row was first mapped to its branch. */
    var mappedAt by GithubPullRequestTable.mappedAt

    /** When the snapshot columns were last refreshed from GitHub. */
    var lastCheckedAt by GithubPullRequestTable.lastCheckedAt
    var createdAt by GithubPullRequestTable.createdAt
    var updatedAt by GithubPullRequestTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(GithubPullRequestTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
