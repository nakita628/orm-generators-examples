package langwatch.models

import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

object SessionTable : IdTable<String>("\"Session\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val sessionToken = text("sessionToken").uniqueIndex("\"Session_sessionToken_key\"")
    val userId = text("userId")
    val expires = pgTimestamp("expires", 3)
    val ipAddress = text("ipAddress").nullable()
    val userAgent = text("userAgent").nullable()

    /**
     * Which of the user's sign-in methods minted this session (D06). NULL on
     * every session that predates the column and on any mint we could not
     * attribute - an ordinary value, not an error. Nothing reads it to decide
     * whether the session is valid, so no session was ended to add it.
     */
    val identifierId = text("identifierId").nullable()

    /**
     * What the sign-in that minted this session PROVED, as RFC 8176 method
     * references (D06). Empty on every session minted before the column, and
     * on any ordinary password sign-in. An organization that requires a second
     * factor reads this when a member reaches its data; it is never read at
     * mint and never ends a session.
     */
    val amr = array("amr", TextColumnType()).nullable().databaseDefault("'{}'") { emptyList() }

    /**
     * Minted while the account had never signed in: the session its sign-up
     * started, which skips the secure-account offer. False on every later one
     * and on every session that predates the column (specs/identity/passkeys.feature).
     */
    val startedBySignUp = bool("startedBySignUp").default(false)

    /**
     * The legacy impersonation payload. NOTHING READS OR WRITES IT any more -
     * the {actor, subject} columns below replaced it, and the migration beside
     * them ended every session that still carried a value.
     *
     * It stays for ONE release, and that is the whole reason it is still here.
     * `prisma migrate deploy` runs at container start, so the first pod of a
     * new release drops a column while the previous release's pods are still
     * serving - and those pods select this column on EVERY authenticated
     * request and fail closed when the row does not come back. Dropping it in
     * the same release that stops reading it signs everybody out for the length
     * of the rollout. Expand now, contract next release.
     */
    val impersonating = jsonb<String>("impersonating", { it }, { it }).nullable()

    /**
     * The impersonation claims (D06), aligned with the authz Principal
     * {actor, subject}. The actor is the operator who really made the request
     * and the subject is the person whose access they are borrowing; both are
     * recorded on every authorization decision, so the audit trail can answer
     * who really did it. These replace the legacy `impersonating` JSON blob,
     * which is dropped in the release AFTER this one.
     *
     * NULL on every ordinary session, which is all of them: an impersonation
     * is a handful of rows at a time and nothing else reads these columns.
     */
    val actorUserId = text("actorUserId").nullable()
    val subjectUserId = text("subjectUserId").nullable()

    /**
     * Why the operator started it. Required to start one
     * (specs/features/backoffice-user-impersonation-reason.feature) and kept
     * beside both people rather than only in the audit log, so a request under
     * impersonation carries its own justification.
     */
    val impersonationReason = text("impersonationReason").nullable()

    /**
     * When the borrowed access lapses. A session outlives its impersonation
     * window: past this the claims are ignored and the operator is themselves
     * again, with nothing ended.
     */
    val impersonationExpiresAt = pgTimestamp("impersonationExpiresAt", 3).nullable()

    /**
     * When this session was last used, for the idle timeout an organization
     * may set (GAC-10, specs/identity/org-session-lifetime.feature).
     *
     * A COLUMN OF OUR OWN rather than `updatedAt`, because better-auth rolls a
     * live session's expiry once per `updateAge` - a day - so `updatedAt` is
     * activity to the nearest day and cannot tell an idle hour from an idle
     * minute. Written only for sessions whose owner belongs to an organization
     * that has set a window, and then at most once per quarter-window, so a
     * deployment where nobody has turned this on writes it never.
     *
     * NULL on every session minted before the column and on every session
     * under no window at all. A null READS AS `updatedAt`, which better-auth
     * keeps rolled to within a day of the last use. That fallback is
     * conservative in the direction that matters - `updatedAt` is never
     * earlier than the real last use, so it can only ever judge a session MORE
     * active than it was, never less - and it is what lets a window turned on
     * today end the sessions that have been idle since February instead of
     * waiting a day for a stamp they have never had.
     */
    val lastSeenAt = pgTimestamp("lastSeenAt", 3).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"Session_pkey\"")

    init {
        index("\"Session_userId_idx\"", false, userId)
        index("\"Session_identifierId_idx\"", false, identifierId)
    }
}
