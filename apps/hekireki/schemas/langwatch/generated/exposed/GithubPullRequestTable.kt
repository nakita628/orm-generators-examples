package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
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
object GithubPullRequestTable : IdTable<String>("\"GithubPullRequest\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val organizationId = text("organizationId")

    /** The GitHub host the repository lives on, e.g. "github.com". */
    val repositoryHost = text("repositoryHost")

    /**
     * "owner/repo", stored LOWERCASED so a lookup matches whatever casing the
     * session reported.
     */
    val repositoryFullName = text("repositoryFullName")

    /** The branch the pull request was opened from. */
    val headBranch = text("headBranch")
    val prNumber = integer("prNumber")
    val htmlUrl = text("htmlUrl")
    val title = text("title")

    /** GitHub's own state string: "open" or "closed". */
    val state = text("state")
    val isDraft = bool("isDraft").default(false)
    val authorLogin = text("authorLogin").nullable()
    val prCreatedAt = pgTimestamp("prCreatedAt", 3)
    val prClosedAt = pgTimestamp("prClosedAt", 3).nullable()
    val prMergedAt = pgTimestamp("prMergedAt", 3).nullable()

    /**
     * GitHub's own `updated_at` for the snapshot this row holds, and the
     * freshness guard on every write to it: a snapshot older than the stored one
     * is skipped. GitHub may deliver webhooks out of order, and without the
     * guard a late `opened` lands after `closed` and reopens a merged pull
     * request. NULL means the row predates the column, so any write is accepted.
     */
    val prUpdatedAt = pgTimestamp("prUpdatedAt", 3).nullable()

    /** When this row was first mapped to its branch. */
    val mappedAt = pgTimestamp("mappedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }

    /** When the snapshot columns were last refreshed from GitHub. */
    val lastCheckedAt =
        pgTimestamp("lastCheckedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"GithubPullRequest_pkey\"")

    init {
        index(
            "GithubPullRequest_organizationId_repositoryHost_repositoryF_idx",
            false,
            organizationId,
            repositoryHost,
            repositoryFullName,
            headBranch,
        )
        uniqueIndex(
            "GithubPullRequest_organizationId_repositoryHost_repositoryF_key",
            organizationId,
            repositoryHost,
            repositoryFullName,
            prNumber,
        )
    }
}
