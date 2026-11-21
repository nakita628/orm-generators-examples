package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * The values an `SsoConnection`'s credential references point AT (D09 - see
 * specs/identity/sso-idp-termination.feature).
 *
 * The aggregate has always carried `clientIdRef` / `secretRef` / `certRefs`
 * rather than values, because a fact may not carry a secret (ADR-101 §4).
 * Until D09 there was nothing behind the references and every writer passed
 * null. This is the thing behind them.
 *
 * `ciphertext` is AES-256-GCM under the deployment's `CREDENTIALS_SECRET`,
 * the same key every other credential in the product is kept under. Nothing
 * reads a row except the projection that builds the engine's dialing
 * configuration; no surface renders one back.
 *
 * Written by the command path rather than the fold, and that split is
 * deliberate: the VALUE is not event truth and cannot be, so the log records
 * WHEN a credential changed (a new reference in a new fact) and never what
 * it was. A replay rebuilds every reference from the log and finds the value
 * still here.
 */
object SsoCredentialTable : IdTable<String>("\"SsoCredential\"") {
    override val id = text("id").entityId()
    val organizationId = text("organizationId")

    /**
     * The connection whose registration minted this. Not a relation, matching
     * `SsoConnection` itself: a projection rebuilt from events must not depend
     * on relational state the log does not carry.
     */
    val connectionId = text("connectionId")

    /** oidc-client-id | oidc-client-secret | saml-idp-config */
    val kind = text("kind")
    val ciphertext = text("ciphertext")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"SsoCredential_pkey\"")

    init {
        index("\"SsoCredential_organizationId_connectionId_idx\"", false, organizationId, connectionId)
    }
}
