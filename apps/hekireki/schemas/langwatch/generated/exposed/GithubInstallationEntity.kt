package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * The organization's GitHub connection: a GitHub App installation mapped to a
 * LangWatch organization. Short-lived (1h) installation access tokens are
 * minted against these rows on demand, by Langy for writes (bot-authored pull
 * requests, scoped to the installation's repositories) and by pull-request
 * linkage for reads. No user token and no secret are stored here, the crown
 * jewel is the App private key, held only in the control-plane env.
 * Spec: specs/integrations/github-connection.feature.
 *
 * installationId is globally unique (one GitHub installation, one row); a
 * single LangWatch organization may own MANY installations (different GitHub
 * accounts), so the org to installation relationship is one-to-many.
 */
class GithubInstallationEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, GithubInstallationEntity>(GithubInstallationTable)

    /** GitHub installation id (numeric; stored as string). Globally unique. */
    var installationId by GithubInstallationTable.installationId
    var organizationId by GithubInstallationTable.organizationId

    /** The GitHub account (organization or user) the App is installed on. */
    var accountLogin by GithubInstallationTable.accountLogin

    /** "Organization" or "User". */
    var accountType by GithubInstallationTable.accountType

    /** GitHub numeric account id (stable across renames). */
    var accountId by GithubInstallationTable.accountId

    /**
     * "all" or "selected" — whether the App can see every repo on the account
     * or only a chosen subset.
     */
    var repositorySelection by GithubInstallationTable.repositorySelection

    /**
     * Cached repository list [{ id, fullName }] when repositorySelection =
     * "selected". Null when "all" (the full set is resolved live at mint time).
     */
    var repositories by GithubInstallationTable.repositories

    /**
     * Set when GitHub suspends the installation; cleared on unsuspend. A
     * suspended installation cannot mint tokens.
     */
    var suspendedAt by GithubInstallationTable.suspendedAt
    var createdAt by GithubInstallationTable.createdAt
    var updatedAt by GithubInstallationTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(GithubInstallationTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
