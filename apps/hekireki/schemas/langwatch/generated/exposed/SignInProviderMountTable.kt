package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object SignInProviderMountTable : IdTable<String>("\"SignInProviderMount\"") {
    /**
     * Fixed, so this table holds one row: licensing's copy of sso's sign-in
     * provider fact (ADR-179).
     */
    override val id = text("id").default("self").entityId()

    /** The configured provider's name; null when the deployment signs in by email. */
    val provider = text("provider").nullable()
    val mounted = bool("mounted")

    /** The last sso event applied: the row's idempotency guard and its version. */
    val lastAppliedEventId = text("lastAppliedEventId")
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"SignInProviderMount_pkey\"")
}
