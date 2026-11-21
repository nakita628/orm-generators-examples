package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
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
object GithubInstallationTable : IdTable<String>("\"GithubInstallation\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()

    /** GitHub installation id (numeric; stored as string). Globally unique. */
    val installationId = text("installationId").uniqueIndex("\"GithubInstallation_installationId_key\"")
    val organizationId = text("organizationId")

    /** The GitHub account (organization or user) the App is installed on. */
    val accountLogin = text("accountLogin")

    /** "Organization" or "User". */
    val accountType = text("accountType")

    /** GitHub numeric account id (stable across renames). */
    val accountId = text("accountId")

    /**
     * "all" or "selected" — whether the App can see every repo on the account
     * or only a chosen subset.
     */
    val repositorySelection = text("repositorySelection")

    /**
     * Cached repository list [{ id, fullName }] when repositorySelection =
     * "selected". Null when "all" (the full set is resolved live at mint time).
     */
    val repositories = jsonb<String>("repositories", { it }, { it }).nullable()

    /**
     * Set when GitHub suspends the installation; cleared on unsuspend. A
     * suspended installation cannot mint tokens.
     */
    val suspendedAt = pgTimestamp("suspendedAt", 3).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"GithubInstallation_pkey\"")

    init {
        index("\"GithubInstallation_organizationId_idx\"", false, organizationId)
    }
}
